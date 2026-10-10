package services

import (
	"fmt"
	"os"
	"time"

	"jk-api/api/http/controllers/v1/dto"
	"jk-api/internal/config"
	"jk-api/internal/database/models"
	"jk-api/internal/errors/bcrypt_err"
	"jk-api/internal/errors/gorm_err"
	"jk-api/pkg/repository/adapter/sql"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthService interface {
	Login(email, password string) (*models.User, error)
	Register(req *dto.RegisterRequest) (*models.User, error)
	GetProfile(token string) (*models.User, error)
	Logout(token string) error
	GenerateAccessToken(user *models.User) (string, error)
	GenerateRefreshToken(user *models.User) (string, error)
	RefreshToken(token string) (string, error)
	DecodeToken(token string) (jwt.MapClaims, error)
	GetPendingTeacherRequests() ([]models.RegistrationRequest, error)
	ProcessTeacherRequest(requestID int64, adminID int64, decision dto.TeacherApprovalDecisionRequest) error
}

type authService struct {
	repo             sql.UserRepository
	refreshTokenRepo sql.RefreshTokenRepository
}

func NewAuthService(
	userRepo sql.UserRepository,
	refreshRepo sql.RefreshTokenRepository,
) *authService {
	return &authService{
		repo:             userRepo,
		refreshTokenRepo: refreshRepo,
	}
}

func (s *authService) Login(email, password string) (*models.User, error) {
	user, err := s.repo.FindUserByEmail(email)

	if err != nil {
		return nil, gorm_err.TranslateGormError(err)
	}

	// ❌ Cek user aktif
	if !user.IsActive {
		return nil, fmt.Errorf("akun tidak aktif")
	}

	// ❌ Cek approval teacher
	if user.Role == models.RoleTeacher && user.Approval == 0 {
		return nil, fmt.Errorf("akun teacher belum di-approve admin")
	} else if user.Role == models.RoleTeacher && user.Approval == 2 {
		return nil, fmt.Errorf("akun teacher ditolak admin")
	}

	// ❌ Cek password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, bcrypt_err.TranslateBcryptError(err)
	}

	return user, nil
}

func (s *authService) Register(req *dto.RegisterRequest) (*models.User, error) {
	// Validasi role di trust boundary: hanya student/teacher yang boleh self-register,
	// supaya role dari body request tidak bisa dipakai jadi super/admin.
	role := models.UserRole(req.Role)
	if role != models.RoleStudent && role != models.RoleTeacher {
		return nil, fmt.Errorf("role tidak valid untuk registrasi")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	if role == models.RoleStudent {
		if req.ClassCode == nil || *req.ClassCode == "" {
			return nil, fmt.Errorf("kode kelas wajib diisi untuk registrasi student")
		}

		var class models.MClass
		if err := config.DB.Where("class_code = ?", *req.ClassCode).First(&class).Error; err != nil {
			return nil, fmt.Errorf("kode kelas tidak ditemukan")
		}

		if !class.IsActive {
			return nil, fmt.Errorf("kelas tidak aktif")
		}

		var school models.MSchool
		if err := config.DB.Where("id = ?", class.SchoolID).First(&school).Error; err != nil {
			return nil, fmt.Errorf("sekolah tidak ditemukan")
		}

		if !school.IsActive {
			return nil, fmt.Errorf("sekolah tidak aktif")
		}

		if class.TeacherID == nil {
			return nil, fmt.Errorf("kelas belum memiliki pengajar yang valid")
		}

		var createdUser models.User
		err = config.DB.Transaction(func(tx *gorm.DB) error {
			createdUser = models.User{
				Name:     req.Name,
				Email:    req.Email,
				Password: string(hashedPassword),
				Role:     models.RoleStudent,
				ClassID:  &class.ID,
				SchoolID: &class.SchoolID,
				Approval: 1,
				IsActive: true,
			}
			return tx.Create(&createdUser).Error
		})

		if err != nil {
			return nil, gorm_err.TranslateGormError(err)
		}

		createdUser.Password = ""
		return &createdUser, nil
	}

	// Alur Registrasi Teacher
	if req.SchoolID == nil && (req.ProposedSchoolName == nil || *req.ProposedSchoolName == "") {
		return nil, fmt.Errorf("sekolah atau usulan nama sekolah baru wajib disertakan")
	}

	if req.SchoolID != nil {
		var school models.MSchool
		if err := config.DB.Where("id = ? AND isactive = true", *req.SchoolID).First(&school).Error; err != nil {
			return nil, fmt.Errorf("sekolah tidak ditemukan atau tidak aktif")
		}
	}

	var createdUser models.User
	err = config.DB.Transaction(func(tx *gorm.DB) error {
		createdUser = models.User{
			Name:     req.Name,
			Email:    req.Email,
			Password: string(hashedPassword),
			Role:     models.RoleTeacher,
			Nip:      req.NIP,
			SchoolID: req.SchoolID,
			Approval: 0, // Menunggu approval admin
			IsActive: true,
		}

		if err := tx.Create(&createdUser).Error; err != nil {
			return err
		}

		regReq := models.RegistrationRequest{
			Name:     req.Name,
			Email:    req.Email,
			Password: string(hashedPassword),
			Nip:      req.NIP,
			Status:   "pending",
		}
		if err := tx.Create(&regReq).Error; err != nil {
			return err
		}

		regSchool := models.RegistrationRequestSchool{
			RegistrationRequestID: regReq.ID,
			ExistingSchoolID:      req.SchoolID,
			ProposedSchoolName:    req.ProposedSchoolName,
			Status:                "pending",
		}
		if err := tx.Create(&regSchool).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, gorm_err.TranslateGormError(err)
	}

	createdUser.Password = ""
	return &createdUser, nil
}

func (s *authService) GetPendingTeacherRequests() ([]models.RegistrationRequest, error) {
	var requests []models.RegistrationRequest
	err := config.DB.
		Preload("School").
		Preload("School.ExistingSchool").
		Where("status = ?", "pending").
		Find(&requests).Error
	if err != nil {
		return nil, gorm_err.TranslateGormError(err)
	}
	return requests, nil
}

func (s *authService) ProcessTeacherRequest(requestID int64, adminID int64, decision dto.TeacherApprovalDecisionRequest) error {
	return config.DB.Transaction(func(tx *gorm.DB) error {
		var regReq models.RegistrationRequest
		if err := tx.Preload("School").Where("id = ?", requestID).First(&regReq).Error; err != nil {
			return fmt.Errorf("permohonan registrasi tidak ditemukan")
		}

		if regReq.Status != "pending" {
			return fmt.Errorf("permohonan registrasi sudah diproses sebelumnya")
		}

		var teacher models.User
		if err := tx.Where("email = ?", regReq.Email).First(&teacher).Error; err != nil {
			return fmt.Errorf("akun guru untuk permohonan ini tidak ditemukan")
		}

		if decision.Action == "approve" {
			var schoolID int64
			if regReq.School != nil && regReq.School.ProposedSchoolName != nil && *regReq.School.ProposedSchoolName != "" {
				newSchool := models.MSchool{
					SchoolName: *regReq.School.ProposedSchoolName,
					IsActive:   true,
				}
				if err := tx.Create(&newSchool).Error; err != nil {
					return fmt.Errorf("gagal membuat data sekolah baru: %w", err)
				}
				schoolID = newSchool.ID
			} else if regReq.School != nil && regReq.School.ExistingSchoolID != nil {
				schoolID = *regReq.School.ExistingSchoolID
			}

			if schoolID != 0 {
				if err := tx.Model(&teacher).Updates(map[string]interface{}{
					"approval":  1,
					"school_id": schoolID,
				}).Error; err != nil {
					return err
				}
			} else {
				if err := tx.Model(&teacher).Update("approval", 1).Error; err != nil {
					return err
				}
			}

			if err := tx.Model(&regReq).Updates(map[string]interface{}{
				"status":      "approved",
				"reviewed_by": adminID,
			}).Error; err != nil {
				return err
			}

			if regReq.School != nil {
				if err := tx.Model(regReq.School).Update("status", "approved").Error; err != nil {
					return err
				}
			}
		} else if decision.Action == "reject" {
			if err := tx.Model(&teacher).Update("approval", 2).Error; err != nil {
				return err
			}

			if err := tx.Model(&regReq).Updates(map[string]interface{}{
				"status":         "rejected",
				"reviewed_by":    adminID,
				"rejection_note": decision.RejectionNote,
			}).Error; err != nil {
				return err
			}

			if regReq.School != nil {
				if err := tx.Model(regReq.School).Update("status", "rejected").Error; err != nil {
					return err
				}
			}
		} else {
			return fmt.Errorf("aksi keputusan tidak valid, harus 'approve' atau 'reject'")
		}

		return nil
	})
}

func (s *authService) Logout(refreshToken string) error {
	return s.refreshTokenRepo.DeleteByToken(refreshToken)
}

func (s *authService) GetProfile(token string) (user *models.User, err error) {
	claims, err := s.DecodeToken(token)
	if err != nil {
		return nil, err
	}
	userID, ok := claims["user_id"].(float64)

	if !ok {
		return nil, err
	}
	user, err = s.repo.WithPreloads("HasClass", "School").FindUserByID(int64(userID))
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *authService) GenerateAccessToken(user *models.User) (string, error) {
	claims := jwt.MapClaims{
		"user_id": user.ID,
		"name":    user.Name,
		"role":    string(user.Role),
		"exp":     time.Now().Add(7 * 24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

func (s *authService) GenerateRefreshToken(user *models.User) (string, error) {
	token := uuid.New().String()

	refresh := models.RefreshToken{
		UserID:    user.ID,
		Token:     token,
		ExpiresAt: time.Now().Add(7 * 24 * time.Hour),
		CreatedAt: time.Now(),
	}

	_, err := s.refreshTokenRepo.Insert(&refresh)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *authService) RefreshToken(token string) (string, error) {
	rt, err := s.refreshTokenRepo.FindByToken(token)
	if err != nil {
		return "", fmt.Errorf("refresh token tidak valid")
	}

	if time.Now().After(rt.ExpiresAt) {
		return "", fmt.Errorf("refresh token expired")
	}

	// 🔥 ambil user
	user, err := s.repo.FindUserByID(rt.UserID)
	if err != nil {
		return "", err
	}

	// 🔥 baru generate access token
	return s.GenerateAccessToken(user)
}

func (s *authService) DecodeToken(token string) (jwt.MapClaims, error) {
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	return claims, err
}

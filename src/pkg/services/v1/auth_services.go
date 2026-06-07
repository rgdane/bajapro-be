package services

import (
	"fmt"
	"log"
	"os"
	"time"

	"jk-api/api/http/controllers/v1/dto"
	"jk-api/internal/database/models"
	"jk-api/internal/errors/bcrypt_err"
	"jk-api/internal/errors/gorm_err"
	"jk-api/pkg/repository/adapter/sql"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
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
}

type authService struct {
	repo sql.UserRepository
	refreshTokenRepo  sql.RefreshTokenRepository
}

func NewAuthService(
	userRepo sql.UserRepository,
	refreshRepo sql.RefreshTokenRepository,
) *authService {
	return &authService{
		repo: userRepo,
		refreshTokenRepo: refreshRepo,
	}
}

func (s *authService) Login(email, password string) (*models.User, error) {
	user, err := s.repo.
		WithPreloads("HasRoles.HasPermissions").
		FindUserByEmail(email)

	if err != nil {
		return nil, gorm_err.TranslateGormError(err)
	}

	// 🔍 DEBUG (optional)
	log.Printf("User ID: %d, Roles count: %d", user.ID, len(user.HasRoles))

	// ❌ Cek user aktif
	if !user.IsActive {
		return nil, fmt.Errorf("akun tidak aktif")
	}

	// ❌ Cek approval teacher
	if user.HasRoles[0].ID == 2 && !user.IsApprovedByAdmin {
		return nil, fmt.Errorf("akun teacher belum di-approve admin")
	}

	// ❌ Cek password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, bcrypt_err.TranslateBcryptError(err)
	}

	return user, nil
}

func (s *authService) Register(req *dto.RegisterRequest) (*models.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	var roleID int64
	var isApproved bool

	switch req.Role {
	case "teacher":
		roleID = 2
		isApproved = false // ❗ harus approval admin
	case "student":
		roleID = 3
		isApproved = true
	default:
		return nil, fmt.Errorf("role tidak valid")
	}

	user := models.User{
		Name:              req.Name,
		Email:             req.Email,
		Password:          string(hashedPassword),
		RoleID:            roleID,
		IsApprovedByAdmin: isApproved,
		IsActive:          true,
	}

	createdUser, err := s.repo.InsertUser(&user)
	if err != nil {
		return nil, gorm_err.TranslateGormError(err)
	}

	return createdUser, nil
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
	user, err = s.repo.WithPreloads("HasRoles.HasPermissions", "HasClass").FindUserByID(int64(userID))
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *authService) GenerateAccessToken(user *models.User) (string, error) {
	var roles []string
	var permissions []string

	for _, role := range user.HasRoles {
		roles = append(roles, role.Name)

		for _, perm := range role.HasPermissions {
			permissions = append(permissions, perm.Name)
		}
	}

	claims := jwt.MapClaims{
		"user_id":    user.ID,
		"name":       user.Name,
		"roles":      roles,
		"permissions": permissions,
		"exp":        time.Now().Add(7 * 24 * time.Hour).Unix(),
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


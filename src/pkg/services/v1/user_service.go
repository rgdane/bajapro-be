package services

import (
	"fmt"
	"jk-api/api/http/controllers/v1/dto"
	"jk-api/internal/config"
	"jk-api/internal/database/models"
	"jk-api/internal/errors/gorm_err"
	"jk-api/pkg/repository/adapter/sql"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService interface {
	WithTx(tx *gorm.DB) UserService

	CreateUser(input *models.User) (*models.User, error)
	UpdateUser(id int64, updates map[string]interface{}, associations map[string]interface{}) (*models.User, error)
	DeleteUser(id int64, isPermanent bool) error
	GetAllUsers(filter dto.UserFilterDto) ([]models.User, error)
	GetUserByID(id int64, filter dto.UserFilterDto) (*models.User, error)
	GetDB() *gorm.DB
	BulkCreateUsers(data []*models.User) ([]*models.User, error)
	BulkUpdateUsers(ids []int64, updates map[string]interface{}, associatons map[string]interface{}) error
	BulkDeleteUsers(ids []int64, isPermanent bool) error
}

type userService struct {
	repo sql.UserRepository
	tx   *gorm.DB
}

func NewUserService(repo sql.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) WithTx(tx *gorm.DB) UserService {
	return &userService{
		repo: s.repo.WithTx(tx),
		tx:   tx,
	}
}

func (s *userService) GetDB() *gorm.DB {
	if s.tx != nil {
		return s.tx
	}
	return config.DB
}

func (s *userService) CreateUser(input *models.User) (*models.User, error) {
	if err := validateUserSchoolClass(input.Role, input.SchoolID, input.ClassID); err != nil {
		return nil, err
	}

	hashedPassword, err := HashPassword(input.Password)
	if err != nil {
		return nil, err
	}
	input.Password = hashedPassword
	input.CreatedAt = time.Now()
	input.UpdatedAt = time.Now()

	data, err := s.repo.InsertUser(input)
	if err != nil {
		return nil, gorm_err.TranslateGormError(err)
	}
	return data, nil
}

func validateUserSchoolClass(role models.UserRole, schoolID, classID *int64) error {
	switch role {
	case models.RoleTeacher:
		if schoolID == nil {
			return fmt.Errorf("teacher wajib memiliki school_id")
		}
	case models.RoleStudent:
		if schoolID == nil || classID == nil {
			return fmt.Errorf("student wajib memiliki school_id dan class_id")
		}
		var class models.MClass
		if err := config.DB.Where("id = ?", *classID).First(&class).Error; err != nil {
			return fmt.Errorf("kelas tidak ditemukan")
		}
		if class.SchoolID != *schoolID {
			return fmt.Errorf("kelas tidak berasal dari sekolah yang sama dengan student")
		}
		if !class.IsActive {
			return fmt.Errorf("kelas tidak aktif")
		}
	case models.RoleAdmin:
		if classID != nil {
			return fmt.Errorf("admin tidak boleh memiliki class_id")
		}
	}
	return nil
}

func (s *userService) UpdateUser(id int64, updates map[string]interface{}, associations map[string]interface{}) (*models.User, error) {
	repo := s.repo

	if len(associations) > 0 {
		assocNames := make([]string, 0, len(associations))
		for name := range associations {
			assocNames = append(assocNames, name)
			delete(updates, name)
		}
		repo = repo.WithAssociations(assocNames...).WithReplacements(associations)
	}

	if new_pwd, ok := updates["new_password"].(string); ok && new_pwd != "" {
		newHashedPassword, err := HashPassword(new_pwd)
		if err != nil {
			return nil, err
		}
		updates["new_password"] = newHashedPassword
	}

	if pwd, ok := updates["password"].(string); ok && pwd != "" {
		hashedPassword, err := HashPassword(pwd)
		if err != nil {
			return nil, err
		}
		updates["password"] = hashedPassword
	}

	_, hasRole := updates["role"]
	_, hasSchool := updates["school_id"]
	_, hasClass := updates["class_id"]
	if hasRole || hasSchool || hasClass {
		existing, err := repo.FindUserByID(id)
		if err != nil {
			return nil, gorm_err.TranslateGormError(err)
		}

		effectiveRole := existing.Role
		effectiveSchoolID := existing.SchoolID
		effectiveClassID := existing.ClassID

		if v, ok := updates["role"].(string); ok && v != "" {
			effectiveRole = models.UserRole(v)
		}
		if v, ok := updates["school_id"].(int64); ok {
			effectiveSchoolID = &v
		}
		if v, ok := updates["school_id"].(*int64); ok {
			effectiveSchoolID = v
		}
		if v, ok := updates["class_id"].(int64); ok {
			effectiveClassID = &v
		}
		if v, ok := updates["class_id"].(*int64); ok {
			effectiveClassID = v
		}

		if !effectiveRole.Valid() {
			return nil, fmt.Errorf("role tidak valid")
		}

		// Role menentukan school/class yang wajib. Bersihkan yang tak relevan
		// agar perubahan role selalu konsisten (mis. student -> teacher).
		switch effectiveRole {
		case models.RoleAdmin:
			effectiveSchoolID = nil
			effectiveClassID = nil
		case models.RoleTeacher:
			effectiveClassID = nil
		}

		if err := validateUserSchoolClass(effectiveRole, effectiveSchoolID, effectiveClassID); err != nil {
			return nil, err
		}

		// Tulis balik nilai efektif agar penghapusan (NULL) ikut tersimpan.
		updates["role"] = string(effectiveRole)
		if effectiveSchoolID != nil {
			updates["school_id"] = *effectiveSchoolID
		} else {
			updates["school_id"] = nil
		}
		if effectiveClassID != nil {
			updates["class_id"] = *effectiveClassID
		} else {
			updates["class_id"] = nil
		}
	}

	data, err := repo.UpdateUser(id, updates)
	if err != nil {
		return nil, gorm_err.TranslateGormError(err)
	}
	return data, nil
}

func (s *userService) DeleteUser(id int64, isPermanent bool) (err error) {
	repo := s.repo
	if isPermanent {
		repo = repo.WithUnscoped()
		err = repo.RemoveUser(id)
		return gorm_err.TranslateGormError(err)
	}

	payload := map[string]interface{}{
		"code": fmt.Sprintf("user-deleted-%d", time.Now().Unix()),
	}

	if _, err = s.UpdateUser(id, payload, nil); err != nil {
		return gorm_err.TranslateGormError(err)
	}

	err = repo.RemoveUser(id)
	return gorm_err.TranslateGormError(err)
}

func (s *userService) GetAllUsers(filter dto.UserFilterDto) ([]models.User, error) {
	repo := s.repo
	if filter.Role != "" {
		repo = repo.WithWhere("users.role = ?", filter.Role)
	}

	if filter.Preload {
		repo = repo.WithPreloads("HasClass", "School")
	}
	if filter.Limit != 0 {
		repo = repo.WithLimit(int(filter.Limit))
	}
	if filter.Cursor != 0 {
		repo = repo.WithCursor(int(filter.Cursor))
	}
	if filter.Name != "" {
		repo = repo.WithWhere("users.name ILIKE ?", "%"+filter.Name+"%")
	}
	if filter.Sort != "" && filter.Order != "" {
		orderClause := filter.Sort + " " + filter.Order
		repo = repo.WithOrder(orderClause)
	}
	if filter.ShowDeleted {
		repo = repo.WithUnscoped().WithWhere("deleted_at IS NOT NULL")
	}

	data, err := repo.FindUser()
	if err != nil {
		return nil, gorm_err.TranslateGormError(err)
	}
	return data, nil
}

func (s *userService) GetUserByID(id int64, filter dto.UserFilterDto) (*models.User, error) {
	repo := s.repo
	if filter.Preload {
		repo = repo.WithPreloads("HasClass", "School")
	}
	data, err := repo.FindUserByID(id)
	if err != nil {
		return nil, gorm_err.TranslateGormError(err)
	}
	return data, nil
}

func HashPassword(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashed), nil
}

func (s *userService) BulkCreateUsers(data []*models.User) ([]*models.User, error) {
	for i, user := range data {
		if user.Password == "" {
			return nil, fmt.Errorf("password untuk user index %d tidak boleh kosong", i)
		}

		hashedPassword, err := HashPassword(user.Password)
		if err != nil {
			return nil, fmt.Errorf("gagal hash password untuk user index %d: %w", i, err)
		}

		user.Password = hashedPassword
		user.CreatedAt = time.Now()
		user.UpdatedAt = time.Now()

	}

	datas, err := s.repo.InsertManyUsers(data)
	if err != nil {
		return nil, gorm_err.TranslateGormError(err)
	}
	return datas, nil
}

func (s *userService) BulkUpdateUsers(
	ids []int64,
	updates map[string]interface{},
	associations map[string]interface{},
) error {
	repo := s.repo
	if _, ok := updates["deleted_at"]; ok {
		repo = repo.WithUnscoped()
		deletedAtValue := updates["deleted_at"]
		var isNil bool
		if deletedAtValue == nil {
			isNil = true
		} else {
			switch v := deletedAtValue.(type) {
			case *time.Time:
				isNil = (v == nil)
			default:
				isNil = false
			}
		}

		if isNil {
			if err := s.userRestore(ids); err != nil {
				return err
			}
		}
	}

	if len(associations) > 0 {
		assocNames := make([]string, 0, len(associations))
		for name := range associations {
			assocNames = append(assocNames, name)
			delete(updates, name)
		}

		return s.GetDB().Transaction(func(tx *gorm.DB) error {
			for _, id := range ids {
				if len(updates) > 0 {
					if err := tx.Model(&models.User{}).
						Where("id = ?", id).
						Updates(updates).Error; err != nil {
						return gorm_err.TranslateGormError(err)
					}
				}

				for name, value := range associations {
					if err := tx.Model(&models.User{ID: id}).
						Association(name).
						Replace(value); err != nil {
						return gorm_err.TranslateGormError(err)
					}
				}
			}
			return nil
		})
	}

	err := repo.UpdateManyUsers(ids, updates)
	return gorm_err.TranslateGormError(err)
}

func (s *userService) BulkDeleteUsers(ids []int64, isPermanent bool) (err error) {
	repo := s.repo

	if isPermanent {
		fmt.Println("PERMANENT DELETE ACTIVE!")
		repo = repo.WithUnscoped()
		err = repo.RemoveManyUsers(ids)
		return gorm_err.TranslateGormError(err)
	}

	payload := map[string]interface{}{
		"code": fmt.Sprintf("user-deleted-%d", time.Now().Unix()),
	}

	if err = s.BulkUpdateUsers(ids, payload, nil); err != nil {
		return err
	}

	err = repo.RemoveManyUsers(ids)
	return gorm_err.TranslateGormError(err)
}

func (s *userService) userRestore(ids []int64) error {
	var users []models.User
	if err := s.GetDB().Unscoped().Where("id IN ? AND deleted_at IS NOT NULL", ids).Find(&users).Error; err != nil {
		return err
	}

	for _, user := range users {
		newCode := user.GenerateUserCode()
		if newCode != "" {
			if err := s.GetDB().Unscoped().Model(&user).Where("id = ?", user.ID).Update("code", newCode).Error; err != nil {
				return err
			}
		}
	}
	return nil
}

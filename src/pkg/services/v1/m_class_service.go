package services

import (
	"errors"
	"fmt"
	"jk-api/api/http/controllers/v1/dto"
	"jk-api/internal/config"
	"jk-api/internal/database/models"
	"jk-api/internal/errors/gorm_err"
	"jk-api/pkg/repository/adapter/sql"

	"gorm.io/gorm"
)

type MClassService interface {
	WithTx(tx *gorm.DB) MClassService

	CreateMClass(input *models.MClass) (*models.MClass, error)
	UpdateMClass(id int64, updates map[string]interface{}, associations map[string]interface{}) (*models.MClass, error)
	DeleteMClass(id int64) error
	GetAllMClasses(filter dto.MClassFilterDto) ([]models.MClass, error)
	GetMClassByID(id int64, filter dto.MClassFilterDto) (*models.MClass, error)
	GetDB() *gorm.DB
}

type mClassService struct {
	repo sql.MClassRepository
	tx   *gorm.DB
}

func NewMClassService(repo sql.MClassRepository) MClassService {
	return &mClassService{repo: repo}
}

func (s *mClassService) WithTx(tx *gorm.DB) MClassService {
	return &mClassService{
		repo: s.repo.WithTx(tx),
		tx:   tx,
	}
}

func (s *mClassService) GetDB() *gorm.DB {
	if s.tx != nil {
		return s.tx
	}
	return config.DB
}

func (s *mClassService) CreateMClass(input *models.MClass) (*models.MClass, error) {
	// Business Logic Validations
	var school models.MSchool
	if err := config.DB.Where("id = ? AND isactive = true", input.SchoolID).First(&school).Error; err != nil {
		return nil, fmt.Errorf("sekolah dengan id %d tidak ditemukan atau tidak aktif", input.SchoolID)
	}

	if input.TeacherID != nil {
		var teacher models.User
		if err := config.DB.Where("id = ? AND role = ? AND isactive = true", *input.TeacherID, models.RoleTeacher).First(&teacher).Error; err != nil {
			return nil, fmt.Errorf("pengajar dengan id %d tidak ditemukan, tidak aktif, atau bukan role teacher", *input.TeacherID)
		}

		if teacher.SchoolID == nil || *teacher.SchoolID != input.SchoolID {
			return nil, fmt.Errorf("pengajar tidak berasal dari sekolah yang sama dengan kelas")
		}
	}

	data, err := s.repo.InsertMClass(input)
	if err != nil {
		return nil, gorm_err.TranslateGormError(err)
	}
	return data, nil
}

func (s *mClassService) UpdateMClass(
	id int64,
	payload map[string]interface{},
	associations map[string]interface{},
) (*models.MClass, error) {

	repo := s.repo

	if len(associations) > 0 {
		var assocNames []string
		for name := range associations {
			assocNames = append(assocNames, name)
		}
		repo = repo.WithAssociations(assocNames...).WithReplacements(associations)
	}

	for key := range associations {
		delete(payload, key)
	}

	existing, err := repo.FindMClassByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("kelas tidak ditemukan")
		}
		return nil, gorm_err.TranslateGormError(err)
	}

	schoolID := existing.SchoolID
	if v, ok := payload["school_id"].(int64); ok {
		schoolID = v
	}

	var school models.MSchool
	if err := config.DB.Where("id = ? AND isactive = true", schoolID).First(&school).Error; err != nil {
		return nil, fmt.Errorf("sekolah dengan id %d tidak ditemukan atau tidak aktif", schoolID)
	}

	teacherID := existing.TeacherID
	if v, ok := payload["teacher_id"].(int64); ok {
		teacherID = &v
	}

	if teacherID != nil {
		var teacher models.User
		if err := config.DB.Where("id = ? AND role = ? AND isactive = true", *teacherID, models.RoleTeacher).First(&teacher).Error; err != nil {
			return nil, fmt.Errorf("pengajar dengan id %d tidak ditemukan, tidak aktif, atau bukan role teacher", *teacherID)
		}
		if teacher.SchoolID == nil || *teacher.SchoolID != schoolID {
			return nil, fmt.Errorf("pengajar tidak berasal dari sekolah yang sama dengan kelas")
		}
	}

	updated, err := repo.UpdateMClass(id, payload)
	if err != nil {
		return nil, gorm_err.TranslateGormError(err)
	}

	return updated, nil
}

func (s *mClassService) DeleteMClass(id int64) error {
	err := s.repo.WithAssociations("Students").RemoveMClass(id)
	return gorm_err.TranslateGormError(err)
}

func (s *mClassService) GetAllMClasses(filter dto.MClassFilterDto) ([]models.MClass, error) {
	repo := s.repo
	if filter.SchoolID != nil {
		repo = repo.WithWhere("school_id = ?", *filter.SchoolID)
	}
	if filter.TeacherID != nil {
		repo = repo.WithWhere("teacher_id = ?", *filter.TeacherID)
	}
	if filter.Preload {
		repo = repo.WithPreloads("School", "Teacher", "Students")
	}
	data, err := repo.FindMClass()
	if err != nil {
		return nil, gorm_err.TranslateGormError(err)
	}
	return data, nil
}

func (s *mClassService) GetMClassByID(id int64, filter dto.MClassFilterDto) (*models.MClass, error) {
	repo := s.repo
	if filter.Preload {
		repo = repo.WithPreloads("School", "Teacher", "Students")
	}
	data, err := repo.FindMClassByID(id)
	if err != nil {
		return nil, gorm_err.TranslateGormError(err)
	}
	return data, nil
}

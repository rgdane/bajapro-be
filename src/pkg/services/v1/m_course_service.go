package services

import (
	"jk-api/api/http/controllers/v1/dto"
	"jk-api/internal/config"
	"jk-api/internal/database/models"
	"jk-api/internal/errors/gorm_err"
	"jk-api/pkg/repository/adapter/sql"

	"gorm.io/gorm"
)

type MCourseService interface {
	WithTx(tx *gorm.DB) MCourseService

	CreateMCourse(input *models.MCourse) (*models.MCourse, error)
	UpdateMCourse(id int64, updates map[string]interface{}, associations map[string]interface{}) (*models.MCourse, error)
	DeleteMCourse(id int64) error
	GetAllMCourses(filter dto.MCourseFilterDto) ([]models.MCourse, error)
	GetMCourseByID(id int64, filter dto.MCourseFilterDto) (*models.MCourse, error)
	GetDB() *gorm.DB
}

type mCourseService struct {
	repo sql.MCourseRepository
	tx   *gorm.DB
}

func NewMCourseService(repo sql.MCourseRepository) MCourseService {
	return &mCourseService{repo: repo}
}

func (s *mCourseService) WithTx(tx *gorm.DB) MCourseService {
	return &mCourseService{
		repo: s.repo.WithTx(tx),
		tx:   tx,
	}
}

func (s *mCourseService) GetDB() *gorm.DB {
	if s.tx != nil {
		return s.tx
	}
	return config.DB
}

func (s *mCourseService) CreateMCourse(input *models.MCourse) (*models.MCourse, error) {
	data, err := s.repo.InsertMCourse(input)
	if err != nil {
		return nil, gorm_err.TranslateGormError(err)
	}
	return data, nil
}

func (s *mCourseService) UpdateMCourse(
	id int64,
	payload map[string]interface{},
	associations map[string]interface{},
) (*models.MCourse, error) {

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

	updated, err := repo.UpdateMCourse(id, payload)
	if err != nil {
		return nil, gorm_err.TranslateGormError(err)
	}

	return updated, nil
}

func (s *mCourseService) DeleteMCourse(id int64) error {
	err := s.repo.RemoveMCourse(id)
	return gorm_err.TranslateGormError(err)
}

func (s *mCourseService) GetAllMCourses(filter dto.MCourseFilterDto) ([]models.MCourse, error) {
	repo := s.repo
	if filter.Preload {
		repo = repo.WithPreloads("Lessons")
	}
	data, err := repo.FindMCourse()
	if err != nil {
		return nil, gorm_err.TranslateGormError(err)
	}
	return data, nil
}

func (s *mCourseService) GetMCourseByID(id int64, filter dto.MCourseFilterDto) (*models.MCourse, error) {
	repo := s.repo
	if filter.Preload {
		repo = repo.WithPreloads("Lessons")
	}
	data, err := repo.FindMCourseByID(id)
	if err != nil {
		return nil, gorm_err.TranslateGormError(err)
	}
	return data, nil
}

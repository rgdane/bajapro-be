package services

import (
	"jk-api/api/http/controllers/v1/dto"
	"jk-api/internal/config"
	"jk-api/internal/database/models"
	"jk-api/internal/errors/gorm_err"
	"jk-api/pkg/repository/adapter/sql"

	"gorm.io/gorm"
)

type MSchoolService interface {
	WithTx(tx *gorm.DB) MSchoolService
	CreateMSchool(input *models.MSchool) (*models.MSchool, error)
	UpdateMSchool(id int64, updates map[string]interface{}) (*models.MSchool, error)
	DeleteMSchool(id int64) error
	GetAllMSchools(filter dto.MSchoolFilterDto) ([]models.MSchool, error)
	GetMSchoolByID(id int64, filter dto.MSchoolFilterDto) (*models.MSchool, error)
	GetDB() *gorm.DB
}

type mSchoolService struct {
	repo sql.MSchoolRepository
	tx   *gorm.DB
}

func NewMSchoolService(repo sql.MSchoolRepository) MSchoolService {
	return &mSchoolService{repo: repo}
}

func (s *mSchoolService) WithTx(tx *gorm.DB) MSchoolService {
	return &mSchoolService{
		repo: s.repo.WithTx(tx),
		tx:   tx,
	}
}

func (s *mSchoolService) GetDB() *gorm.DB {
	if s.tx != nil {
		return s.tx
	}
	return config.DB
}

func (s *mSchoolService) CreateMSchool(input *models.MSchool) (*models.MSchool, error) {
	data, err := s.repo.InsertMSchool(input)
	if err != nil {
		return nil, gorm_err.TranslateGormError(err)
	}
	return data, nil
}

func (s *mSchoolService) UpdateMSchool(id int64, updates map[string]interface{}) (*models.MSchool, error) {
	updated, err := s.repo.UpdateMSchool(id, updates)
	if err != nil {
		return nil, gorm_err.TranslateGormError(err)
	}
	return updated, nil
}

func (s *mSchoolService) DeleteMSchool(id int64) error {
	err := s.repo.RemoveMSchool(id)
	return gorm_err.TranslateGormError(err)
}

func (s *mSchoolService) GetAllMSchools(filter dto.MSchoolFilterDto) ([]models.MSchool, error) {
	repo := s.repo
	if filter.Name != "" {
		repo = repo.WithWhere("school_name ILIKE ?", "%"+filter.Name+"%")
	}
	if filter.Preload {
		repo = repo.WithPreloads("Classes", "Users")
	}
	data, err := repo.FindMSchool()
	if err != nil {
		return nil, gorm_err.TranslateGormError(err)
	}
	return data, nil
}

func (s *mSchoolService) GetMSchoolByID(id int64, filter dto.MSchoolFilterDto) (*models.MSchool, error) {
	repo := s.repo
	if filter.Preload {
		repo = repo.WithPreloads("Classes", "Users")
	}
	data, err := repo.FindMSchoolByID(id)
	if err != nil {
		return nil, gorm_err.TranslateGormError(err)
	}
	return data, nil
}

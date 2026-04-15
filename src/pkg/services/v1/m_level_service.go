package services

import (
	"jk-api/api/http/controllers/v1/dto"
	"jk-api/internal/config"
	"jk-api/internal/database/models"
	"jk-api/internal/errors/gorm_err"
	"jk-api/pkg/repository/adapter/sql"

	"gorm.io/gorm"
)

type MLevelService interface {
	WithTx(tx *gorm.DB) MLevelService

	CreateMLevel(input *models.MLevel) (*models.MLevel, error)
	UpdateMLevel(id int64, updates map[string]interface{}, associations map[string]interface{}) (*models.MLevel, error)
	DeleteMLevel(id int64) error
	GetAllMLevels(filter dto.MLevelFilterDto) ([]models.MLevel, error)
	GetMLevelByID(id int64, filter dto.MLevelFilterDto) (*models.MLevel, error)
	GetDB() *gorm.DB
}

type mLevelService struct {
	repo sql.MLevelRepository
	tx   *gorm.DB
}

func NewMLevelService(repo sql.MLevelRepository) MLevelService {
	return &mLevelService{repo: repo}
}

func (s *mLevelService) WithTx(tx *gorm.DB) MLevelService {
	return &mLevelService{
		repo: s.repo.WithTx(tx),
		tx:   tx,
	}
}

func (s *mLevelService) GetDB() *gorm.DB {
	if s.tx != nil {
		return s.tx
	}
	return config.DB
}

func (s *mLevelService) CreateMLevel(input *models.MLevel) (*models.MLevel, error) {
	data, err := s.repo.InsertMLevel(input)
	if err != nil {
		return nil, gorm_err.TranslateGormError(err)
	}
	return data, nil
}

func (s *mLevelService) UpdateMLevel(
	id int64,
	payload map[string]interface{},
	associations map[string]interface{},
) (*models.MLevel, error) {

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

	updated, err := repo.UpdateMLevel(id, payload)
	if err != nil {
		return nil, gorm_err.TranslateGormError(err)
	}

	return updated, nil
}

func (s *mLevelService) DeleteMLevel(id int64) error {
	err := s.repo.RemoveMLevel(id)
	return gorm_err.TranslateGormError(err)
}

func (s *mLevelService) GetAllMLevels(filter dto.MLevelFilterDto) ([]models.MLevel, error) {
	repo := s.repo
	// if filter.Preload {
	// 	repo = repo.WithPreloads("Teachers", "Students")
	// }
	data, err := repo.FindMLevel()
	if err != nil {
		return nil, gorm_err.TranslateGormError(err)
	}
	return data, nil
}

func (s *mLevelService) GetMLevelByID(id int64, filter dto.MLevelFilterDto) (*models.MLevel, error) {
	repo := s.repo
	// if filter.Preload {
	// 	repo = repo.WithPreloads("Teachers", "Students")
	// }
	data, err := repo.FindMLevelByID(id)
	if err != nil {
		return nil, gorm_err.TranslateGormError(err)
	}
	return data, nil
}

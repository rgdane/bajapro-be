package services

import (
	"jk-api/api/http/controllers/v1/dto"
	"jk-api/internal/config"
	"jk-api/internal/database/models"
	"jk-api/internal/errors/gorm_err"
	"jk-api/pkg/repository/adapter/sql"

	"gorm.io/gorm"
)

type MBadgeSettingsService interface {
	WithTx(tx *gorm.DB) MBadgeSettingsService

	CreateMBadgeSettings(input *models.MBadgeSettings) (*models.MBadgeSettings, error)
	UpdateMBadgeSettings(id int64, updates map[string]interface{}, associations map[string]interface{}) (*models.MBadgeSettings, error)
	DeleteMBadgeSettings(id int64) error
	GetAllMBadgeSettings(filter dto.MBadgeSettingsFilterDto) ([]models.MBadgeSettings, error)
	GetMBadgeSettingsByID(id int64, filter dto.MBadgeSettingsFilterDto) (*models.MBadgeSettings, error)
	GetDB() *gorm.DB
}

type mBadgeSettingsService struct {
	repo sql.MBadgeSettingsRepository
	tx   *gorm.DB
}

func NewMBadgeSettingsService(repo sql.MBadgeSettingsRepository) MBadgeSettingsService {
	return &mBadgeSettingsService{repo: repo}
}

func (s *mBadgeSettingsService) WithTx(tx *gorm.DB) MBadgeSettingsService {
	return &mBadgeSettingsService{
		repo: s.repo.WithTx(tx),
		tx:   tx,
	}
}

func (s *mBadgeSettingsService) GetDB() *gorm.DB {
	if s.tx != nil {
		return s.tx
	}
	return config.DB
}

func (s *mBadgeSettingsService) CreateMBadgeSettings(input *models.MBadgeSettings) (*models.MBadgeSettings, error) {
	data, err := s.repo.InsertMBadgeSettings(input)
	if err != nil {
		return nil, gorm_err.TranslateGormError(err)
	}
	return data, nil
}

func (s *mBadgeSettingsService) UpdateMBadgeSettings(
	id int64,
	payload map[string]interface{},
	associations map[string]interface{},
) (*models.MBadgeSettings, error) {

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

	updated, err := repo.UpdateMBadgeSettings(id, payload)
	if err != nil {
		return nil, gorm_err.TranslateGormError(err)
	}

	return updated, nil
}

func (s *mBadgeSettingsService) DeleteMBadgeSettings(id int64) error {
	err := s.repo.RemoveMBadgeSettings(id)
	return gorm_err.TranslateGormError(err)
}

func (s *mBadgeSettingsService) GetAllMBadgeSettings(filter dto.MBadgeSettingsFilterDto) ([]models.MBadgeSettings, error) {
	repo := s.repo
	// if filter.Preload {
	// 	repo = repo.WithPreloads("Teachers", "Students")
	// }
	data, err := repo.FindMBadgeSettings()
	if err != nil {
		return nil, gorm_err.TranslateGormError(err)
	}
	return data, nil
}

func (s *mBadgeSettingsService) GetMBadgeSettingsByID(id int64, filter dto.MBadgeSettingsFilterDto) (*models.MBadgeSettings, error) {
	repo := s.repo
	// if filter.Preload {
	// 	repo = repo.WithPreloads("Teachers", "Students")
	// }
	data, err := repo.FindMBadgeSettingsByID(id)
	if err != nil {
		return nil, gorm_err.TranslateGormError(err)
	}
	return data, nil
}

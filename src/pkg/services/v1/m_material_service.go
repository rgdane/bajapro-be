package services

import (
	"fmt"
	"jk-api/api/http/controllers/v1/dto"
	"jk-api/internal/config"
	"jk-api/internal/database/models"
	"jk-api/internal/errors/gorm_err"
	"jk-api/pkg/repository/adapter/sql"
	"time"

	"gorm.io/gorm"
)

type MMaterialService interface {
	WithTx(tx *gorm.DB) MMaterialService

	CreateMMaterial(input *models.MMaterials) (*models.MMaterials, error)
	UpdateMMaterial(id int64, updates map[string]interface{}) (*models.MMaterials, error)
	DeleteMMaterial(id int64) error
	GetAllMMaterials(filter dto.MMaterialFilterDto) ([]models.MMaterials, error)
	GetMMaterialByID(id int64, filter dto.MMaterialFilterDto) (*models.MMaterials, error)
	GetMMaterialsByIDs(ids []int64) ([]*models.MMaterials, error)
	GetDB() *gorm.DB
	BulkCreateMMaterials(data []*models.MMaterials) ([]*models.MMaterials, error)
	BulkUpdateMMaterials(ids []int64, updates map[string]interface{}) error
	BulkDeleteMMaterials(ids []int64) error
}

type mMaterialService struct {
	repo sql.MMaterialRepository
	tx   *gorm.DB
}

func NewMMaterialService(repo sql.MMaterialRepository) MMaterialService {
	return &mMaterialService{repo: repo}
}

func (s *mMaterialService) WithTx(tx *gorm.DB) MMaterialService {
	return &mMaterialService{
		repo: s.repo.WithTx(tx),
		tx:   tx,
	}
}

func (s *mMaterialService) GetDB() *gorm.DB {
	if s.tx != nil {
		return s.tx
	}
	return config.DB
}

func (s *mMaterialService) CreateMMaterial(input *models.MMaterials) (*models.MMaterials, error) {
	input.CreatedAt = time.Now()
	input.UpdatedAt = time.Now()

	data, err := s.repo.InsertMMaterial(input)
	if err != nil {
		return nil, gorm_err.TranslateGormError(err)
	}
	return data, nil
}

func (s *mMaterialService) UpdateMMaterial(id int64, updates map[string]interface{}) (*models.MMaterials, error) {
	repo := s.repo

	if _, ok := updates["deleted_at"]; ok {
		repo = repo.WithUnscoped()
	}

	if _, err := repo.FindMMaterialByID(id); err != nil {
		return nil, gorm_err.TranslateGormError(err)
	}

	updates["updated_at"] = time.Now()
	fmt.Println(updates)

	data, err := s.repo.UpdateMMaterial(id, updates)
	if err != nil {
		return nil, gorm_err.TranslateGormError(err)
	}
	return data, nil
}

func (s *mMaterialService) DeleteMMaterial(id int64) error {
	err := s.repo.RemoveMMaterial(id)
	return gorm_err.TranslateGormError(err)
}

func (s *mMaterialService) GetAllMMaterials(filter dto.MMaterialFilterDto) ([]models.MMaterials, error) {
	repo := s.repo

	if filter.Preload {
		repo = repo.WithPreloads("SubLesson")
	}

	if filter.Sort != "" && filter.Order != "" {
		orderClause := filter.Sort + " " + filter.Order
		repo = repo.WithOrder(orderClause)
	}

	if filter.ShowDeleted {
		repo = repo.WithUnscoped().WithWhere("m_materials.deleted_at IS NOT NULL")
	}
	fmt.Println("filter", filter)

	data, err := repo.FindMMaterials()
	if err != nil {
		return nil, gorm_err.TranslateGormError(err)
	}
	return data, nil
}

func (s *mMaterialService) GetMMaterialByID(id int64, filter dto.MMaterialFilterDto) (*models.MMaterials, error) {
	repo := s.repo
	if filter.Preload {
		repo = repo.WithPreloads("SubLesson")
	}
	data, err := repo.FindMMaterialByID(id)
	if err != nil {
		return nil, gorm_err.TranslateGormError(err)
	}
	return data, nil
}

func (s *mMaterialService) BulkCreateMMaterials(data []*models.MMaterials) ([]*models.MMaterials, error) {
	datas, err := s.repo.InsertManyMMaterials(data)
	if err != nil {
		return nil, gorm_err.TranslateGormError(err)
	}
	return datas, nil
}

func (s *mMaterialService) BulkUpdateMMaterials(ids []int64, updates map[string]interface{}) error {
	repo := s.repo

	if _, ok := updates["deleted_at"]; ok {
		repo = s.repo.WithUnscoped()
	}

	err := repo.UpdateManyMMaterials(ids, updates)
	return gorm_err.TranslateGormError(err)
}

func (s *mMaterialService) GetMMaterialsByIDs(ids []int64) ([]*models.MMaterials, error) {
	data, err := s.repo.FindMMaterialsByIDs(ids)
	if err != nil {
		return nil, gorm_err.TranslateGormError(err)
	}
	return data, nil
}

func (s *mMaterialService) BulkDeleteMMaterials(ids []int64) error {
	err := s.repo.RemoveManyMMaterials(ids)
	return gorm_err.TranslateGormError(err)
}
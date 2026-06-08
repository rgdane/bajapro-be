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

type MLessonService interface {
	WithTx(tx *gorm.DB) MLessonService

	CreateMLesson(input *models.MLesson) (*models.MLesson, error)
	UpdateMLesson(id int64, updates map[string]interface{}) (*models.MLesson, error)
	DeleteMLesson(id int64) error
	GetAllMLessons(filter dto.MLessonFilterDto) ([]models.MLesson, error)
	GetMLessonByID(id int64, filter dto.MLessonFilterDto) (*models.MLesson, error)
	GetMLessonsByIDs(ids []int64) ([]*models.MLesson, error)
	GetDB() *gorm.DB
	BulkCreateMLessons(data []*models.MLesson) ([]*models.MLesson, error)
	BulkUpdateMLessons(ids []int64, updates map[string]interface{}) error
	BulkDeleteMLessons(ids []int64) error
}

type mLessonService struct {
	repo sql.MLessonRepository
	tx   *gorm.DB
}

func NewMLessonService(repo sql.MLessonRepository) MLessonService {
	return &mLessonService{repo: repo}
}

func (s *mLessonService) WithTx(tx *gorm.DB) MLessonService {
	return &mLessonService{
		repo: s.repo.WithTx(tx),
		tx:   tx,
	}
}

func (s *mLessonService) GetDB() *gorm.DB {
	if s.tx != nil {
		return s.tx
	}
	return config.DB
}

func (s *mLessonService) CreateMLesson(input *models.MLesson) (*models.MLesson, error) {
	input.CreatedAt = time.Now()
	input.UpdatedAt = time.Now()

	data, err := s.repo.InsertMLesson(input)
	if err != nil {
		return nil, gorm_err.TranslateGormError(err)
	}
	return data, nil
}

func (s *mLessonService) UpdateMLesson(id int64, updates map[string]interface{}) (*models.MLesson, error) {
	repo := s.repo

	if _, ok := updates["deleted_at"]; ok {
		repo = repo.WithUnscoped()
	}

	if _, err := repo.FindMLessonByID(id); err != nil {
		return nil, gorm_err.TranslateGormError(err)
	}

	updates["updated_at"] = time.Now()
	fmt.Println(updates)

	data, err := s.repo.UpdateMLesson(id, updates)
	if err != nil {
		return nil, gorm_err.TranslateGormError(err)
	}
	return data, nil
}

func (s *mLessonService) DeleteMLesson(id int64) error {
	err := s.repo.RemoveMLesson(id)
	return gorm_err.TranslateGormError(err)
}

func (s *mLessonService) GetAllMLessons(filter dto.MLessonFilterDto) ([]models.MLesson, error) {
	repo := s.repo

	if filter.Preload {
		repo = repo.WithPreloads("Course", "Level")
	}

	if filter.Sort != "" && filter.Order != "" {
		orderClause := filter.Sort + " " + filter.Order
		repo = repo.WithOrder(orderClause)
	}

	if filter.ShowDeleted {
		repo = repo.WithUnscoped().WithWhere("m_lessons.deleted_at IS NOT NULL")
	}
	fmt.Println("filter", filter)

	data, err := repo.FindMLesson()
	if err != nil {
		return nil, gorm_err.TranslateGormError(err)
	}
	return data, nil
}

func (s *mLessonService) GetMLessonByID(id int64, filter dto.MLessonFilterDto) (*models.MLesson, error) {
	repo := s.repo
	if filter.Preload {
		repo = repo.WithPreloads("Course", "Level")
	}
	data, err := repo.FindMLessonByID(id)
	if err != nil {
		return nil, gorm_err.TranslateGormError(err)
	}
	return data, nil
}

func (s *mLessonService) BulkCreateMLessons(data []*models.MLesson) ([]*models.MLesson, error) {
	datas, err := s.repo.InsertManyMLessons(data)
	if err != nil {
		return nil, gorm_err.TranslateGormError(err)
	}
	return datas, nil
}

func (s *mLessonService) BulkUpdateMLessons(ids []int64, updates map[string]interface{}) error {
	repo := s.repo

	if _, ok := updates["deleted_at"]; ok {
		repo = s.repo.WithUnscoped()
	}

	err := repo.UpdateManyMLessons(ids, updates)
	return gorm_err.TranslateGormError(err)
}

func (s *mLessonService) GetMLessonsByIDs(ids []int64) ([]*models.MLesson, error) {
	data, err := s.repo.FindMLessonsByIDs(ids)
	if err != nil {
		return nil, gorm_err.TranslateGormError(err)
	}
	return data, nil
}

func (s *mLessonService) BulkDeleteMLessons(ids []int64) error {
	err := s.repo.RemoveManyMLessons(ids)
	return gorm_err.TranslateGormError(err)
}
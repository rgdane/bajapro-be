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

type MSubLessonService interface {
	WithTx(tx *gorm.DB) MSubLessonService

	CreateMSubLesson(input *models.MSubLesson) (*models.MSubLesson, error)
	UpdateMSubLesson(id int64, updates map[string]interface{}) (*models.MSubLesson, error)
	DeleteMSubLesson(id int64) error
		GetAllMSubLessons(filter dto.MSubLessonFilterDto) ([]models.MSubLesson, error)
	GetMSubLessonByID(id int64, filter dto.MSubLessonFilterDto) (*models.MSubLesson, error)
	GetMSubLessonsByIDs(ids []int64) ([]*models.MSubLesson, error)
	GetDB() *gorm.DB
	BulkCreateMSubLessons(data []*models.MSubLesson) ([]*models.MSubLesson, error)
	BulkUpdateMSubLessons(ids []int64, updates map[string]interface{}) error
	BulkDeleteMSubLessons(ids []int64) error
}

type mSubLessonService struct {
	repo sql.MSubLessonRepository
	tx   *gorm.DB
}

func NewMSubLessonService(repo sql.MSubLessonRepository) MSubLessonService {
	return &mSubLessonService{repo: repo}
}

func (s *mSubLessonService) WithTx(tx *gorm.DB) MSubLessonService {
	return &mSubLessonService{
		repo: s.repo.WithTx(tx),
		tx:   tx,
	}
}

func (s *mSubLessonService) GetDB() *gorm.DB {
	if s.tx != nil {
		return s.tx
	}
	return config.DB
}

func (s *mSubLessonService) CreateMSubLesson(input *models.MSubLesson) (*models.MSubLesson, error) {
	input.CreatedAt = time.Now()
	input.UpdatedAt = time.Now()

	data, err := s.repo.InsertMSubLesson(input)
	if err != nil {
		return nil, gorm_err.TranslateGormError(err)
	}
	return data, nil
}

func (s *mSubLessonService) UpdateMSubLesson(id int64, updates map[string]interface{}) (*models.MSubLesson, error) {
	repo := s.repo

	if _, ok := updates["deleted_at"]; ok {
		repo = repo.WithUnscoped()
	}

	if _, err := repo.FindMSubLessonByID(id); err != nil {
		return nil, gorm_err.TranslateGormError(err)
	}

	updates["updated_at"] = time.Now()
	fmt.Println(updates)

	data, err := s.repo.UpdateMSubLesson(id, updates)
	if err != nil {
		return nil, gorm_err.TranslateGormError(err)
	}
	return data, nil
}

func (s *mSubLessonService) DeleteMSubLesson(id int64) error {
	err := s.repo.RemoveMSubLesson(id)
	return gorm_err.TranslateGormError(err)
}

func (s *mSubLessonService) GetAllMSubLessons(filter dto.MSubLessonFilterDto) ([]models.MSubLesson, error) {
	repo := s.repo

	if filter.Preload {
		repo = repo.WithPreloads("Lesson")
	}

	if filter.Sort != "" && filter.Order != "" {
		orderClause := filter.Sort + " " + filter.Order
		repo = repo.WithOrder(orderClause)
	}

	if filter.ShowDeleted {
		repo = repo.WithUnscoped().WithWhere("m_sub_lessons.deleted_at IS NOT NULL")
	}
	fmt.Println("filter", filter)

	data, err := repo.FindMSubLessons()
	if err != nil {
		return nil, gorm_err.TranslateGormError(err)
	}
	return data, nil
}

func (s *mSubLessonService) GetMSubLessonByID(id int64, filter dto.MSubLessonFilterDto) (*models.MSubLesson, error) {
	repo := s.repo
	if filter.Preload {
		repo = repo.WithPreloads("Lesson")
	}
	data, err := repo.FindMSubLessonByID(id)
	if err != nil {
		return nil, gorm_err.TranslateGormError(err)
	}
	return data, nil
}

func (s *mSubLessonService) BulkCreateMSubLessons(data []*models.MSubLesson) ([]*models.MSubLesson, error) {
	datas, err := s.repo.InsertManyMSubLessons(data)
	if err != nil {
		return nil, gorm_err.TranslateGormError(err)
	}
	return datas, nil
}

func (s *mSubLessonService) BulkUpdateMSubLessons(ids []int64, updates map[string]interface{}) error {
	repo := s.repo

	if _, ok := updates["deleted_at"]; ok {
		repo = s.repo.WithUnscoped()
	}

	err := repo.UpdateManyMSubLessons(ids, updates)
	return gorm_err.TranslateGormError(err)
}

func (s *mSubLessonService) GetMSubLessonsByIDs(ids []int64) ([]*models.MSubLesson, error) {
	data, err := s.repo.FindMSubLessonsByIDs(ids)
	if err != nil {
		return nil, gorm_err.TranslateGormError(err)
	}
	return data, nil
}

func (s *mSubLessonService) BulkDeleteMSubLessons(ids []int64) error {
	err := s.repo.RemoveManyMSubLessons(ids)
	return gorm_err.TranslateGormError(err)
}
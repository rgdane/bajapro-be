package services

import (
	"jk-api/internal/config"
	"jk-api/internal/database/models"
	"jk-api/internal/errors/gorm_err"
	"jk-api/pkg/repository/adapter/sql"

	"gorm.io/gorm"
)

type TCodeQuestionService interface {
	WithTx(tx *gorm.DB) TCodeQuestionService
	GetCodeQuestionsBySubLessonID(subLessonID int64) ([]models.TCodeQuestion, error)
	GetCodeQuestionByID(id int64) (*models.TCodeQuestion, error)
	CreateCodeQuestion(data *models.TCodeQuestion) (*models.TCodeQuestion, error)
	GetDB() *gorm.DB
}

type tCodeQuestionService struct {
	repo sql.TCodeQuestionRepository
	tx   *gorm.DB
}

func NewTCodeQuestionService(repo sql.TCodeQuestionRepository) TCodeQuestionService {
	return &tCodeQuestionService{repo: repo}
}

func (s *tCodeQuestionService) WithTx(tx *gorm.DB) TCodeQuestionService {
	return &tCodeQuestionService{
		repo: s.repo.WithTx(tx),
		tx:   tx,
	}
}

func (s *tCodeQuestionService) GetDB() *gorm.DB {
	if s.tx != nil {
		return s.tx
	}
	return config.DB
}

func (s *tCodeQuestionService) GetCodeQuestionsBySubLessonID(subLessonID int64) ([]models.TCodeQuestion, error) {
	data, err := s.repo.FindTCodeQuestionsBySubLessonID(subLessonID)
	if err != nil {
		return nil, gorm_err.TranslateGormError(err)
	}
	return data, nil
}

func (s *tCodeQuestionService) CreateCodeQuestion(data *models.TCodeQuestion) (*models.TCodeQuestion, error) {
	data, err := s.repo.CreateTCodeQuestion(data)
	if err != nil {
		return nil, gorm_err.TranslateGormError(err)
	}
	return data, nil
}

func (s *tCodeQuestionService) GetCodeQuestionByID(id int64) (*models.TCodeQuestion, error) {
	data, err := s.repo.FindTCodeQuestionByID(id)
	if err != nil {
		return nil, gorm_err.TranslateGormError(err)
	}
	return data, nil
}




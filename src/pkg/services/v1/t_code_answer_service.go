package services

import (
	"jk-api/internal/config"
	"jk-api/internal/database/models"
	"jk-api/internal/errors/gorm_err"
	"jk-api/pkg/repository/adapter/sql"

	"gorm.io/gorm"
)

type TCodeAnswerService interface {
	WithTx(tx *gorm.DB) TCodeAnswerService
	GetTCodeAnswersByCodeQuestionID(codeQuestionID int64) ([]models.TCodeAnswer, error)
	CreateTCodeAnswer(data *models.TCodeAnswer, userID int64) (*models.TCodeAnswer, error)
	GetDB() *gorm.DB
}

type tCodeAnswerService struct {
	repo sql.TCodeAnswerRepository
	tx   *gorm.DB
}

func NewTCodeAnswerService(repo sql.TCodeAnswerRepository) TCodeAnswerService {
	return &tCodeAnswerService{repo: repo}
}

func (s *tCodeAnswerService) WithTx(tx *gorm.DB) TCodeAnswerService {
	return &tCodeAnswerService{
		repo: s.repo.WithTx(tx),
		tx:   tx,
	}
}

func (s *tCodeAnswerService) GetDB() *gorm.DB {
	if s.tx != nil {
		return s.tx
	}
	return config.DB
}

func (s *tCodeAnswerService) GetTCodeAnswersByCodeQuestionID(codeQuestionID int64) ([]models.TCodeAnswer, error) {
	data, err := s.repo.FindTCodeAnswersByCodeQuestionID(codeQuestionID)
	if err != nil {
		return nil, gorm_err.TranslateGormError(err)
	}
	return data, nil
}

func (s *tCodeAnswerService) CreateTCodeAnswer(data *models.TCodeAnswer, userID int64) (*models.TCodeAnswer, error) {
	data, err := s.repo.CreateTCodeAnswer(data)
	if err != nil {
		return nil, gorm_err.TranslateGormError(err)
	}
	return data, nil
}





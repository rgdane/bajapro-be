package services

import (
	"jk-api/internal/config"
	"jk-api/internal/database/models"
	"jk-api/internal/errors/gorm_err"
	"jk-api/pkg/repository/adapter/sql"

	"gorm.io/gorm"
)

type TEssayAnswerService interface {
	WithTx(tx *gorm.DB) TEssayAnswerService
	GetTEssayAnswersByEssayQuestionIDAndUserID(essayQuestionID, userID int64) (*models.TEssayAnswer, error)
	CreateTEssayAnswer(data *models.TEssayAnswer, userID int64) (*models.TEssayAnswer, error)
	GetDB() *gorm.DB
}

type tEssayAnswerService struct {
	repo sql.TEssayAnswerRepository
	tx   *gorm.DB
}

func NewTEssayAnswerService(repo sql.TEssayAnswerRepository) TEssayAnswerService {
	return &tEssayAnswerService{repo: repo}
}

func (s *tEssayAnswerService) WithTx(tx *gorm.DB) TEssayAnswerService {
	return &tEssayAnswerService{
		repo: s.repo.WithTx(tx),
		tx:   tx,
	}
}

func (s *tEssayAnswerService) GetDB() *gorm.DB {
	if s.tx != nil {
		return s.tx
	}
	return config.DB
}

func (s *tEssayAnswerService) GetTEssayAnswersByEssayQuestionIDAndUserID(essayQuestionID, userID int64) (*models.TEssayAnswer, error) {
	data, err := s.repo.FindTEssayAnswersByEssayQuestionIDAndUserID(essayQuestionID, userID)
	if err != nil {
		return nil, gorm_err.TranslateGormError(err)
	}
	return data, nil
}

func (s *tEssayAnswerService) CreateTEssayAnswer(data *models.TEssayAnswer, userID int64) (*models.TEssayAnswer, error) {
	data, err := s.repo.CreateTEssayAnswer(data)
	if err != nil {
		return nil, gorm_err.TranslateGormError(err)
	}
	return data, nil
}





package services

import (
	"jk-api/internal/config"
	"jk-api/internal/database/models"
	"jk-api/internal/errors/gorm_err"
	"jk-api/pkg/repository/adapter/sql"

	"gorm.io/gorm"
)

type EssayQuestionService interface {
	WithTx(tx *gorm.DB) EssayQuestionService
	GetEssayQuestionsByCodeQuestionID(codeQuestionID int64) ([]models.EssayQuestion, error)
	GetEssayQuestionByID(id int64) (*models.EssayQuestion, error)
	CreateEssayQuestion(data *models.EssayQuestion) (*models.EssayQuestion, error)
	GetDB() *gorm.DB
}

type essayQuestionService struct {
	repo sql.EssayQuestionRepository
	tx   *gorm.DB
}

func NewEssayQuestionService(repo sql.EssayQuestionRepository) EssayQuestionService {
	return &essayQuestionService{repo: repo}
}

func (s *essayQuestionService) WithTx(tx *gorm.DB) EssayQuestionService {
	return &essayQuestionService{
		repo: s.repo.WithTx(tx),
		tx:   tx,
	}
}

func (s *essayQuestionService) GetDB() *gorm.DB {
	if s.tx != nil {
		return s.tx
	}
	return config.DB
}

func (s *essayQuestionService) GetEssayQuestionsByCodeQuestionID(codeQuestionID int64) ([]models.EssayQuestion, error) {
	data, err := s.repo.FindEssayQuestionsByCodeQuestionID(codeQuestionID)
	if err != nil {
		return nil, gorm_err.TranslateGormError(err)
	}
	return data, nil
}

func (s *essayQuestionService) CreateEssayQuestion(data *models.EssayQuestion) (*models.EssayQuestion, error) {
	data, err := s.repo.CreateEssayQuestion(data)
	if err != nil {
		return nil, gorm_err.TranslateGormError(err)
	}
	return data, nil
}

func (s *essayQuestionService) GetEssayQuestionByID(id int64) (*models.EssayQuestion, error) {
	data, err := s.repo.FindEssayQuestionByID(id)
	if err != nil {
		return nil, gorm_err.TranslateGormError(err)
	}
	return data, nil
}




package services

import (
	"jk-api/internal/config"
	"jk-api/internal/database/models"
	"jk-api/internal/errors/gorm_err"
	"jk-api/pkg/repository/adapter/sql"

	"gorm.io/gorm"
)

type TWonderingScoreService interface {
	WithTx(tx *gorm.DB) TWonderingScoreService
	GetTWonderingScoresBySubLessonIDAndUserID(subLessonID, userID int64) (*models.TWonderingScore, error)
	CreateTWonderingScore(data *models.TWonderingScore, userID int64) (*models.TWonderingScore, error)
	GetDB() *gorm.DB
}

type tWonderingScoreService struct {
	repo sql.TWonderingScoreRepository
	tx   *gorm.DB
}

func NewTWonderingScoreService(repo sql.TWonderingScoreRepository) TWonderingScoreService {
	return &tWonderingScoreService{repo: repo}
}

func (s *tWonderingScoreService) WithTx(tx *gorm.DB) TWonderingScoreService {
	return &tWonderingScoreService{
		repo: s.repo.WithTx(tx),
		tx:   tx,
	}
}

func (s *tWonderingScoreService) GetDB() *gorm.DB {
	if s.tx != nil {
		return s.tx
	}
	return config.DB
}

func (s *tWonderingScoreService) GetTWonderingScoresBySubLessonIDAndUserID(subLessonID, userID int64) (*models.TWonderingScore, error) {
	data, err := s.repo.FindTWonderingScoresBySubLessonIDAndUserID(subLessonID, userID)
	if err != nil {
		return nil, gorm_err.TranslateGormError(err)
	}
	return data, nil
}

func (s *tWonderingScoreService) CreateTWonderingScore(data *models.TWonderingScore, userID int64) (*models.TWonderingScore, error) {
	data, err := s.repo.CreateTWonderingScore(data)
	if err != nil {
		return nil, gorm_err.TranslateGormError(err)
	}
	return data, nil
}





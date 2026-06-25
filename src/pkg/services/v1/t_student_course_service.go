package services

import (
	"jk-api/api/http/controllers/v1/dto"
	"jk-api/internal/config"
	"jk-api/internal/database/models"
	"jk-api/internal/errors/gorm_err"
	"jk-api/pkg/repository/adapter/sql"

	"gorm.io/gorm"
)

type TStudentCourseService interface {
	WithTx(tx *gorm.DB) TStudentCourseService

	EnrollTStudentCourse(input *models.TStudentCourse) (*models.TStudentCourse, error)
	GetMyCourse(userID int64, filter dto.TStudentCourseFilterDto, ) ([]models.TStudentCourse, error)
	GetTStudentCourseByID(id int64, filter dto.TStudentCourseFilterDto, userID int64) (*models.TStudentCourse, error)
	GetDB() *gorm.DB
}

type tStudentCourseService struct {
	repo sql.TStudentCourseRepository
	tx   *gorm.DB
}

func NewTStudentCourseService(repo sql.TStudentCourseRepository) TStudentCourseService {
	return &tStudentCourseService{repo: repo}
}

func (s *tStudentCourseService) WithTx(tx *gorm.DB) TStudentCourseService {
	return &tStudentCourseService{
		repo: s.repo.WithTx(tx),
		tx:   tx,
	}
}

func (s *tStudentCourseService) GetDB() *gorm.DB {
	if s.tx != nil {
		return s.tx
	}
	return config.DB
}

func (s *tStudentCourseService) EnrollTStudentCourse(input *models.TStudentCourse) (*models.TStudentCourse, error) {
	data, err := s.repo.EnrollCourse(input)
	if err != nil {
		return nil, gorm_err.TranslateGormError(err)
	}
	return data, nil
}

func (s *tStudentCourseService) GetTStudentCourseByID(id int64, filter dto.TStudentCourseFilterDto, userID int64) (*models.TStudentCourse, error) {
	repo := s.repo
	if filter.Preload {
		repo = repo.WithPreloads("Course", "Badge")
	}
	data, err := s.repo.FindByID(id, userID)
	if err != nil {
		return nil, gorm_err.TranslateGormError(err)
	}
	return data, nil
}

func (s *tStudentCourseService) GetMyCourse(UserID int64, filter dto.TStudentCourseFilterDto) ([]models.TStudentCourse, error) {
	repo := s.repo
	if filter.Preload {
		repo = repo.WithPreloads("Course", "Badge")
	}
	data, err := repo.FindMyCourse(UserID)
	if err != nil {
		return nil, gorm_err.TranslateGormError(err)
	}
	return data, nil
}

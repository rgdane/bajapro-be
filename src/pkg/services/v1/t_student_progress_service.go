package services

import (
	"fmt"
	"jk-api/internal/config"
	"jk-api/internal/database/models"
	"jk-api/internal/errors/gorm_err"
	"jk-api/pkg/repository/adapter/sql"

	"gorm.io/gorm"
)

type TStudentProgressService interface {
	WithTx(tx *gorm.DB) TStudentProgressService

	CompleteTStudentProgress(input *models.TStudentProgress) (*models.TStudentProgress, error)
	GetDB() *gorm.DB
}

type tStudentProgressService struct {
	repo sql.TStudentProgressRepository
	tx   *gorm.DB
}

func NewTStudentProgressService(repo sql.TStudentProgressRepository) TStudentProgressService {
	return &tStudentProgressService{repo: repo}
}

func (s *tStudentProgressService) WithTx(tx *gorm.DB) TStudentProgressService {
	return &tStudentProgressService{
		repo: s.repo.WithTx(tx),
		tx:   tx,
	}
}

func (s *tStudentProgressService) GetDB() *gorm.DB {
	if s.tx != nil {
		return s.tx
	}
	return config.DB
}

func (s *tStudentProgressService) CompleteTStudentProgress(
	input *models.TStudentProgress,
) (*models.TStudentProgress, error) {

	// sudah pernah complete?
	existing, _ := s.repo.FindByUserAndSubLesson(
		input.UserID,
		input.SubLessonID,
	)

	if existing != nil {
		return existing, nil
	}

	progress, err := s.repo.CompleteTStudentProgress(input)
	if err != nil {
		return nil, gorm_err.TranslateGormError(err)
	}

	courseID, err := s.repo.FindCourseIDBySubLesson(
		input.SubLessonID,
	)

	if err != nil {
		return nil, err
	}

	totalSubLesson, err := s.repo.CountTotalSubLessonByCourse(
		courseID,
	)

	if err != nil {
		return nil, err
	}

	completedSubLesson, err := s.repo.CountCompletedByCourse(
		input.UserID,
		courseID,
	)

	if err != nil {
		return nil, err
	}

	var percentage float64

	fmt.Println("CourseID:", courseID)
	fmt.Println("Completed:", completedSubLesson)
	fmt.Println("Total:", totalSubLesson)

	if totalSubLesson > 0 {
		percentage = (float64(completedSubLesson) / float64(totalSubLesson)) * 100
	
	}

	fmt.Println("Percentage:", percentage)
	err = s.repo.UpdateStudentCourseProgress(
		input.UserID,
		courseID,
		percentage,
	)

	if err != nil {
		return nil, err
	}

	return progress, nil
}

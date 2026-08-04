package services

import (
	"jk-api/api/http/controllers/v1/dto"
	"jk-api/internal/config"
	"jk-api/internal/database/models"
	"jk-api/internal/errors/gorm_err"
	"jk-api/pkg/repository/adapter/sql"

	"gorm.io/gorm"
)

type MCourseService interface {
	WithTx(tx *gorm.DB) MCourseService

	CreateMCourse(input *models.MCourse) (*models.MCourse, error)
	UpdateMCourse(id int64, updates map[string]interface{}, associations map[string]interface{}) (*models.MCourse, error)
	DeleteMCourse(id int64) error
	GetAllMCourses(filter dto.MCourseFilterDto) ([]dto.MCourseResponseDto, error)
	GetMCourseDetailByID(
	id int64,
) (*dto.MCourseDetailResponse, error)
	GetDB() *gorm.DB
}

type mCourseService struct {
	repo sql.MCourseRepository
	levelRepo sql.MLevelRepository
	tx   *gorm.DB
}

func NewMCourseService(repo sql.MCourseRepository, levelRepo sql.MLevelRepository) MCourseService {
	return &mCourseService{repo: repo, levelRepo: levelRepo}
}

func (s *mCourseService) WithTx(tx *gorm.DB) MCourseService {
	return &mCourseService{
		repo: s.repo.WithTx(tx),
		levelRepo: s.levelRepo.WithTx(tx),
		tx:   tx,
	}
}

func (s *mCourseService) GetDB() *gorm.DB {
	if s.tx != nil {
		return s.tx
	}
	return config.DB
}

func (s *mCourseService) CreateMCourse(input *models.MCourse) (*models.MCourse, error) {
	data, err := s.repo.InsertMCourse(input)
	if err != nil {
		return nil, gorm_err.TranslateGormError(err)
	}
	return data, nil
}

func (s *mCourseService) UpdateMCourse(
	id int64,
	payload map[string]interface{},
	associations map[string]interface{},
) (*models.MCourse, error) {

	repo := s.repo

	if len(associations) > 0 {
		var assocNames []string
		for name := range associations {
			assocNames = append(assocNames, name)
		}
		repo = repo.WithAssociations(assocNames...).WithReplacements(associations)
	}

	for key := range associations {
		delete(payload, key)
	}

	updated, err := repo.UpdateMCourse(id, payload)
	if err != nil {
		return nil, gorm_err.TranslateGormError(err)
	}

	return updated, nil
}

func (s *mCourseService) DeleteMCourse(id int64) error {
	err := s.repo.RemoveMCourse(id)
	return gorm_err.TranslateGormError(err)
}

func (s *mCourseService) GetAllMCourses(
	filter dto.MCourseFilterDto,
) ([]dto.MCourseResponseDto, error) {

	data, err := s.repo.FindMCourse()

	if err != nil {
		return nil, gorm_err.TranslateGormError(err)
	}

	var response []dto.MCourseResponseDto

	for _, course := range data {

		lessonsCount := len(course.Lessons)

		subLessonsCount := 0

		for _, lesson := range course.Lessons {
			subLessonsCount += len(lesson.SubLessons)
		}

		response = append(response, dto.MCourseResponseDto{
			ID:              course.ID,
			CourseName:      course.CourseName,
			Description:     course.Description,
			ImgThumbnail:    course.ImgThumbnail,
			Published:       course.Published,
			IsActive:        course.IsActive,
			CreatedAt:       course.CreatedAt,
			UpdatedAt:       course.UpdatedAt,
			LessonCount:     lessonsCount,
			SubLessonCount:  subLessonsCount,
		})
	}

	return response, nil
}
func (s *mCourseService) GetMCourseDetailByID(
	id int64,
) (*dto.MCourseDetailResponse, error) {
	course, err := s.repo.FindMCourseByID(id)
	if err != nil {
		return nil, gorm_err.TranslateGormError(err)
	}

	levels, err := s.levelRepo.FindActiveLevels()
	if err != nil {
		return nil, gorm_err.TranslateGormError(err)
	}

	type levelStats struct {
		LessonCount    int
		SubLessonCount int
	}

	statsByLevelID := make(map[int64]levelStats)

	for _, lesson := range course.Lessons {
		stats := statsByLevelID[lesson.LevelID]

		stats.LessonCount++
		stats.SubLessonCount += len(lesson.SubLessons)

		statsByLevelID[lesson.LevelID] = stats
	}

	responseLevels := make([]dto.MCourseLevelResponse, 0, len(levels))

	for _, level := range levels {
		stats := statsByLevelID[level.ID]

		responseLevels = append(
			responseLevels,
			dto.MCourseLevelResponse{
				ID:             level.ID,
				LevelName:      level.LevelName,
				Description:    level.Description,
				IsActive:       level.IsActive,
				CreatedAt:      level.CreatedAt,
				UpdatedAt:      level.UpdatedAt,
				LessonCount:    stats.LessonCount,
				SubLessonCount: stats.SubLessonCount,
				IsLocked:       false,
			},
		)
	}

	return &dto.MCourseDetailResponse{
		ID:           course.ID,
		CourseName:   course.CourseName,
		Description:  course.Description,
		ImgThumbnail: course.ImgThumbnail,
		Published:    course.Published,
		IsActive:     course.IsActive,
		CreatedAt:    course.CreatedAt,
		UpdatedAt:    course.UpdatedAt,
		Levels:       responseLevels,
	}, nil
}
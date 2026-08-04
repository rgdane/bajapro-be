package services

import (
	"errors"
	"jk-api/api/http/controllers/v1/dto"
	"jk-api/api/http/controllers/v1/mapper"
	"jk-api/internal/config"
	"jk-api/internal/database/models"
	"jk-api/internal/errors/gorm_err"
	"jk-api/pkg/repository/adapter/sql"
	"math"
	"sort"
	"strconv"

	"gorm.io/gorm"
)

type TStudentCourseService interface {
	WithTx(tx *gorm.DB) TStudentCourseService

	EnrollTStudentCourse(input *models.TStudentCourse) (*models.TStudentCourse, error)
	GetMyCourse(userID int64, filter dto.TStudentCourseFilterDto) ([]models.TStudentCourse, error)
	GetTStudentCourseByID(id int64, filter dto.TStudentCourseFilterDto, userID int64) (*models.TStudentCourse, error)
	FindLeaderboard(
    courseID int64,
    classID *int64,
) ([]models.Leaderboard, error)
	GetDashboard(userID int64) (*dto.DashboardResponseDto, error)
	GetCourseProgress(
		studentID int64,
		courseID int64,
	) (*dto.StudentCourseProgressResponse, error)
	GetDB() *gorm.DB
}

type tStudentCourseService struct {
	repo                sql.TStudentCourseRepository
	studentCourseRepo   sql.TStudentCourseRepository
	subLessonRepo       sql.MSubLessonRepository
	studentProgressRepo sql.TStudentProgressRepository
	badgeRepo           sql.MBadgeSettingsRepository
	lessonRepo          sql.MLessonRepository
	userRepo            sql.UserRepository
	tx                  *gorm.DB
}

func NewTStudentCourseService(
	studentCourseRepo sql.TStudentCourseRepository,
	subLessonRepo sql.MSubLessonRepository,
	studentProgressRepo sql.TStudentProgressRepository,
	badgeRepo sql.MBadgeSettingsRepository,
	lessonRepo sql.MLessonRepository,
	userRepo sql.UserRepository,
) TStudentCourseService {

	return &tStudentCourseService{
		repo:                studentCourseRepo,
		studentCourseRepo:   studentCourseRepo,
		subLessonRepo:       subLessonRepo,
		studentProgressRepo: studentProgressRepo,
		badgeRepo:           badgeRepo,
		lessonRepo:          lessonRepo,
		userRepo:            userRepo,
	}
}

func (s *tStudentCourseService) WithTx(tx *gorm.DB) TStudentCourseService {
	return &tStudentCourseService{
		repo:                s.repo.WithTx(tx),
		studentCourseRepo:   s.studentCourseRepo.WithTx(tx),
		subLessonRepo:       s.subLessonRepo.WithTx(tx),
		studentProgressRepo: s.studentProgressRepo.WithTx(tx),
		badgeRepo:           s.badgeRepo.WithTx(tx),
		lessonRepo:          s.lessonRepo.WithTx(tx),
		userRepo:            s.userRepo.WithTx(tx),
		tx:                  tx,
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

func (s *tStudentCourseService) FindLeaderboard(
    courseID int64,
    classID *int64,
) ([]models.Leaderboard, error) {

    data, err := s.repo.FindLeaderboard(courseID, classID)
    if err != nil {
        return nil, gorm_err.TranslateGormError(err)
    }

    return data, nil
}

func (s *tStudentCourseService) GetDashboard(userID int64) (*dto.DashboardResponseDto, error) {

	// 1. enrolled course
	enrolledCourses, err := s.studentCourseRepo.FindMyCourse(userID)
	if err != nil {
		return nil, err
	}

	// 2. total sub lesson
	materialsTotal, err := s.subLessonRepo.CountActive()
	if err != nil {
		return nil, err
	}

	// 3. completed material
	completedMaterials, err := s.studentProgressRepo.CountCompleted(userID)
	if err != nil {
		return nil, err
	}

	// 4. badge
	badges, err := s.badgeRepo.FindMBadgeSettings()
	if err != nil {
		return nil, err
	}

	// 5. latest progress
	latestProgress, err := s.studentProgressRepo.FindLatest(userID, 2)
	if err != nil {
		return nil, err
	}

	// 6. leaderboard data
	studentCourses, err := s.studentCourseRepo.FindAllScore()
	if err != nil {
		return nil, err
	}

	students, err := s.userRepo.FindUser()
	if err != nil {
		return nil, err
	}

	//----------------------------------------
	// Leaderboard
	//----------------------------------------

	scoreMap := map[int64]int{}

	for _, sc := range studentCourses {
		scoreMap[sc.UserID] += sc.TotalScore
	}

	leaderboard := []dto.LeaderboardResponseDto{}

	for _, u := range students {
		leaderboard = append(leaderboard, dto.LeaderboardResponseDto{
			ID:    u.ID,
			Name:  u.Name,
			TotalScore: scoreMap[u.ID],
		})
	}

	sort.Slice(leaderboard, func(i, j int) bool {
		return leaderboard[i].TotalScore > leaderboard[j].TotalScore
	})

	top5 := leaderboard
	if len(top5) > 5 {
		top5 = top5[:5]
	}

	//----------------------------------------
	// total score
	//----------------------------------------

	totalScore := 0

	for _, c := range enrolledCourses {
		totalScore += c.TotalScore
	}

	//----------------------------------------
	// progress
	//----------------------------------------

	overallProgress := 0

	if materialsTotal > 0 {
		overallProgress = int(math.Round(
			float64(completedMaterials) /
				float64(materialsTotal) * 100,
		))
	}

	//----------------------------------------
	// badge
	//----------------------------------------

	badgeName := "Beginner"
	badgeImage := ""

	for _, b := range badges {

		if totalScore >= b.MinScore &&
			totalScore <= b.MaxScore {

			badgeName = b.Name
			badgeImage = b.Image
			break
		}

	}

	//----------------------------------------
	// rank
	//----------------------------------------

	rank := "-"

	for i, l := range leaderboard {

		if l.ID == userID {
			rank = strconv.Itoa(i + 1)
			break
		}

	}

	//----------------------------------------
	// lesson history
	//----------------------------------------

	lessonHistory := []dto.LessonHistoryDto{}

	for _, ec := range enrolledCourses {

		totalSub, _ := s.lessonRepo.CountSubLesson(ec.CourseID)

		completedSub, _ := s.studentProgressRepo.CountCompletedByCourse(
			userID,
			ec.CourseID,
		)

		progress := 0

		if totalSub > 0 {
			progress = int(math.Round(
				float64(completedSub) /
					float64(totalSub) * 100,
			))
		}

		lessonHistory = append(
			lessonHistory,
			dto.LessonHistoryDto{
				ID:         ec.ID,
				StudentID:  ec.UserID,
				CourseID:   ec.CourseID,
				TotalScore: ec.TotalScore,
				BadgeID:    ec.BadgeID,
				Course: dto.CourseDto{
					ID:          ec.Course.ID,
					CourseName:  ec.Course.CourseName,
					Image:       ec.Course.ImgThumbnail,
					Description: ec.Course.Description,
				},
				ProgressPercent: progress,
			},
		)
	}

	//----------------------------------------


	//----------------------------------------

	return &dto.DashboardResponseDto{
		EnrolledCount:      len(enrolledCourses),
		MaterialsTotal:     materialsTotal,
		MaterialsCompleted: completedMaterials,
		OverallProgress:    overallProgress,
		LatestProgress:     mapper.StudentProgressModelsToRecentProgressDto(latestProgress),
		LessonHistory:      lessonHistory,
		Achievement: dto.AchievementDto{
			Badge:      badgeName,
			BadgeImage: badgeImage,
			Rank:       rank,
			TotalScore: totalScore,
		},
		Top5Leaderboard: top5,
	}, nil
}

func (s *tStudentCourseService) GetCourseProgress(
	studentID int64,
	courseID int64,
) (*dto.StudentCourseProgressResponse, error) {
	enrollment, err := s.repo.FindByUserIDAndCourseID(
		studentID,
		courseID,
	)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, gorm_err.TranslateGormError(err)
	}

	badge := enrollment.Badge

	if badge == nil {
		foundBadge, badgeErr := s.badgeRepo.FindMBadgeByScore(
			enrollment.TotalScore,
		)

		if badgeErr != nil &&
			!errors.Is(badgeErr, gorm.ErrRecordNotFound) {
			return nil, gorm_err.TranslateGormError(badgeErr)
		}

		if badgeErr == nil {
			badge = foundBadge
		}
	}

	var badgeResponse *dto.BadgeResponse

	if badge != nil {
		badgeResponse = &dto.BadgeResponse{
			ID:        badge.ID,
			Name:      badge.Name,
			Image:     badge.Image,
			MinScore:  badge.MinScore,
			MaxScore:  badge.MaxScore,
			IsActive:  badge.IsActive,
			CreatedAt: badge.CreatedAt,
			UpdatedAt: badge.UpdatedAt,
		}
	}

	return &dto.StudentCourseProgressResponse{
		ID:         enrollment.ID,
		StudentID:  enrollment.UserID,
		CourseID:   enrollment.CourseID,
		TotalScore: enrollment.TotalScore,
		BadgeID:    enrollment.BadgeID,
		IsActive:   enrollment.IsActive,
		CreatedAt:  enrollment.CreatedAt,
		UpdatedAt:  enrollment.UpdatedAt,
		Badge:      badgeResponse,
	}, nil
}
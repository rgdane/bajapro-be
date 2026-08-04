package sql

import (
	"jk-api/internal/database/models"

	"gorm.io/gorm"
)

type TStudentCourseRepository interface {
	WithTx(tx *gorm.DB) TStudentCourseRepository
	WithPreloads(preloads ...string) TStudentCourseRepository
	WithAssociations(associations ...string) TStudentCourseRepository
	WithReplacements(replacements map[string]interface{}) TStudentCourseRepository
	WithJoins(joins ...string) TStudentCourseRepository
	WithWhere(query interface{}, args ...interface{}) TStudentCourseRepository
	WithOrder(order string) TStudentCourseRepository
	WithLimit(limit int) TStudentCourseRepository
	WithCursor(cursor int) TStudentCourseRepository

	EnrollCourse(data *models.TStudentCourse) (*models.TStudentCourse, error)
	FindMyCourse(UserID int64) ([]models.TStudentCourse, error)
	FindByID(id int64, userID int64) (*models.TStudentCourse, error)
	FindByUserIDAndCourseID(userID int64, courseID int64) (*models.TStudentCourse, error)
	UpdateTotalScoreByID(data *models.TStudentCourse) (*models.TStudentCourse, error)
	FindLeaderboard(courseID int64, classID *int64) ([]models.Leaderboard, error)
	FindAllScore() ([]models.TStudentCourse, error)
	
}

package sql

import (
	"jk-api/internal/database/models"

	"gorm.io/gorm"
)

type TStudentProgressRepository interface {
	WithTx(tx *gorm.DB) TStudentProgressRepository
	WithPreloads(preloads ...string) TStudentProgressRepository
	WithAssociations(associations ...string) TStudentProgressRepository
	WithReplacements(replacements map[string]interface{}) TStudentProgressRepository
	WithJoins(joins ...string) TStudentProgressRepository
	WithWhere(query interface{}, args ...interface{}) TStudentProgressRepository
	WithOrder(order string) TStudentProgressRepository
	WithLimit(limit int) TStudentProgressRepository
	WithCursor(cursor int) TStudentProgressRepository

	CompleteTStudentProgress(
		data *models.TStudentProgress,
	) (*models.TStudentProgress, error)

	FindByUserAndSubLesson(
		userID int64,
		subLessonID int64,
	) (*models.TStudentProgress, error)

	FindCourseIDBySubLesson(
		subLessonID int64,
	) (int64, error)

	CountTotalSubLessonByCourse(
		courseID int64,
	) (int64, error)

	CountCompletedByCourse(
		userID int64,
		courseID int64,
	) (int64, error)

	UpdateStudentCourseProgress(
		userID int64,
		courseID int64,
		percentage float64,
	) error
}

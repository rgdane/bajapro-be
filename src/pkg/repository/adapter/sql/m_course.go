package sql

import (
	"jk-api/internal/database/models"

	"gorm.io/gorm"
)

type MCourseRepository interface {
	WithTx(tx *gorm.DB) MCourseRepository
	WithPreloads(preloads ...string) MCourseRepository
	WithAssociations(associations ...string) MCourseRepository
	WithReplacements(replacements map[string]interface{}) MCourseRepository
	WithJoins(joins ...string) MCourseRepository
	WithWhere(query interface{}, args ...interface{}) MCourseRepository
	WithOrder(order string) MCourseRepository
	WithLimit(limit int) MCourseRepository
	WithCursor(cursor int) MCourseRepository

	InsertMCourse(data *models.MCourse) (*models.MCourse, error)
	UpdateMCourse(id int64, updates map[string]interface{}) (*models.MCourse, error)
	UpdateManyMCourses(ids []int64, updates map[string]interface{}) error
	RemoveMCourse(id int64) error
	RemoveManyMCourses(ids []int64) error

	FindMCourse() ([]models.MCourse, error)
	FindMCourseByID(id int64) (*models.MCourse, error)
}

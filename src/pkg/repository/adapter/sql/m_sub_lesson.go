package sql

import (
	"jk-api/internal/database/models"

	"gorm.io/gorm"
)

type MSubLessonRepository interface {
	WithTx(tx *gorm.DB) MSubLessonRepository
	WithPreloads(preloads ...string) MSubLessonRepository
	WithAssociations(associations ...string) MSubLessonRepository
	WithReplacements(replacements map[string]interface{}) MSubLessonRepository
	WithJoins(joins ...string) MSubLessonRepository
	WithWhere(query interface{}, args ...interface{}) MSubLessonRepository
	WithOrder(order string) MSubLessonRepository
	WithLimit(limit int) MSubLessonRepository
	WithCursor(cursor int) MSubLessonRepository
	WithUnscoped() MSubLessonRepository

	InsertMSubLesson(data *models.MSubLesson) (*models.MSubLesson, error)
	InsertManyMSubLessons(data []*models.MSubLesson) ([]*models.MSubLesson, error)
	UpdateMSubLesson(id int64, updates map[string]interface{}) (*models.MSubLesson, error)
	UpdateManyMSubLessons(ids []int64, updates map[string]interface{}) error
	RemoveMSubLesson(id int64) error
	RemoveManyMSubLessons(ids []int64) error

	FindMSubLessons() ([]models.MSubLesson, error)
	FindMSubLessonByID(id int64) (*models.MSubLesson, error)
	FindMSubLessonsByIDs(ids []int64) ([]*models.MSubLesson, error)
}

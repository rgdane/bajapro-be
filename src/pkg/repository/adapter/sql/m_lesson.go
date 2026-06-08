package sql

import (
	"jk-api/internal/database/models"

	"gorm.io/gorm"
)

type MLessonRepository interface {
	WithTx(tx *gorm.DB) MLessonRepository
	WithPreloads(preloads ...string) MLessonRepository
	WithAssociations(associations ...string) MLessonRepository
	WithReplacements(replacements map[string]interface{}) MLessonRepository
	WithJoins(joins ...string) MLessonRepository
	WithWhere(query interface{}, args ...interface{}) MLessonRepository
	WithOrder(order string) MLessonRepository
	WithLimit(limit int) MLessonRepository
	WithCursor(cursor int) MLessonRepository
	WithUnscoped() MLessonRepository

	InsertMLesson(data *models.MLesson) (*models.MLesson, error)
	InsertManyMLessons(data []*models.MLesson) ([]*models.MLesson, error)
	UpdateMLesson(id int64, updates map[string]interface{}) (*models.MLesson, error)
	UpdateManyMLessons(ids []int64, updates map[string]interface{}) error
	RemoveMLesson(id int64) error
	RemoveManyMLessons(ids []int64) error

	FindMLesson() ([]models.MLesson, error)
	FindMLessonByID(id int64) (*models.MLesson, error)
	FindMLessonsByIDs(ids []int64) ([]*models.MLesson, error)
}

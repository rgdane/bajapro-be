package sql

import (
	"jk-api/internal/database/models"

	"gorm.io/gorm"
)

type LevelRepository interface {
	WithTx(tx *gorm.DB) LevelRepository
	WithPreloads(preloads ...string) LevelRepository
	WithAssociations(associations ...string) LevelRepository
	WithReplacements(replacements map[string]interface{}) LevelRepository
	WithJoins(joins ...string) LevelRepository
	WithWhere(query interface{}, args ...interface{}) LevelRepository
	WithOrder(order string) LevelRepository
	WithLimit(limit int) LevelRepository
	WithCursor(cursor int) LevelRepository
	WithUnscoped() LevelRepository

	InsertLevel(data *models.Level) (*models.Level, error)
	InsertManyLevels(data []*models.Level) ([]*models.Level, error)
	UpdateLevel(id int64, updates map[string]interface{}) (*models.Level, error)
	UpdateManyLevels(ids []int64, updates map[string]interface{}) error
	RemoveLevel(id int64) error
	RemoveManyLevels(ids []int64) error

	FindLevel() ([]models.Level, error)
	FindLevelByID(id int64) (*models.Level, error)
	FindLevelsByIDs(ids []int64) ([]*models.Level, error)
}

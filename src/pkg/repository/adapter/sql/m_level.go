package sql

import (
	"jk-api/internal/database/models"

	"gorm.io/gorm"
)

type MLevelRepository interface {
	WithTx(tx *gorm.DB) MLevelRepository
	WithPreloads(preloads ...string) MLevelRepository
	WithAssociations(associations ...string) MLevelRepository
	WithReplacements(replacements map[string]interface{}) MLevelRepository
	WithJoins(joins ...string) MLevelRepository
	WithWhere(query interface{}, args ...interface{}) MLevelRepository
	WithOrder(order string) MLevelRepository
	WithLimit(limit int) MLevelRepository
	WithCursor(cursor int) MLevelRepository

	InsertMLevel(data *models.MLevel) (*models.MLevel, error)
	UpdateMLevel(id int64, updates map[string]interface{}) (*models.MLevel, error)
	RemoveMLevel(id int64) error

	FindMLevel() ([]models.MLevel, error)
	FindMLevelByID(id int64) (*models.MLevel, error)
}

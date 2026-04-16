package sql

import (
	"jk-api/internal/database/models"

	"gorm.io/gorm"
)

type MMaterialRepository interface {
	WithTx(tx *gorm.DB) MMaterialRepository
	WithPreloads(preloads ...string) MMaterialRepository
	WithAssociations(associations ...string) MMaterialRepository
	WithReplacements(replacements map[string]interface{}) MMaterialRepository
	WithJoins(joins ...string) MMaterialRepository
	WithWhere(query interface{}, args ...interface{}) MMaterialRepository
	WithOrder(order string) MMaterialRepository
	WithLimit(limit int) MMaterialRepository
	WithCursor(cursor int) MMaterialRepository
	WithUnscoped() MMaterialRepository

	InsertMMaterial(data *models.MMaterials) (*models.MMaterials, error)
	InsertManyMMaterials(data []*models.MMaterials) ([]*models.MMaterials, error)
	UpdateMMaterial(id int64, updates map[string]interface{}) (*models.MMaterials, error)
	UpdateManyMMaterials(ids []int64, updates map[string]interface{}) error
	RemoveMMaterial(id int64) error
	RemoveManyMMaterials(ids []int64) error

	FindMMaterials() ([]models.MMaterials, error)
	FindMMaterialByID(id int64) (*models.MMaterials, error)
	FindMMaterialsByIDs(ids []int64) ([]*models.MMaterials, error)
}

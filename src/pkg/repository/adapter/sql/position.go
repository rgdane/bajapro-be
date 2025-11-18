package sql

import (
	"jk-api/internal/database/models"

	"gorm.io/gorm"
)

type PositionRepository interface {
	WithTx(tx *gorm.DB) PositionRepository
	WithPreloads(preloads ...string) PositionRepository
	WithAssociations(associations ...string) PositionRepository
	WithReplacements(replacements map[string]interface{}) PositionRepository
	WithJoins(joins ...string) PositionRepository
	WithWhere(query interface{}, args ...interface{}) PositionRepository
	WithOrder(order string) PositionRepository
	WithLimit(limit int) PositionRepository
	WithCursor(cursor int) PositionRepository
	WithUnscoped() PositionRepository

	InsertPosition(data *models.Position) (*models.Position, error)
	InsertManyPositions(data []*models.Position) ([]*models.Position, error)
	UpdatePosition(id int64, updates map[string]interface{}) (*models.Position, error)
	UpdateManyPositions(ids []int64, updates map[string]interface{}) error
	RemovePosition(id int64) error
	RemoveManyPositions(ids []int64) error

	FindPosition() ([]models.Position, error)
	FindPositionByID(id int64) (*models.Position, error)
	FindPositionsByIDs(ids []int64) ([]*models.Position, error)
}

package sql

import (
	"jk-api/internal/database/models"

	"gorm.io/gorm"
)

type DivisionRepository interface {
	WithTx(tx *gorm.DB) DivisionRepository
	WithPreloads(preloads ...string) DivisionRepository
	WithAssociations(associations ...string) DivisionRepository
	WithReplacements(replacements map[string]interface{}) DivisionRepository
	WithJoins(joins ...string) DivisionRepository
	WithWhere(query interface{}, args ...interface{}) DivisionRepository
	WithOrder(order string) DivisionRepository
	WithLimit(limit int) DivisionRepository
	WithCursor(cursor int) DivisionRepository
	WithUnscoped() DivisionRepository

	InsertDivision(data *models.Division) (*models.Division, error)
	InsertManyDivisions(data []*models.Division) ([]*models.Division, error)
	UpdateDivision(id int64, updates map[string]interface{}) (*models.Division, error)
	UpdateManyDivisions(ids []int64, updates map[string]interface{}) error
	RemoveDivision(id int64) error
	RemoveManyDivisions(ids []int64) error

	FindDivision() ([]models.Division, error)
	FindDivisionByID(id int64) (*models.Division, error)
	FindDivisionsByIDs(ids []int64) ([]*models.Division, error)
}

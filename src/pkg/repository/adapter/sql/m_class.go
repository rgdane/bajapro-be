package sql

import (
	"jk-api/internal/database/models"

	"gorm.io/gorm"
)

type MClassRepository interface {
	WithTx(tx *gorm.DB) MClassRepository
	WithPreloads(preloads ...string) MClassRepository
	WithAssociations(associations ...string) MClassRepository
	WithReplacements(replacements map[string]interface{}) MClassRepository
	WithJoins(joins ...string) MClassRepository
	WithWhere(query interface{}, args ...interface{}) MClassRepository
	WithOrder(order string) MClassRepository
	WithLimit(limit int) MClassRepository
	WithCursor(cursor int) MClassRepository

	InsertMClass(data *models.MClass) (*models.MClass, error)
	UpdateMClass(id int64, updates map[string]interface{}) (*models.MClass, error)
	UpdateManyMClasses(ids []int64, updates map[string]interface{}) error
	RemoveMClass(id int64) error
	RemoveManyMClasses(ids []int64) error

	FindMClass() ([]models.MClass, error)
	FindMClassByID(id int64) (*models.MClass, error)
}

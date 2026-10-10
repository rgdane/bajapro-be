package sql

import (
	"jk-api/internal/database/models"

	"gorm.io/gorm"
)

type MSchoolRepository interface {
	WithTx(tx *gorm.DB) MSchoolRepository
	WithPreloads(preloads ...string) MSchoolRepository
	WithAssociations(associations ...string) MSchoolRepository
	WithReplacements(replacements map[string]interface{}) MSchoolRepository
	WithJoins(joins ...string) MSchoolRepository
	WithWhere(query interface{}, args ...interface{}) MSchoolRepository
	WithOrder(order string) MSchoolRepository
	WithLimit(limit int) MSchoolRepository
	WithCursor(cursor int) MSchoolRepository

	InsertMSchool(data *models.MSchool) (*models.MSchool, error)
	UpdateMSchool(id int64, updates map[string]interface{}) (*models.MSchool, error)
	RemoveMSchool(id int64) error
	FindMSchool() ([]models.MSchool, error)
	FindMSchoolByID(id int64) (*models.MSchool, error)
}

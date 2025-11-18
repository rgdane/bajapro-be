package sql

import (
	"jk-api/internal/database/models"

	"gorm.io/gorm"
)

type RoleRepository interface {
	WithTx(tx *gorm.DB) RoleRepository
	WithPreloads(preloads ...string) RoleRepository
	WithAssociations(associations ...string) RoleRepository
	WithReplacements(replacements map[string]interface{}) RoleRepository
	WithJoins(joins ...string) RoleRepository
	WithWhere(query interface{}, args ...interface{}) RoleRepository
	WithOrder(order string) RoleRepository
	WithLimit(limit int) RoleRepository
	WithCursor(cursor int) RoleRepository

	InsertRole(data *models.Role) (*models.Role, error)
	UpdateRole(id int64, updates map[string]interface{}) (*models.Role, error)
	UpdateManyRoles(ids []int64, updates map[string]interface{}) error
	RemoveRole(id int64) error
	RemoveManyRoles(ids []int64) error

	FindRole() ([]models.Role, error)
	FindRoleByID(id int64) (*models.Role, error)
}

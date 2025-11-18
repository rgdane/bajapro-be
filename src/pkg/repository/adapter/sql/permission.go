package sql

import (
	"jk-api/internal/database/models"

	"gorm.io/gorm"
)

type PermissionRepository interface {
	WithTx(tx *gorm.DB) PermissionRepository
	WithPreloads(preloads ...string) PermissionRepository
	WithAssociations(associations ...string) PermissionRepository
	WithReplacements(replacements map[string]interface{}) PermissionRepository
	WithJoins(joins ...string) PermissionRepository
	WithWhere(query interface{}, args ...interface{}) PermissionRepository
	WithOrder(order string) PermissionRepository
	WithLimit(limit int) PermissionRepository
	WithCursor(cursor int) PermissionRepository

	InsertPermission(data *models.Permission) (*models.Permission, error)
	UpdatePermission(id int64, updates map[string]interface{}) (*models.Permission, error)
	UpdateManyPermissions(ids []int64, updates map[string]interface{}) error
	RemovePermission(id int64) error
	RemoveManyPermissions(ids []int64) error

	FindPermission() ([]models.Permission, error)
	FindPermissionByID(id int64) (*models.Permission, error)
}

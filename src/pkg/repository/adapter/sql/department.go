package sql

import (
	"jk-api/internal/database/models"

	"gorm.io/gorm"
)

type DepartmentRepository interface {
	WithTx(tx *gorm.DB) DepartmentRepository
	WithPreloads(preloads ...string) DepartmentRepository
	WithAssociations(associations ...string) DepartmentRepository
	WithReplacements(replacements map[string]interface{}) DepartmentRepository
	WithJoins(joins ...string) DepartmentRepository
	WithWhere(query interface{}, args ...interface{}) DepartmentRepository
	WithOrder(order string) DepartmentRepository
	WithLimit(limit int) DepartmentRepository
	WithCursor(cursor int) DepartmentRepository
	WithUnscoped() DepartmentRepository

	InsertDepartment(data *models.Department) (*models.Department, error)
	InsertManyDepartments(data []*models.Department) ([]*models.Department, error)
	UpdateDepartment(id int64, updates map[string]interface{}) (*models.Department, error)
	UpdateManyDepartments(ids []int64, updates map[string]interface{}) error
	RemoveDepartment(id int64) error
	RemoveManyDepartments(ids []int64) error

	FindDepartment() ([]models.Department, error)
	FindDepartmentByID(id int64) (*models.Department, error)
	FindDepartmentsByIDs(ids []int64) ([]*models.Department, error)
}

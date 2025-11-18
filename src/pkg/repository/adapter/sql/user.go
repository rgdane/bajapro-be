package sql

import (
	"jk-api/internal/database/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	WithTx(tx *gorm.DB) UserRepository
	WithPreloads(preloads ...string) UserRepository
	WithAssociations(associations ...string) UserRepository
	WithReplacements(replacements map[string]interface{}) UserRepository
	WithJoins(joins ...string) UserRepository
	WithWhere(query interface{}, args ...interface{}) UserRepository
	WithOrder(order string) UserRepository
	WithLimit(limit int) UserRepository
	WithCursor(cursor int) UserRepository
	WithUnscoped() UserRepository

	InsertUser(data *models.User) (*models.User, error)
	InsertManyUsers(data []*models.User) ([]*models.User, error)

	UpdateUser(id int64, updates map[string]interface{}) (*models.User, error)
	UpdateManyUsers(ids []int64, updates map[string]interface{}) error
	RemoveUser(id int64) error
	RemoveManyUsers(ids []int64) error

	FindUser() ([]models.User, error)
	FindUserByID(id int64) (*models.User, error)
	FindUserByEmail(email string) (*models.User, error)
}

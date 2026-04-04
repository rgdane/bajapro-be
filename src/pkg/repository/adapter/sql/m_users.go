package sql

import (
	"jk-api/internal/database/models"

	"gorm.io/gorm"
)

type MUsersRepository interface {
	WithTx(tx *gorm.DB) MUsersRepository
	WithPreloads(preloads ...string) MUsersRepository
	WithAssociations(associations ...string) MUsersRepository
	WithReplacements(replacements map[string]interface{}) MUsersRepository
	WithJoins(joins ...string) MUsersRepository
	WithWhere(query interface{}, args ...interface{}) MUsersRepository
	WithOrder(order string) MUsersRepository
	WithLimit(limit int) MUsersRepository
	WithCursor(cursor int) MUsersRepository
	WithUnscoped() MUsersRepository

	InsertMUser(data *models.MUsers) (*models.MUsers, error)
	InsertManyMUsers(data []*models.MUsers) ([]*models.MUsers, error)

	UpdateMUser(id int64, updates map[string]interface{}) (*models.MUsers, error)
	UpdateManyMUsers(ids []int64, updates map[string]interface{}) error
	RemoveMUser(id int64) error
	RemoveManyMUsers(ids []int64) error

	FindMUser() ([]models.MUsers, error)
	FindMUserByID(id int64) (*models.MUsers, error)
	FindMUserByEmail(email string) (*models.MUsers, error)

	GetDB() *gorm.DB
}

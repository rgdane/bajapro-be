// adapter/sql/base.go
package sql

import "gorm.io/gorm"

type BaseRepository[T any] interface {
	WithPreloads(...string) BaseRepository[T]
	WithJoins(...string) BaseRepository[T]
	WithWhere(func(*gorm.DB) *gorm.DB) BaseRepository[T]
	WithOrder(string) BaseRepository[T]
	WithLimit(int) BaseRepository[T]
	WithCursor(int) BaseRepository[T]
	WithAssociations(...string) BaseRepository[T]
	WithReplacements(map[string]interface{}) BaseRepository[T]

	FindAll() ([]T, error)
	FindByID(id any) (*T, error)
	FindAllPtr() ([]*T, error)
	Create(*T) error
	UpdateByID(id int64, updates map[string]interface{}) error
	UpdateMany(ids []int64, updates map[string]interface{}) error
	Delete(id any) error

	Sum(string, any) error
}

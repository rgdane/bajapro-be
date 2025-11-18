package sql

import (
	"jk-api/internal/database/models"

	"gorm.io/gorm"
)

type TitleRepository interface {
	WithTx(tx *gorm.DB) TitleRepository
	WithPreloads(preloads ...string) TitleRepository
	WithAssociations(associations ...string) TitleRepository
	WithReplacements(replacements map[string]interface{}) TitleRepository
	WithJoins(joins ...string) TitleRepository
	WithWhere(query interface{}, args ...interface{}) TitleRepository
	WithOrder(order string) TitleRepository
	WithLimit(limit int) TitleRepository
	WithCursor(cursor int) TitleRepository
	WithUnscoped() TitleRepository

	InsertTitle(data *models.Title) (*models.Title, error)
	InsertManyTitles(data []*models.Title) ([]*models.Title, error)
	UpdateTitle(id int64, updates map[string]interface{}) (*models.Title, error)
	UpdateManyTitles(ids []int64, updates map[string]interface{}) error
	RemoveTitle(id int64) error
	RemoveManyTitles(ids []int64) error

	FindTitle() ([]models.Title, error)
	FindTitleByID(id int64) (*models.Title, error)
	FindTitlesByIDs(ids []int64) ([]*models.Title, error)
}

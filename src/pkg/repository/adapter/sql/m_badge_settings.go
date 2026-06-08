package sql

import (
	"jk-api/internal/database/models"

	"gorm.io/gorm"
)

type MBadgeSettingsRepository interface {
	WithTx(tx *gorm.DB) MBadgeSettingsRepository
	WithPreloads(preloads ...string) MBadgeSettingsRepository
	WithAssociations(associations ...string) MBadgeSettingsRepository
	WithReplacements(replacements map[string]interface{}) MBadgeSettingsRepository
	WithJoins(joins ...string) MBadgeSettingsRepository
	WithWhere(query interface{}, args ...interface{}) MBadgeSettingsRepository
	WithOrder(order string) MBadgeSettingsRepository
	WithLimit(limit int) MBadgeSettingsRepository
	WithCursor(cursor int) MBadgeSettingsRepository

	InsertMBadgeSettings(data *models.MBadgeSettings) (*models.MBadgeSettings, error)
	UpdateMBadgeSettings(id int64, updates map[string]interface{}) (*models.MBadgeSettings, error)
	UpdateManyMBadgeSettings(ids []int64, updates map[string]interface{}) error
	RemoveMBadgeSettings(id int64) error
	RemoveManyMBadgeSettings(ids []int64) error

	FindMBadgeSettings() ([]models.MBadgeSettings, error)
	FindMBadgeSettingsByID(id int64) (*models.MBadgeSettings, error)
}

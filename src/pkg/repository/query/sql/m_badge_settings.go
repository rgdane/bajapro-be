package sql

import (
	"jk-api/internal/config"
	"jk-api/internal/database/models"
	adapter "jk-api/pkg/repository/adapter/sql"
	"jk-api/pkg/repository/query/sql/builder"

	"gorm.io/gorm"
)

type mBadgeSettingsRepository struct {
	db           *gorm.DB
	preloads     []string
	associations []string
	replacements map[string]interface{}
	joins        []string
	whereClauses []func(*gorm.DB) *gorm.DB
	order        string
	limit        *int
	cursor       *int
}

func NewMBadgeSettingsRepository() adapter.MBadgeSettingsRepository {
	return &mBadgeSettingsRepository{db: config.DB}
}

// --- 🔁 Chainable Configs ---

func (repo *mBadgeSettingsRepository) clone() *mBadgeSettingsRepository {
	clone := *repo
	return &clone
}

func (repo *mBadgeSettingsRepository) WithTx(tx *gorm.DB) adapter.MBadgeSettingsRepository {
	clone := repo.clone()
	clone.db = tx
	return clone
}

func (repo *mBadgeSettingsRepository) WithPreloads(preloads ...string) adapter.MBadgeSettingsRepository {
	clone := repo.clone()
	clone.preloads = append(clone.preloads, preloads...)
	return clone
}

func (repo *mBadgeSettingsRepository) WithAssociations(associations ...string) adapter.MBadgeSettingsRepository {
	clone := repo.clone()
	clone.associations = append(clone.associations, associations...)
	return clone
}

func (repo *mBadgeSettingsRepository) WithReplacements(replacements map[string]interface{}) adapter.MBadgeSettingsRepository {
	clone := repo.clone()
	clone.replacements = replacements
	return clone
}

func (repo *mBadgeSettingsRepository) WithJoins(joins ...string) adapter.MBadgeSettingsRepository {
	clone := repo.clone()
	clone.joins = append(clone.joins, joins...)
	return clone
}

func (repo *mBadgeSettingsRepository) WithWhere(query interface{}, args ...interface{}) adapter.MBadgeSettingsRepository {
	clone := repo.clone()
	clone.whereClauses = append(clone.whereClauses, func(db *gorm.DB) *gorm.DB {
		return db.Where(query, args...)
	})
	return clone
}

func (repo *mBadgeSettingsRepository) WithOrder(order string) adapter.MBadgeSettingsRepository {
	clone := repo.clone()
	clone.order = order
	return clone
}

func (repo *mBadgeSettingsRepository) WithLimit(limit int) adapter.MBadgeSettingsRepository {
	clone := repo.clone()
	clone.limit = &limit
	return clone
}

func (repo *mBadgeSettingsRepository) WithCursor(cursor int) adapter.MBadgeSettingsRepository {
	clone := repo.clone()
	clone.cursor = &cursor
	return clone
}

// --- 🔧 Query Builder Helper ---

func (repo *mBadgeSettingsRepository) getQueryBuilder() *builder.QueryBuilder[models.MBadgeSettings] {
	qb := builder.NewQueryBuilder[models.MBadgeSettings](repo.db).
		WithPreloads(repo.preloads...).
		WithAssociations(repo.associations...).
		WithReplacements(repo.replacements).
		WithJoins(repo.joins...).
		WithOrder(repo.order)

	for _, w := range repo.whereClauses {
		qb = qb.WithWhere(w)
	}
	if repo.limit != nil {
		qb = qb.WithLimit(*repo.limit)
	}
	if repo.cursor != nil {
		qb = qb.WithCursor(*repo.cursor)
	}
	return qb
}

// --- 🔧 CRUD Methods ---

func (repo *mBadgeSettingsRepository) InsertMBadgeSettings(data *models.MBadgeSettings) (*models.MBadgeSettings, error) {
	if err := repo.getQueryBuilder().Create(data); err != nil {
		return nil, err
	}
	return data, nil
}

func (repo *mBadgeSettingsRepository) UpdateMBadgeSettings(id int64, updates map[string]interface{}) (*models.MBadgeSettings, error) {
	return repo.getQueryBuilder().UpdateByID(id, updates)
}

func (repo *mBadgeSettingsRepository) UpdateManyMBadgeSettings(ids []int64, updates map[string]interface{}) error {
	_, err := repo.getQueryBuilder().UpdateMany(ids, updates)
	return err
}

func (repo *mBadgeSettingsRepository) RemoveMBadgeSettings(id int64) error {
	return repo.getQueryBuilder().Delete(id)
}

func (repo *mBadgeSettingsRepository) RemoveManyMBadgeSettings(ids []int64) error {
	return repo.getQueryBuilder().WithWhere(func(db *gorm.DB) *gorm.DB {
		return db.Where("id IN ?", ids)
	}).Delete(nil)
}

func (repo *mBadgeSettingsRepository) FindMBadgeSettings() ([]models.MBadgeSettings, error) {
	return repo.getQueryBuilder().FindAll()
}

func (repo *mBadgeSettingsRepository) FindMBadgeSettingsByID(id int64) (*models.MBadgeSettings, error) {
	return repo.getQueryBuilder().FindByID(id)
}

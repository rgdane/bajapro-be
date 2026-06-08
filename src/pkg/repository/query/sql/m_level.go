package sql

import (
	"jk-api/internal/config"
	"jk-api/internal/database/models"
	adapter "jk-api/pkg/repository/adapter/sql"
	"jk-api/pkg/repository/query/sql/builder"

	"gorm.io/gorm"
)

type mLevelRepository struct {
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

func NewMLevelRepository() adapter.MLevelRepository {
	return &mLevelRepository{db: config.DB}
}

// --- 🔁 Chainable Configs ---

func (repo *mLevelRepository) clone() *mLevelRepository {
	clone := *repo
	return &clone
}

func (repo *mLevelRepository) WithTx(tx *gorm.DB) adapter.MLevelRepository {
	clone := repo.clone()
	clone.db = tx
	return clone
}

func (repo *mLevelRepository) WithPreloads(preloads ...string) adapter.MLevelRepository {
	clone := repo.clone()
	clone.preloads = append(clone.preloads, preloads...)
	return clone
}

func (repo *mLevelRepository) WithAssociations(associations ...string) adapter.MLevelRepository {
	clone := repo.clone()
	clone.associations = append(clone.associations, associations...)
	return clone
}

func (repo *mLevelRepository) WithReplacements(replacements map[string]interface{}) adapter.MLevelRepository {
	clone := repo.clone()
	clone.replacements = replacements
	return clone
}

func (repo *mLevelRepository) WithJoins(joins ...string) adapter.MLevelRepository {
	clone := repo.clone()
	clone.joins = append(clone.joins, joins...)
	return clone
}

func (repo *mLevelRepository) WithWhere(query interface{}, args ...interface{}) adapter.MLevelRepository {
	clone := repo.clone()
	clone.whereClauses = append(clone.whereClauses, func(db *gorm.DB) *gorm.DB {
		return db.Where(query, args...)
	})
	return clone
}

func (repo *mLevelRepository) WithOrder(order string) adapter.MLevelRepository {
	clone := repo.clone()
	clone.order = order
	return clone
}

func (repo *mLevelRepository) WithLimit(limit int) adapter.MLevelRepository {
	clone := repo.clone()
	clone.limit = &limit
	return clone
}

func (repo *mLevelRepository) WithCursor(cursor int) adapter.MLevelRepository {
	clone := repo.clone()
	clone.cursor = &cursor
	return clone
}

// --- 🔧 Query Builder Helper ---

func (repo *mLevelRepository) getQueryBuilder() *builder.QueryBuilder[models.MLevel] {
	qb := builder.NewQueryBuilder[models.MLevel](repo.db).
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

func (repo *mLevelRepository) InsertMLevel(data *models.MLevel) (*models.MLevel, error) {
	if err := repo.getQueryBuilder().Create(data); err != nil {
		return nil, err
	}
	return data, nil
}

func (repo *mLevelRepository) UpdateMLevel(id int64, updates map[string]interface{}) (*models.MLevel, error) {
	return repo.getQueryBuilder().UpdateByID(id, updates)
}

func (repo *mLevelRepository) UpdateManyMLevels(ids []int64, updates map[string]interface{}) error {
	_, err := repo.getQueryBuilder().UpdateMany(ids, updates)
	return err
}

func (repo *mLevelRepository) RemoveMLevel(id int64) error {
	return repo.getQueryBuilder().Delete(id)
}

func (repo *mLevelRepository) RemoveManyMLevels(ids []int64) error {
	return repo.getQueryBuilder().WithWhere(func(db *gorm.DB) *gorm.DB {
		return db.Where("id IN ?", ids)
	}).Delete(nil)
}

func (repo *mLevelRepository) FindMLevel() ([]models.MLevel, error) {
	return repo.getQueryBuilder().FindAll()
}

func (repo *mLevelRepository) FindMLevelByID(id int64) (*models.MLevel, error) {
	return repo.getQueryBuilder().FindByID(id)
}



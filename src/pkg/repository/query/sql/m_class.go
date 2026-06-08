package sql

import (
	"jk-api/internal/config"
	"jk-api/internal/database/models"
	adapter "jk-api/pkg/repository/adapter/sql"
	"jk-api/pkg/repository/query/sql/builder"

	"gorm.io/gorm"
)

type mClassRepository struct {
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

func NewMClassRepository() adapter.MClassRepository {
	return &mClassRepository{db: config.DB}
}

// --- 🔁 Chainable Configs ---

func (repo *mClassRepository) clone() *mClassRepository {
	clone := *repo
	return &clone
}

func (repo *mClassRepository) WithTx(tx *gorm.DB) adapter.MClassRepository {
	clone := repo.clone()
	clone.db = tx
	return clone
}

func (repo *mClassRepository) WithPreloads(preloads ...string) adapter.MClassRepository {
	clone := repo.clone()
	clone.preloads = append(clone.preloads, preloads...)
	return clone
}

func (repo *mClassRepository) WithAssociations(associations ...string) adapter.MClassRepository {
	clone := repo.clone()
	clone.associations = append(clone.associations, associations...)
	return clone
}

func (repo *mClassRepository) WithReplacements(replacements map[string]interface{}) adapter.MClassRepository {
	clone := repo.clone()
	clone.replacements = replacements
	return clone
}

func (repo *mClassRepository) WithJoins(joins ...string) adapter.MClassRepository {
	clone := repo.clone()
	clone.joins = append(clone.joins, joins...)
	return clone
}

func (repo *mClassRepository) WithWhere(query interface{}, args ...interface{}) adapter.MClassRepository {
	clone := repo.clone()
	clone.whereClauses = append(clone.whereClauses, func(db *gorm.DB) *gorm.DB {
		return db.Where(query, args...)
	})
	return clone
}

func (repo *mClassRepository) WithOrder(order string) adapter.MClassRepository {
	clone := repo.clone()
	clone.order = order
	return clone
}

func (repo *mClassRepository) WithLimit(limit int) adapter.MClassRepository {
	clone := repo.clone()
	clone.limit = &limit
	return clone
}

func (repo *mClassRepository) WithCursor(cursor int) adapter.MClassRepository {
	clone := repo.clone()
	clone.cursor = &cursor
	return clone
}

// --- 🔧 Query Builder Helper ---

func (repo *mClassRepository) getQueryBuilder() *builder.QueryBuilder[models.MClass] {
	qb := builder.NewQueryBuilder[models.MClass](repo.db).
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

func (repo *mClassRepository) InsertMClass(data *models.MClass) (*models.MClass, error) {
	if err := repo.getQueryBuilder().Create(data); err != nil {
		return nil, err
	}
	return data, nil
}

func (repo *mClassRepository) UpdateMClass(id int64, updates map[string]interface{}) (*models.MClass, error) {
	return repo.getQueryBuilder().UpdateByID(id, updates)
}

func (repo *mClassRepository) UpdateManyMClasses(ids []int64, updates map[string]interface{}) error {
	_, err := repo.getQueryBuilder().UpdateMany(ids, updates)
	return err
}

func (repo *mClassRepository) RemoveMClass(id int64) error {
	return repo.getQueryBuilder().Delete(id)
}

func (repo *mClassRepository) RemoveManyMClasses(ids []int64) error {
	return repo.getQueryBuilder().WithWhere(func(db *gorm.DB) *gorm.DB {
		return db.Where("id IN ?", ids)
	}).Delete(nil)
}

func (repo *mClassRepository) FindMClass() ([]models.MClass, error) {
	return repo.getQueryBuilder().FindAll()
}

func (repo *mClassRepository) FindMClassByID(id int64) (*models.MClass, error) {
	return repo.getQueryBuilder().FindByID(id)
}

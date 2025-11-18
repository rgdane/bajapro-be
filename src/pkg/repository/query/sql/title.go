package sql

import (
	"jk-api/internal/config"
	"jk-api/internal/database/models"
	adapter "jk-api/pkg/repository/adapter/sql"
	"jk-api/pkg/repository/query/sql/builder"

	"gorm.io/gorm"
)

type titleRepository struct {
	db           *gorm.DB
	preloads     []string
	associations []string
	replacements map[string]interface{}
	joins        []string
	whereClauses []func(*gorm.DB) *gorm.DB
	order        string
	limit        *int
	cursor       *int
	unscoped     bool
}

func NewTitleRepository() adapter.TitleRepository {
	return &titleRepository{db: config.DB}
}

// --- 🔁 Chainable Configs ---

func (repo *titleRepository) clone() *titleRepository {
	clone := *repo
	return &clone
}

func (repo *titleRepository) WithTx(tx *gorm.DB) adapter.TitleRepository {
	clone := repo.clone()
	clone.db = tx
	return clone
}

func (repo *titleRepository) WithPreloads(preloads ...string) adapter.TitleRepository {
	clone := repo.clone()
	clone.preloads = append(clone.preloads, preloads...)
	return clone
}

func (repo *titleRepository) WithAssociations(associations ...string) adapter.TitleRepository {
	clone := repo.clone()
	clone.associations = append(clone.associations, associations...)
	return clone
}

func (repo *titleRepository) WithReplacements(replacements map[string]interface{}) adapter.TitleRepository {
	clone := repo.clone()
	clone.replacements = replacements
	return clone
}

func (repo *titleRepository) WithJoins(joins ...string) adapter.TitleRepository {
	clone := repo.clone()
	clone.joins = append(clone.joins, joins...)
	return clone
}

func (repo *titleRepository) WithWhere(query interface{}, args ...interface{}) adapter.TitleRepository {
	clone := repo.clone()
	clone.whereClauses = append(clone.whereClauses, func(db *gorm.DB) *gorm.DB {
		return db.Where(query, args...)
	})
	return clone
}

func (repo *titleRepository) WithOrder(order string) adapter.TitleRepository {
	clone := repo.clone()
	clone.order = order
	return clone
}

func (repo *titleRepository) WithLimit(limit int) adapter.TitleRepository {
	clone := repo.clone()
	clone.limit = &limit
	return clone
}

func (repo *titleRepository) WithCursor(cursor int) adapter.TitleRepository {
	clone := repo.clone()
	clone.cursor = &cursor
	return clone
}

func (repo *titleRepository) WithUnscoped() adapter.TitleRepository {
	clone := repo.clone()
	clone.unscoped = true
	return clone
}

// --- 🔧 Builder ---

func (repo *titleRepository) getQueryBuilder() *builder.QueryBuilder[models.Title] {
	db := repo.db
	if repo.unscoped {
		db = db.Unscoped()
	}
	qb := builder.NewQueryBuilder[models.Title](db).
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

// --- ⚙️ CRUD ---

func (repo *titleRepository) InsertTitle(data *models.Title) (*models.Title, error) {
	if err := repo.getQueryBuilder().Create(data); err != nil {
		return nil, err
	}
	return data, nil
}

func (repo *titleRepository) InsertManyTitles(data []*models.Title) ([]*models.Title, error) {
	if err := repo.getQueryBuilder().CreateMany(data); err != nil {
		return nil, err
	}
	return data, nil
}

func (repo *titleRepository) UpdateTitle(id int64, updates map[string]interface{}) (*models.Title, error) {
	return repo.getQueryBuilder().UpdateByID(id, updates)
}

func (repo *titleRepository) UpdateManyTitles(ids []int64, updates map[string]interface{}) error {
	_, err := repo.getQueryBuilder().UpdateMany(ids, updates)
	return err
}

func (repo *titleRepository) RemoveTitle(id int64) error {
	return repo.getQueryBuilder().Delete(id)
}

func (repo *titleRepository) RemoveManyTitles(ids []int64) error {
	return repo.getQueryBuilder().Delete(ids)
}

// --- 🔍 Query ---

func (repo *titleRepository) FindTitle() ([]models.Title, error) {
	return repo.getQueryBuilder().FindAll()
}

func (repo *titleRepository) FindTitleByID(id int64) (*models.Title, error) {
	return repo.getQueryBuilder().FindByID(id)
}

func (repo *titleRepository) FindTitlesByIDs(ids []int64) ([]*models.Title, error) {
	if len(ids) == 0 {
		return []*models.Title{}, nil
	}

	return repo.getQueryBuilder().WithWhere(func(db *gorm.DB) *gorm.DB {
		return db.Where("id IN ?", ids)
	}).FindAllPtr()
}

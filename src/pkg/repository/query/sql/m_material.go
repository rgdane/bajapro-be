package sql

import (
	"jk-api/internal/config"
	"jk-api/internal/database/models"
	adapter "jk-api/pkg/repository/adapter/sql"
	"jk-api/pkg/repository/query/sql/builder"

	"gorm.io/gorm"
)

type mMaterialRepository struct {
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

func NewMMaterialRepository() adapter.MMaterialRepository {
	return &mMaterialRepository{db: config.DB}
}

// --- 🔁 Chainable Configs ---

func (repo *mMaterialRepository) clone() *mMaterialRepository {
	clone := *repo
	return &clone
}

func (repo *mMaterialRepository) WithTx(tx *gorm.DB) adapter.MMaterialRepository {
	clone := repo.clone()
	clone.db = tx
	return clone
}

func (repo *mMaterialRepository) WithPreloads(preloads ...string) adapter.MMaterialRepository {
	clone := repo.clone()
	clone.preloads = append(clone.preloads, preloads...)
	return clone
}

func (repo *mMaterialRepository) WithAssociations(associations ...string) adapter.MMaterialRepository {
	clone := repo.clone()
	clone.associations = append(clone.associations, associations...)
	return clone
}

func (repo *mMaterialRepository) WithReplacements(replacements map[string]interface{}) adapter.MMaterialRepository {
	clone := repo.clone()
	clone.replacements = replacements
	return clone
}

func (repo *mMaterialRepository) WithJoins(joins ...string) adapter.MMaterialRepository {
	clone := repo.clone()
	clone.joins = append(clone.joins, joins...)
	return clone
}

func (repo *mMaterialRepository) WithWhere(query interface{}, args ...interface{}) adapter.MMaterialRepository {
	clone := repo.clone()
	clone.whereClauses = append(clone.whereClauses, func(db *gorm.DB) *gorm.DB {
		return db.Where(query, args...)
	})
	return clone
}

func (repo *mMaterialRepository) WithOrder(order string) adapter.MMaterialRepository {
	clone := repo.clone()
	clone.order = order
	return clone
}

func (repo *mMaterialRepository) WithLimit(limit int) adapter.MMaterialRepository {
	clone := repo.clone()
	clone.limit = &limit
	return clone
}

func (repo *mMaterialRepository) WithCursor(cursor int) adapter.MMaterialRepository {
	clone := repo.clone()
	clone.cursor = &cursor
	return clone
}

func (repo *mMaterialRepository) WithUnscoped() adapter.MMaterialRepository {
	clone := repo.clone()
	clone.unscoped = true
	return clone
}

// --- 🔧 Builder ---

func (repo *mMaterialRepository) getQueryBuilder() *builder.QueryBuilder[models.MMaterials] {
	db := repo.db
	if repo.unscoped {
		db = db.Unscoped()
	}

	qb := builder.NewQueryBuilder[models.MMaterials](db).
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

func (repo *mMaterialRepository) InsertMMaterial(data *models.MMaterials) (*models.MMaterials, error) {
	if err := repo.getQueryBuilder().Create(data); err != nil {
		return nil, err
	}
	return data, nil
}

func (repo *mMaterialRepository) InsertManyMMaterials(data []*models.MMaterials) ([]*models.MMaterials, error) {
	if err := repo.getQueryBuilder().CreateMany(data); err != nil {
		return nil, err
	}
	return data, nil
}

func (repo *mMaterialRepository) UpdateMMaterial(id int64, updates map[string]interface{}) (*models.MMaterials, error) {
	return repo.getQueryBuilder().UpdateByID(id, updates)
}

func (repo *mMaterialRepository) UpdateManyMMaterials(ids []int64, updates map[string]interface{}) error {
	_, err := repo.getQueryBuilder().UpdateMany(ids, updates)
	return err
}

func (repo *mMaterialRepository) RemoveMMaterial(id int64) error {
	return repo.getQueryBuilder().Delete(id)
}

func (repo *mMaterialRepository) RemoveManyMMaterials(ids []int64) error {
	return repo.getQueryBuilder().Delete(ids)
}

func (repo *mMaterialRepository) FindMMaterials() ([]models.MMaterials, error) {
	return repo.getQueryBuilder().FindAll()
}

func (repo *mMaterialRepository) FindMMaterialByID(id int64) (*models.MMaterials, error) {
	return repo.getQueryBuilder().FindByID(id)
}

func (repo *mMaterialRepository) FindMMaterialsByIDs(ids []int64) ([]*models.MMaterials, error) {
	if len(ids) == 0 {
		return []*models.MMaterials{}, nil
	}

	return repo.getQueryBuilder().WithWhere(func(db *gorm.DB) *gorm.DB {
		return db.Where("id IN ?", ids)
	}).FindAllPtr()
}

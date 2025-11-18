package sql

import (
	"jk-api/internal/config"
	"jk-api/internal/database/models"
	adapter "jk-api/pkg/repository/adapter/sql"
	"jk-api/pkg/repository/query/sql/builder"

	"gorm.io/gorm"
)

type divisionRepository struct {
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

func NewDivisionRepository() adapter.DivisionRepository {
	return &divisionRepository{db: config.DB}
}

// --- 🔁 Chainable Methods ---

func (repo *divisionRepository) clone() *divisionRepository {
	clone := *repo
	return &clone
}

func (repo *divisionRepository) WithTx(tx *gorm.DB) adapter.DivisionRepository {
	clone := repo.clone()
	clone.db = tx
	return clone
}

func (repo *divisionRepository) WithPreloads(preloads ...string) adapter.DivisionRepository {
	clone := repo.clone()
	clone.preloads = append(clone.preloads, preloads...)
	return clone
}

func (repo *divisionRepository) WithAssociations(associations ...string) adapter.DivisionRepository {
	clone := repo.clone()
	clone.associations = append(clone.associations, associations...)
	return clone
}

func (repo *divisionRepository) WithReplacements(replacements map[string]interface{}) adapter.DivisionRepository {
	clone := repo.clone()
	clone.replacements = replacements
	return clone
}

func (repo *divisionRepository) WithJoins(joins ...string) adapter.DivisionRepository {
	clone := repo.clone()
	clone.joins = append(clone.joins, joins...)
	return clone
}

func (repo *divisionRepository) WithWhere(query interface{}, args ...interface{}) adapter.DivisionRepository {
	clone := repo.clone()
	clone.whereClauses = append(clone.whereClauses, func(db *gorm.DB) *gorm.DB {
		return db.Where(query, args...)
	})
	return clone
}
func (repo *divisionRepository) WithOrder(order string) adapter.DivisionRepository {
	clone := repo.clone()
	clone.order = order
	return clone
}

func (repo *divisionRepository) WithLimit(limit int) adapter.DivisionRepository {
	clone := repo.clone()
	clone.limit = &limit
	return clone
}

func (repo *divisionRepository) WithCursor(cursor int) adapter.DivisionRepository {
	clone := repo.clone()
	clone.cursor = &cursor
	return clone
}

func (repo *divisionRepository) WithUnscoped() adapter.DivisionRepository {
	clone := repo.clone()
	clone.unscoped = true
	return clone
}

// --- 🧱 Builder Helper ---

func (repo *divisionRepository) getQueryBuilder() *builder.QueryBuilder[models.Division] {
	db := repo.db
	if repo.unscoped {
		db = db.Unscoped()
	}
	qb := builder.NewQueryBuilder[models.Division](db).
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

// --- 📦 CRUD Methods ---

func (repo *divisionRepository) InsertDivision(data *models.Division) (*models.Division, error) {
	if err := repo.getQueryBuilder().Create(data); err != nil {
		return nil, err
	}
	return data, nil
}

func (repo *divisionRepository) InsertManyDivisions(data []*models.Division) ([]*models.Division, error) {
	if err := repo.getQueryBuilder().CreateMany(data); err != nil {
		return nil, err
	}
	return data, nil
}

func (repo *divisionRepository) UpdateDivision(id int64, updates map[string]interface{}) (*models.Division, error) {
	return repo.getQueryBuilder().UpdateByID(id, updates)
}

func (repo *divisionRepository) UpdateManyDivisions(ids []int64, updates map[string]interface{}) error {
	_, err := repo.getQueryBuilder().UpdateMany(ids, updates)
	return err
}

func (repo *divisionRepository) RemoveDivision(id int64) error {
	return repo.getQueryBuilder().Delete(id)
}

func (repo *divisionRepository) RemoveManyDivisions(ids []int64) error {
	return repo.getQueryBuilder().Delete(ids)
}

func (repo *divisionRepository) FindDivision() ([]models.Division, error) {
	return repo.getQueryBuilder().FindAll()
}

func (repo *divisionRepository) FindDivisionByID(id int64) (*models.Division, error) {
	return repo.getQueryBuilder().FindByID(id)
}

func (repo *divisionRepository) FindDivisionsByIDs(ids []int64) ([]*models.Division, error) {
	if len(ids) == 0 {
		return []*models.Division{}, nil
	}

	return repo.getQueryBuilder().WithWhere(func(db *gorm.DB) *gorm.DB {
		return db.Where("id IN ?", ids)
	}).FindAllPtr()
}

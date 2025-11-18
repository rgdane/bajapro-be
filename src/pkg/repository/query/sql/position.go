package sql

import (
	"jk-api/internal/config"
	"jk-api/internal/database/models"
	adapter "jk-api/pkg/repository/adapter/sql"
	"jk-api/pkg/repository/query/sql/builder"

	"gorm.io/gorm"
)

type positionRepository struct {
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

func NewPositionRepository() adapter.PositionRepository {
	return &positionRepository{db: config.DB}
}

// --- 🔁 Chainable Methods ---

func (repo *positionRepository) clone() *positionRepository {
	clone := *repo
	return &clone
}

func (repo *positionRepository) WithTx(tx *gorm.DB) adapter.PositionRepository {
	clone := repo.clone()
	clone.db = tx
	return clone
}

func (repo *positionRepository) WithPreloads(preloads ...string) adapter.PositionRepository {
	clone := repo.clone()
	clone.preloads = append(clone.preloads, preloads...)
	return clone
}

func (repo *positionRepository) WithAssociations(associations ...string) adapter.PositionRepository {
	clone := repo.clone()
	clone.associations = append(clone.associations, associations...)
	return clone
}

func (repo *positionRepository) WithReplacements(replacements map[string]interface{}) adapter.PositionRepository {
	clone := repo.clone()
	clone.replacements = replacements
	return clone
}

func (repo *positionRepository) WithJoins(joins ...string) adapter.PositionRepository {
	clone := repo.clone()
	clone.joins = append(clone.joins, joins...)
	return clone
}

func (repo *positionRepository) WithWhere(query interface{}, args ...interface{}) adapter.PositionRepository {
	clone := repo.clone()
	clone.whereClauses = append(clone.whereClauses, func(db *gorm.DB) *gorm.DB {
		return db.Where(query, args...)
	})
	return clone
}
func (repo *positionRepository) WithOrder(order string) adapter.PositionRepository {
	clone := repo.clone()
	clone.order = order
	return clone
}

func (repo *positionRepository) WithLimit(limit int) adapter.PositionRepository {
	clone := repo.clone()
	clone.limit = &limit
	return clone
}

func (repo *positionRepository) WithCursor(cursor int) adapter.PositionRepository {
	clone := repo.clone()
	clone.cursor = &cursor
	return clone
}

func (repo *positionRepository) WithUnscoped() adapter.PositionRepository {
	clone := repo.clone()
	clone.unscoped = true
	return clone
}

// --- 🧱 Builder Helper ---

func (repo *positionRepository) getQueryBuilder() *builder.QueryBuilder[models.Position] {
	db := repo.db
	if repo.unscoped {
		db = db.Unscoped()
	}
	qb := builder.NewQueryBuilder[models.Position](db).
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

func (repo *positionRepository) InsertPosition(data *models.Position) (*models.Position, error) {
	if err := repo.getQueryBuilder().Create(data); err != nil {
		return nil, err
	}
	return data, nil
}

func (repo *positionRepository) InsertManyPositions(data []*models.Position) ([]*models.Position, error) {
	if err := repo.getQueryBuilder().CreateMany(data); err != nil {
		return nil, err
	}
	return data, nil
}

func (repo *positionRepository) UpdatePosition(id int64, updates map[string]interface{}) (*models.Position, error) {
	return repo.getQueryBuilder().UpdateByID(id, updates)
}

func (repo *positionRepository) UpdateManyPositions(ids []int64, updates map[string]interface{}) error {
	_, err := repo.getQueryBuilder().UpdateMany(ids, updates)
	return err
}

func (repo *positionRepository) RemovePosition(id int64) error {
	return repo.getQueryBuilder().Delete(id)
}

func (repo *positionRepository) RemoveManyPositions(ids []int64) error {
	return repo.getQueryBuilder().Delete(ids)
}

func (repo *positionRepository) FindPosition() ([]models.Position, error) {
	return repo.getQueryBuilder().FindAll()
}

func (repo *positionRepository) FindPositionByID(id int64) (*models.Position, error) {
	return repo.getQueryBuilder().FindByID(id)
}

func (repo *positionRepository) FindPositionsByIDs(ids []int64) ([]*models.Position, error) {
	if len(ids) == 0 {
		return []*models.Position{}, nil
	}

	return repo.getQueryBuilder().WithWhere(func(db *gorm.DB) *gorm.DB {
		return db.Where("id IN ?", ids)
	}).FindAllPtr()
}

package sql

import (
	"jk-api/internal/config"
	"jk-api/internal/database/models"
	adapter "jk-api/pkg/repository/adapter/sql"
	"jk-api/pkg/repository/query/sql/builder"

	"gorm.io/gorm"
)

type mSchoolRepository struct {
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

func NewMSchoolRepository() adapter.MSchoolRepository {
	return &mSchoolRepository{db: config.DB}
}

func (repo *mSchoolRepository) clone() *mSchoolRepository {
	clone := *repo
	return &clone
}

func (repo *mSchoolRepository) WithTx(tx *gorm.DB) adapter.MSchoolRepository {
	clone := repo.clone()
	clone.db = tx
	return clone
}

func (repo *mSchoolRepository) WithPreloads(preloads ...string) adapter.MSchoolRepository {
	clone := repo.clone()
	clone.preloads = append(clone.preloads, preloads...)
	return clone
}

func (repo *mSchoolRepository) WithAssociations(associations ...string) adapter.MSchoolRepository {
	clone := repo.clone()
	clone.associations = append(clone.associations, associations...)
	return clone
}

func (repo *mSchoolRepository) WithReplacements(replacements map[string]interface{}) adapter.MSchoolRepository {
	clone := repo.clone()
	clone.replacements = replacements
	return clone
}

func (repo *mSchoolRepository) WithJoins(joins ...string) adapter.MSchoolRepository {
	clone := repo.clone()
	clone.joins = append(clone.joins, joins...)
	return clone
}

func (repo *mSchoolRepository) WithWhere(query interface{}, args ...interface{}) adapter.MSchoolRepository {
	clone := repo.clone()
	clone.whereClauses = append(clone.whereClauses, func(db *gorm.DB) *gorm.DB {
		return db.Where(query, args...)
	})
	return clone
}

func (repo *mSchoolRepository) WithOrder(order string) adapter.MSchoolRepository {
	clone := repo.clone()
	clone.order = order
	return clone
}

func (repo *mSchoolRepository) WithLimit(limit int) adapter.MSchoolRepository {
	clone := repo.clone()
	clone.limit = &limit
	return clone
}

func (repo *mSchoolRepository) WithCursor(cursor int) adapter.MSchoolRepository {
	clone := repo.clone()
	clone.cursor = &cursor
	return clone
}

func (repo *mSchoolRepository) getQueryBuilder() *builder.QueryBuilder[models.MSchool] {
	qb := builder.NewQueryBuilder[models.MSchool](repo.db).
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

func (repo *mSchoolRepository) InsertMSchool(data *models.MSchool) (*models.MSchool, error) {
	if err := repo.getQueryBuilder().Create(data); err != nil {
		return nil, err
	}
	return data, nil
}

func (repo *mSchoolRepository) UpdateMSchool(id int64, updates map[string]interface{}) (*models.MSchool, error) {
	return repo.getQueryBuilder().UpdateByID(id, updates)
}

func (repo *mSchoolRepository) RemoveMSchool(id int64) error {
	return repo.getQueryBuilder().Delete(id)
}

func (repo *mSchoolRepository) FindMSchool() ([]models.MSchool, error) {
	return repo.getQueryBuilder().FindAll()
}

func (repo *mSchoolRepository) FindMSchoolByID(id int64) (*models.MSchool, error) {
	return repo.getQueryBuilder().FindByID(id)
}

package sql

import (
	"jk-api/internal/config"
	"jk-api/internal/database/models"
	adapter "jk-api/pkg/repository/adapter/sql"
	"jk-api/pkg/repository/query/sql/builder"

	"gorm.io/gorm"
)

type mCourseRepository struct {
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

func NewMCourseRepository() adapter.MCourseRepository {
	return &mCourseRepository{db: config.DB}
}

// --- 🔁 Chainable Configs ---

func (repo *mCourseRepository) clone() *mCourseRepository {
	clone := *repo
	return &clone
}

func (repo *mCourseRepository) WithTx(tx *gorm.DB) adapter.MCourseRepository {
	clone := repo.clone()
	clone.db = tx
	return clone
}

func (repo *mCourseRepository) WithPreloads(preloads ...string) adapter.MCourseRepository {
	clone := repo.clone()
	clone.preloads = append(clone.preloads, preloads...)
	return clone
}

func (repo *mCourseRepository) WithAssociations(associations ...string) adapter.MCourseRepository {
	clone := repo.clone()
	clone.associations = append(clone.associations, associations...)
	return clone
}

func (repo *mCourseRepository) WithReplacements(replacements map[string]interface{}) adapter.MCourseRepository {
	clone := repo.clone()
	clone.replacements = replacements
	return clone
}

func (repo *mCourseRepository) WithJoins(joins ...string) adapter.MCourseRepository {
	clone := repo.clone()
	clone.joins = append(clone.joins, joins...)
	return clone
}

func (repo *mCourseRepository) WithWhere(query interface{}, args ...interface{}) adapter.MCourseRepository {
	clone := repo.clone()
	clone.whereClauses = append(clone.whereClauses, func(db *gorm.DB) *gorm.DB {
		return db.Where(query, args...)
	})
	return clone
}

func (repo *mCourseRepository) WithOrder(order string) adapter.MCourseRepository {
	clone := repo.clone()
	clone.order = order
	return clone
}

func (repo *mCourseRepository) WithLimit(limit int) adapter.MCourseRepository {
	clone := repo.clone()
	clone.limit = &limit
	return clone
}

func (repo *mCourseRepository) WithCursor(cursor int) adapter.MCourseRepository {
	clone := repo.clone()
	clone.cursor = &cursor
	return clone
}

// --- 🔧 Query Builder Helper ---

func (repo *mCourseRepository) getQueryBuilder() *builder.QueryBuilder[models.MCourse] {
	qb := builder.NewQueryBuilder[models.MCourse](repo.db).
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

func (repo *mCourseRepository) InsertMCourse(data *models.MCourse) (*models.MCourse, error) {
	if err := repo.getQueryBuilder().Create(data); err != nil {
		return nil, err
	}
	return data, nil
}

func (repo *mCourseRepository) UpdateMCourse(id int64, updates map[string]interface{}) (*models.MCourse, error) {
	return repo.getQueryBuilder().UpdateByID(id, updates)
}

func (repo *mCourseRepository) UpdateManyMCourses(ids []int64, updates map[string]interface{}) error {
	_, err := repo.getQueryBuilder().UpdateMany(ids, updates)
	return err
}

func (repo *mCourseRepository) RemoveMCourse(id int64) error {
	return repo.getQueryBuilder().Delete(id)
}

func (repo *mCourseRepository) RemoveManyMCourses(ids []int64) error {
	return repo.getQueryBuilder().WithWhere(func(db *gorm.DB) *gorm.DB {
		return db.Where("id IN ?", ids)
	}).Delete(nil)
}

func (repo *mCourseRepository) FindMCourse() ([]models.MCourse, error) {
	return repo.getQueryBuilder().FindAll()
}

func (repo *mCourseRepository) FindMCourseByID(id int64) (*models.MCourse, error) {
	return repo.getQueryBuilder().FindByID(id)
}

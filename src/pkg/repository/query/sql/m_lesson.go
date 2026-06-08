package sql

import (
	"jk-api/internal/config"
	"jk-api/internal/database/models"
	adapter "jk-api/pkg/repository/adapter/sql"
	"jk-api/pkg/repository/query/sql/builder"

	"gorm.io/gorm"
)

type mLessonRepository struct {
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

func NewMLessonRepository() adapter.MLessonRepository {
	return &mLessonRepository{db: config.DB}
}

// --- 🔁 Chainable Configs ---

func (repo *mLessonRepository) clone() *mLessonRepository {
	clone := *repo
	return &clone
}

func (repo *mLessonRepository) WithTx(tx *gorm.DB) adapter.MLessonRepository {
	clone := repo.clone()
	clone.db = tx
	return clone
}

func (repo *mLessonRepository) WithPreloads(preloads ...string) adapter.MLessonRepository {
	clone := repo.clone()
	clone.preloads = append(clone.preloads, preloads...)
	return clone
}

func (repo *mLessonRepository) WithAssociations(associations ...string) adapter.MLessonRepository {
	clone := repo.clone()
	clone.associations = append(clone.associations, associations...)
	return clone
}

func (repo *mLessonRepository) WithReplacements(replacements map[string]interface{}) adapter.MLessonRepository {
	clone := repo.clone()
	clone.replacements = replacements
	return clone
}

func (repo *mLessonRepository) WithJoins(joins ...string) adapter.MLessonRepository {
	clone := repo.clone()
	clone.joins = append(clone.joins, joins...)
	return clone
}

func (repo *mLessonRepository) WithWhere(query interface{}, args ...interface{}) adapter.MLessonRepository {
	clone := repo.clone()
	clone.whereClauses = append(clone.whereClauses, func(db *gorm.DB) *gorm.DB {
		return db.Where(query, args...)
	})
	return clone
}

func (repo *mLessonRepository) WithOrder(order string) adapter.MLessonRepository {
	clone := repo.clone()
	clone.order = order
	return clone
}

func (repo *mLessonRepository) WithLimit(limit int) adapter.MLessonRepository {
	clone := repo.clone()
	clone.limit = &limit
	return clone
}

func (repo *mLessonRepository) WithCursor(cursor int) adapter.MLessonRepository {
	clone := repo.clone()
	clone.cursor = &cursor
	return clone
}

func (repo *mLessonRepository) WithUnscoped() adapter.MLessonRepository {
	clone := repo.clone()
	clone.unscoped = true
	return clone
}

// --- 🔧 Builder ---

func (repo *mLessonRepository) getQueryBuilder() *builder.QueryBuilder[models.MLesson] {
	db := repo.db
	if repo.unscoped {
		db = db.Unscoped()
	}

	qb := builder.NewQueryBuilder[models.MLesson](db).
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

func (repo *mLessonRepository) InsertMLesson(data *models.MLesson) (*models.MLesson, error) {
	if err := repo.getQueryBuilder().Create(data); err != nil {
		return nil, err
	}
	return data, nil
}

func (repo *mLessonRepository) InsertManyMLessons(data []*models.MLesson) ([]*models.MLesson, error) {
	if err := repo.getQueryBuilder().CreateMany(data); err != nil {
		return nil, err
	}
	return data, nil
}

func (repo *mLessonRepository) UpdateMLesson(id int64, updates map[string]interface{}) (*models.MLesson, error) {
	return repo.getQueryBuilder().UpdateByID(id, updates)
}

func (repo *mLessonRepository) UpdateManyMLessons(ids []int64, updates map[string]interface{}) error {
	_, err := repo.getQueryBuilder().UpdateMany(ids, updates)
	return err
}

func (repo *mLessonRepository) RemoveMLesson(id int64) error {
	return repo.getQueryBuilder().Delete(id)
}

func (repo *mLessonRepository) RemoveManyMLessons(ids []int64) error {
	return repo.getQueryBuilder().Delete(ids)
}

func (repo *mLessonRepository) FindMLesson() ([]models.MLesson, error) {
	return repo.getQueryBuilder().FindAll()
}

func (repo *mLessonRepository) FindMLessonByID(id int64) (*models.MLesson, error) {
	return repo.getQueryBuilder().FindByID(id)
}

func (repo *mLessonRepository) FindMLessonsByIDs(ids []int64) ([]*models.MLesson, error) {
	if len(ids) == 0 {
		return []*models.MLesson{}, nil
	}

	return repo.getQueryBuilder().WithWhere(func(db *gorm.DB) *gorm.DB {
		return db.Where("id IN ?", ids)
	}).FindAllPtr()
}

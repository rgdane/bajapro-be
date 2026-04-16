package sql

import (
	"jk-api/internal/config"
	"jk-api/internal/database/models"
	adapter "jk-api/pkg/repository/adapter/sql"
	"jk-api/pkg/repository/query/sql/builder"

	"gorm.io/gorm"
)

type mSubLessonRepository struct {
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

func NewMSubLessonRepository() adapter.MSubLessonRepository {
	return &mSubLessonRepository{db: config.DB}
}

// --- 🔁 Chainable Configs ---

func (repo *mSubLessonRepository) clone() *mSubLessonRepository {
	clone := *repo
	return &clone
}

func (repo *mSubLessonRepository) WithTx(tx *gorm.DB) adapter.MSubLessonRepository {
	clone := repo.clone()
	clone.db = tx
	return clone
}

func (repo *mSubLessonRepository) WithPreloads(preloads ...string) adapter.MSubLessonRepository {
	clone := repo.clone()
	clone.preloads = append(clone.preloads, preloads...)
	return clone
}

func (repo *mSubLessonRepository) WithAssociations(associations ...string) adapter.MSubLessonRepository {
	clone := repo.clone()
	clone.associations = append(clone.associations, associations...)
	return clone
}

func (repo *mSubLessonRepository) WithReplacements(replacements map[string]interface{}) adapter.MSubLessonRepository {
	clone := repo.clone()
	clone.replacements = replacements
	return clone
}

func (repo *mSubLessonRepository) WithJoins(joins ...string) adapter.MSubLessonRepository {
	clone := repo.clone()
	clone.joins = append(clone.joins, joins...)
	return clone
}

func (repo *mSubLessonRepository) WithWhere(query interface{}, args ...interface{}) adapter.MSubLessonRepository {
	clone := repo.clone()
	clone.whereClauses = append(clone.whereClauses, func(db *gorm.DB) *gorm.DB {
		return db.Where(query, args...)
	})
	return clone
}

func (repo *mSubLessonRepository) WithOrder(order string) adapter.MSubLessonRepository {
	clone := repo.clone()
	clone.order = order
	return clone
}

func (repo *mSubLessonRepository) WithLimit(limit int) adapter.MSubLessonRepository {
	clone := repo.clone()
	clone.limit = &limit
	return clone
}

func (repo *mSubLessonRepository) WithCursor(cursor int) adapter.MSubLessonRepository {
	clone := repo.clone()
	clone.cursor = &cursor
	return clone
}

func (repo *mSubLessonRepository) WithUnscoped() adapter.MSubLessonRepository {
	clone := repo.clone()
	clone.unscoped = true
	return clone
}

// --- 🔧 Builder ---

func (repo *mSubLessonRepository) getQueryBuilder() *builder.QueryBuilder[models.MSubLesson] {
	db := repo.db
	if repo.unscoped {
		db = db.Unscoped()
	}

	qb := builder.NewQueryBuilder[models.MSubLesson](db).
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

func (repo *mSubLessonRepository) InsertMSubLesson(data *models.MSubLesson) (*models.MSubLesson, error) {
	if err := repo.getQueryBuilder().Create(data); err != nil {
		return nil, err
	}
	return data, nil
}

func (repo *mSubLessonRepository) InsertManyMSubLessons(data []*models.MSubLesson) ([]*models.MSubLesson, error) {
	if err := repo.getQueryBuilder().CreateMany(data); err != nil {
		return nil, err
	}
	return data, nil
}

func (repo *mSubLessonRepository) UpdateMSubLesson(id int64, updates map[string]interface{}) (*models.MSubLesson, error) {
	return repo.getQueryBuilder().UpdateByID(id, updates)
}

func (repo *mSubLessonRepository) UpdateManyMSubLessons(ids []int64, updates map[string]interface{}) error {
	_, err := repo.getQueryBuilder().UpdateMany(ids, updates)
	return err
}

func (repo *mSubLessonRepository) RemoveMSubLesson(id int64) error {
	return repo.getQueryBuilder().Delete(id)
}

func (repo *mSubLessonRepository) RemoveManyMSubLessons(ids []int64) error {
	return repo.getQueryBuilder().Delete(ids)
}

func (repo *mSubLessonRepository) FindMSubLessons() ([]models.MSubLesson, error) {
	return repo.getQueryBuilder().FindAll()
}

func (repo *mSubLessonRepository) FindMSubLessonByID(id int64) (*models.MSubLesson, error) {
	return repo.getQueryBuilder().FindByID(id)
}

func (repo *mSubLessonRepository) FindMSubLessonsByIDs(ids []int64) ([]*models.MSubLesson, error) {
	if len(ids) == 0 {
		return []*models.MSubLesson{}, nil
	}

	return repo.getQueryBuilder().WithWhere(func(db *gorm.DB) *gorm.DB {
		return db.Where("id IN ?", ids)
	}).FindAllPtr()
}

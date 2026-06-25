package sql

import (
	"jk-api/internal/config"
	"jk-api/internal/database/models"
	adapter "jk-api/pkg/repository/adapter/sql"
	"jk-api/pkg/repository/query/sql/builder"

	"gorm.io/gorm"
)

type tWonderingScoreRepository struct {
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

func NewTWonderingScoreRepository() adapter.TWonderingScoreRepository {
	return &tWonderingScoreRepository{db: config.DB}
}

// --- 🔁 Chainable Configs ---

func (repo *tWonderingScoreRepository) clone() *tWonderingScoreRepository {
	clone := *repo
	return &clone
}

func (repo *tWonderingScoreRepository) WithTx(tx *gorm.DB) adapter.TWonderingScoreRepository {
	clone := repo.clone()
	clone.db = tx
	return clone
}

func (repo *tWonderingScoreRepository) WithPreloads(preloads ...string) adapter.TWonderingScoreRepository {
	clone := repo.clone()
	clone.preloads = append(clone.preloads, preloads...)
	return clone
}

func (repo *tWonderingScoreRepository) WithAssociations(associations ...string) adapter.TWonderingScoreRepository {
	clone := repo.clone()
	clone.associations = append(clone.associations, associations...)
	return clone
}

func (repo *tWonderingScoreRepository) WithReplacements(replacements map[string]interface{}) adapter.TWonderingScoreRepository {
	clone := repo.clone()
	clone.replacements = replacements
	return clone
}

func (repo *tWonderingScoreRepository) WithJoins(joins ...string) adapter.TWonderingScoreRepository {
	clone := repo.clone()
	clone.joins = append(clone.joins, joins...)
	return clone
}

func (repo *tWonderingScoreRepository) WithWhere(query interface{}, args ...interface{}) adapter.TWonderingScoreRepository {
	clone := repo.clone()
	clone.whereClauses = append(clone.whereClauses, func(db *gorm.DB) *gorm.DB {
		return db.Where(query, args...)
	})
	return clone
}

func (repo *tWonderingScoreRepository) WithOrder(order string) adapter.TWonderingScoreRepository {
	clone := repo.clone()
	clone.order = order
	return clone
}

func (repo *tWonderingScoreRepository) WithLimit(limit int) adapter.TWonderingScoreRepository {
	clone := repo.clone()
	clone.limit = &limit
	return clone
}

func (repo *tWonderingScoreRepository) WithCursor(cursor int) adapter.TWonderingScoreRepository {
	clone := repo.clone()
	clone.cursor = &cursor
	return clone
}

// --- 🔧 Query Builder Helper ---

func (repo *tWonderingScoreRepository) getQueryBuilder() *builder.QueryBuilder[models.TWonderingScore] {
	qb := builder.NewQueryBuilder[models.TWonderingScore](repo.db).
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

func (repo *tWonderingScoreRepository) CreateTWonderingScore(data *models.TWonderingScore) (*models.TWonderingScore, error) {
	if err := repo.getQueryBuilder().Create(data); err != nil {
		return nil, err
	}
	return data, nil
}


func (repo *tWonderingScoreRepository) FindTWonderingScoresBySubLessonIDAndUserID(subLessonID, userID int64) (*models.TWonderingScore, error) {
	return repo.getQueryBuilder().
		WithWhere(func(db *gorm.DB) *gorm.DB {
			return db.Where("sub_lesson_id = ? AND user_id = ?", subLessonID, userID)
		}).
		WithPreloads("SubLesson", "User").
		FindOne()
}
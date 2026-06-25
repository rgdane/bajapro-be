package sql

import (
	"jk-api/internal/config"
	"jk-api/internal/database/models"
	adapter "jk-api/pkg/repository/adapter/sql"
	"jk-api/pkg/repository/query/sql/builder"

	"gorm.io/gorm"
)

type tEssayAnswerRepository struct {
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

func NewTEssayAnswerRepository() adapter.TEssayAnswerRepository {
	return &tEssayAnswerRepository{db: config.DB}
}

// --- 🔁 Chainable Configs ---

func (repo *tEssayAnswerRepository) clone() *tEssayAnswerRepository {
	clone := *repo
	return &clone
}

func (repo *tEssayAnswerRepository) WithTx(tx *gorm.DB) adapter.TEssayAnswerRepository {
	clone := repo.clone()
	clone.db = tx
	return clone
}

func (repo *tEssayAnswerRepository) WithPreloads(preloads ...string) adapter.TEssayAnswerRepository {
	clone := repo.clone()
	clone.preloads = append(clone.preloads, preloads...)
	return clone
}

func (repo *tEssayAnswerRepository) WithAssociations(associations ...string) adapter.TEssayAnswerRepository {
	clone := repo.clone()
	clone.associations = append(clone.associations, associations...)
	return clone
}

func (repo *tEssayAnswerRepository) WithReplacements(replacements map[string]interface{}) adapter.TEssayAnswerRepository {
	clone := repo.clone()
	clone.replacements = replacements
	return clone
}

func (repo *tEssayAnswerRepository) WithJoins(joins ...string) adapter.TEssayAnswerRepository {
	clone := repo.clone()
	clone.joins = append(clone.joins, joins...)
	return clone
}

func (repo *tEssayAnswerRepository) WithWhere(query interface{}, args ...interface{}) adapter.TEssayAnswerRepository {
	clone := repo.clone()
	clone.whereClauses = append(clone.whereClauses, func(db *gorm.DB) *gorm.DB {
		return db.Where(query, args...)
	})
	return clone
}

func (repo *tEssayAnswerRepository) WithOrder(order string) adapter.TEssayAnswerRepository {
	clone := repo.clone()
	clone.order = order
	return clone
}

func (repo *tEssayAnswerRepository) WithLimit(limit int) adapter.TEssayAnswerRepository {
	clone := repo.clone()
	clone.limit = &limit
	return clone
}

func (repo *tEssayAnswerRepository) WithCursor(cursor int) adapter.TEssayAnswerRepository {
	clone := repo.clone()
	clone.cursor = &cursor
	return clone
}

// --- 🔧 Query Builder Helper ---

func (repo *tEssayAnswerRepository) getQueryBuilder() *builder.QueryBuilder[models.TEssayAnswer] {
	qb := builder.NewQueryBuilder[models.TEssayAnswer](repo.db).
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

func (repo *tEssayAnswerRepository) CreateTEssayAnswer(data *models.TEssayAnswer) (*models.TEssayAnswer, error) {
	if err := repo.getQueryBuilder().Create(data); err != nil {
		return nil, err
	}
	return data, nil
}


func (repo *tEssayAnswerRepository) FindTEssayAnswersByEssayQuestionIDAndUserID(essayQuestionID, userID int64) (*models.TEssayAnswer, error) {
	return repo.getQueryBuilder().
		WithWhere(func(db *gorm.DB) *gorm.DB {
			return db.Where("essay_question_id = ? AND user_id = ?", essayQuestionID, userID)
		}).
		WithPreloads("EssayQuestion", "User").
		FindOne()
}
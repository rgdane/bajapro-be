package sql

import (
	"jk-api/internal/config"
	"jk-api/internal/database/models"
	adapter "jk-api/pkg/repository/adapter/sql"
	"jk-api/pkg/repository/query/sql/builder"

	"gorm.io/gorm"
)

type tCodeAnswerRepository struct {
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

func NewTCodeAnswerRepository() adapter.TCodeAnswerRepository {
	return &tCodeAnswerRepository{db: config.DB}
}

// --- 🔁 Chainable Configs ---

func (repo *tCodeAnswerRepository) clone() *tCodeAnswerRepository {
	clone := *repo
	return &clone
}

func (repo *tCodeAnswerRepository) WithTx(tx *gorm.DB) adapter.TCodeAnswerRepository {
	clone := repo.clone()
	clone.db = tx
	return clone
}

func (repo *tCodeAnswerRepository) WithPreloads(preloads ...string) adapter.TCodeAnswerRepository {
	clone := repo.clone()
	clone.preloads = append(clone.preloads, preloads...)
	return clone
}

func (repo *tCodeAnswerRepository) WithAssociations(associations ...string) adapter.TCodeAnswerRepository {
	clone := repo.clone()
	clone.associations = append(clone.associations, associations...)
	return clone
}

func (repo *tCodeAnswerRepository) WithReplacements(replacements map[string]interface{}) adapter.TCodeAnswerRepository {
	clone := repo.clone()
	clone.replacements = replacements
	return clone
}

func (repo *tCodeAnswerRepository) WithJoins(joins ...string) adapter.TCodeAnswerRepository {
	clone := repo.clone()
	clone.joins = append(clone.joins, joins...)
	return clone
}

func (repo *tCodeAnswerRepository) WithWhere(query interface{}, args ...interface{}) adapter.TCodeAnswerRepository {
	clone := repo.clone()
	clone.whereClauses = append(clone.whereClauses, func(db *gorm.DB) *gorm.DB {
		return db.Where(query, args...)
	})
	return clone
}

func (repo *tCodeAnswerRepository) WithOrder(order string) adapter.TCodeAnswerRepository {
	clone := repo.clone()
	clone.order = order
	return clone
}

func (repo *tCodeAnswerRepository) WithLimit(limit int) adapter.TCodeAnswerRepository {
	clone := repo.clone()
	clone.limit = &limit
	return clone
}

func (repo *tCodeAnswerRepository) WithCursor(cursor int) adapter.TCodeAnswerRepository {
	clone := repo.clone()
	clone.cursor = &cursor
	return clone
}

// --- 🔧 Query Builder Helper ---

func (repo *tCodeAnswerRepository) getQueryBuilder() *builder.QueryBuilder[models.TCodeAnswer] {
	qb := builder.NewQueryBuilder[models.TCodeAnswer](repo.db).
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

func (repo *tCodeAnswerRepository) CreateTCodeAnswer(data *models.TCodeAnswer) (*models.TCodeAnswer, error) {
	if err := repo.getQueryBuilder().Create(data); err != nil {
		return nil, err
	}
	return data, nil
}


func (repo *tCodeAnswerRepository) FindTCodeAnswersByCodeQuestionID(codeQuestionID int64) ([]models.TCodeAnswer, error) {
	return repo.getQueryBuilder().
		WithWhere(func(db *gorm.DB) *gorm.DB {
			return db.Where("code_question_id = ?", codeQuestionID)
		}).
		WithPreloads("CodeQuestion", "User").
		FindAll()
}
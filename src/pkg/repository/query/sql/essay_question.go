package sql

import (
	"jk-api/internal/config"
	"jk-api/internal/database/models"
	adapter "jk-api/pkg/repository/adapter/sql"
	"jk-api/pkg/repository/query/sql/builder"

	"gorm.io/gorm"
)

type essayQuestionRepository struct {
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

func NewEssayQuestionRepository() adapter.EssayQuestionRepository {
	return &essayQuestionRepository{db: config.DB}
}

// --- 🔁 Chainable Configs ---

func (repo *essayQuestionRepository) clone() *essayQuestionRepository {
	clone := *repo
	return &clone
}

func (repo *essayQuestionRepository) WithTx(tx *gorm.DB) adapter.EssayQuestionRepository {
	clone := repo.clone()
	clone.db = tx
	return clone
}

func (repo *essayQuestionRepository) WithPreloads(preloads ...string) adapter.EssayQuestionRepository {
	clone := repo.clone()
	clone.preloads = append(clone.preloads, preloads...)
	return clone
}

func (repo *essayQuestionRepository) WithAssociations(associations ...string) adapter.EssayQuestionRepository {
	clone := repo.clone()
	clone.associations = append(clone.associations, associations...)
	return clone
}

func (repo *essayQuestionRepository) WithReplacements(replacements map[string]interface{}) adapter.EssayQuestionRepository {
	clone := repo.clone()
	clone.replacements = replacements
	return clone
}

func (repo *essayQuestionRepository) WithJoins(joins ...string) adapter.EssayQuestionRepository {
	clone := repo.clone()
	clone.joins = append(clone.joins, joins...)
	return clone
}

func (repo *essayQuestionRepository) WithWhere(query interface{}, args ...interface{}) adapter.EssayQuestionRepository {
	clone := repo.clone()
	clone.whereClauses = append(clone.whereClauses, func(db *gorm.DB) *gorm.DB {
		return db.Where(query, args...)
	})
	return clone
}

func (repo *essayQuestionRepository) WithOrder(order string) adapter.EssayQuestionRepository {
	clone := repo.clone()
	clone.order = order
	return clone
}

func (repo *essayQuestionRepository) WithLimit(limit int) adapter.EssayQuestionRepository {
	clone := repo.clone()
	clone.limit = &limit
	return clone
}

func (repo *essayQuestionRepository) WithCursor(cursor int) adapter.EssayQuestionRepository {
	clone := repo.clone()
	clone.cursor = &cursor
	return clone
}

// --- 🔧 Query Builder Helper ---

func (repo *essayQuestionRepository) getQueryBuilder() *builder.QueryBuilder[models.EssayQuestion] {
	qb := builder.NewQueryBuilder[models.EssayQuestion](repo.db).
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

func (repo *essayQuestionRepository) CreateEssayQuestion(data *models.EssayQuestion) (*models.EssayQuestion, error) {
	if err := repo.getQueryBuilder().Create(data); err != nil {
		return nil, err
	}
	return data, nil
}

func (repo *essayQuestionRepository) FindEssayQuestionByID(id int64) (*models.EssayQuestion, error) {
	return repo.getQueryBuilder().
		WithWhere(func(db *gorm.DB) *gorm.DB {
			return db.Where("id = ?", id)
		}).
		WithPreloads("CodeQuestion").
		FindOne()
}

func (repo *essayQuestionRepository) FindEssayQuestionsByCodeQuestionID(codeQuestionID int64) ([]models.EssayQuestion, error) {
	return repo.getQueryBuilder().
		WithWhere(func(db *gorm.DB) *gorm.DB {
			return db.Where("code_question_id = ?", codeQuestionID)
		}).
		WithPreloads("CodeQuestion").
		FindAll()
}
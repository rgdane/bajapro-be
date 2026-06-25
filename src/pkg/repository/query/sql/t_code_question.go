package sql

import (
	"jk-api/internal/config"
	"jk-api/internal/database/models"
	adapter "jk-api/pkg/repository/adapter/sql"
	"jk-api/pkg/repository/query/sql/builder"

	"gorm.io/gorm"
)

type tCodeQuestionRepository struct {
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

func NewTCodeQuestionRepository() adapter.TCodeQuestionRepository {
	return &tCodeQuestionRepository{db: config.DB}
}

// --- 🔁 Chainable Configs ---

func (repo *tCodeQuestionRepository) clone() *tCodeQuestionRepository {
	clone := *repo
	return &clone
}

func (repo *tCodeQuestionRepository) WithTx(tx *gorm.DB) adapter.TCodeQuestionRepository {
	clone := repo.clone()
	clone.db = tx
	return clone
}

func (repo *tCodeQuestionRepository) WithPreloads(preloads ...string) adapter.TCodeQuestionRepository {
	clone := repo.clone()
	clone.preloads = append(clone.preloads, preloads...)
	return clone
}

func (repo *tCodeQuestionRepository) WithAssociations(associations ...string) adapter.TCodeQuestionRepository {
	clone := repo.clone()
	clone.associations = append(clone.associations, associations...)
	return clone
}

func (repo *tCodeQuestionRepository) WithReplacements(replacements map[string]interface{}) adapter.TCodeQuestionRepository {
	clone := repo.clone()
	clone.replacements = replacements
	return clone
}

func (repo *tCodeQuestionRepository) WithJoins(joins ...string) adapter.TCodeQuestionRepository {
	clone := repo.clone()
	clone.joins = append(clone.joins, joins...)
	return clone
}

func (repo *tCodeQuestionRepository) WithWhere(query interface{}, args ...interface{}) adapter.TCodeQuestionRepository {
	clone := repo.clone()
	clone.whereClauses = append(clone.whereClauses, func(db *gorm.DB) *gorm.DB {
		return db.Where(query, args...)
	})
	return clone
}

func (repo *tCodeQuestionRepository) WithOrder(order string) adapter.TCodeQuestionRepository {
	clone := repo.clone()
	clone.order = order
	return clone
}

func (repo *tCodeQuestionRepository) WithLimit(limit int) adapter.TCodeQuestionRepository {
	clone := repo.clone()
	clone.limit = &limit
	return clone
}

func (repo *tCodeQuestionRepository) WithCursor(cursor int) adapter.TCodeQuestionRepository {
	clone := repo.clone()
	clone.cursor = &cursor
	return clone
}

// --- 🔧 Query Builder Helper ---

func (repo *tCodeQuestionRepository) getQueryBuilder() *builder.QueryBuilder[models.CodeQuestion] {
	qb := builder.NewQueryBuilder[models.CodeQuestion](repo.db).
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

func (repo *tCodeQuestionRepository) CreateTCodeQuestion(data *models.CodeQuestion) (*models.CodeQuestion, error) {
	if err := repo.getQueryBuilder().Create(data); err != nil {
		return nil, err
	}
	return data, nil
}

func (repo *tCodeQuestionRepository) FindTCodeQuestionByID(id int64) (*models.CodeQuestion, error) {
	return repo.getQueryBuilder().
		WithWhere(func(db *gorm.DB) *gorm.DB {
			return db.Where("id = ?", id)
		}).
		WithPreloads("CodeAnswers", "EssayQuestions", "CodeHistoryLogs").
		FindOne()
}

func (repo *tCodeQuestionRepository) FindTCodeQuestionsBySubLessonID(subLessonID int64) ([]models.CodeQuestion, error) {
	return repo.getQueryBuilder().
		WithWhere(func(db *gorm.DB) *gorm.DB {
			return db.Where("sub_lesson_id = ?", subLessonID)
		}).
		WithPreloads("SubLesson", "CodeAnswers", "EssayQuestions", "CodeHistoryLogs").
		FindAll()
}
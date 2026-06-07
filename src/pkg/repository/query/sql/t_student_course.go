package sql

import (
	"jk-api/internal/config"
	"jk-api/internal/database/models"
	adapter "jk-api/pkg/repository/adapter/sql"
	"jk-api/pkg/repository/query/sql/builder"

	"gorm.io/gorm"
)

type tStudentCourseRepository struct {
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

func NewTStudentCourseRepository() adapter.TStudentCourseRepository {
	return &tStudentCourseRepository{db: config.DB}
}

// --- 🔁 Chainable Configs ---

func (repo *tStudentCourseRepository) clone() *tStudentCourseRepository {
	clone := *repo
	return &clone
}

func (repo *tStudentCourseRepository) WithTx(tx *gorm.DB) adapter.TStudentCourseRepository {
	clone := repo.clone()
	clone.db = tx
	return clone
}

func (repo *tStudentCourseRepository) WithPreloads(preloads ...string) adapter.TStudentCourseRepository {
	clone := repo.clone()
	clone.preloads = append(clone.preloads, preloads...)
	return clone
}

func (repo *tStudentCourseRepository) WithAssociations(associations ...string) adapter.TStudentCourseRepository {
	clone := repo.clone()
	clone.associations = append(clone.associations, associations...)
	return clone
}

func (repo *tStudentCourseRepository) WithReplacements(replacements map[string]interface{}) adapter.TStudentCourseRepository {
	clone := repo.clone()
	clone.replacements = replacements
	return clone
}

func (repo *tStudentCourseRepository) WithJoins(joins ...string) adapter.TStudentCourseRepository {
	clone := repo.clone()
	clone.joins = append(clone.joins, joins...)
	return clone
}

func (repo *tStudentCourseRepository) WithWhere(query interface{}, args ...interface{}) adapter.TStudentCourseRepository {
	clone := repo.clone()
	clone.whereClauses = append(clone.whereClauses, func(db *gorm.DB) *gorm.DB {
		return db.Where(query, args...)
	})
	return clone
}

func (repo *tStudentCourseRepository) WithOrder(order string) adapter.TStudentCourseRepository {
	clone := repo.clone()
	clone.order = order
	return clone
}

func (repo *tStudentCourseRepository) WithLimit(limit int) adapter.TStudentCourseRepository {
	clone := repo.clone()
	clone.limit = &limit
	return clone
}

func (repo *tStudentCourseRepository) WithCursor(cursor int) adapter.TStudentCourseRepository {
	clone := repo.clone()
	clone.cursor = &cursor
	return clone
}

// --- 🔧 Query Builder Helper ---

func (repo *tStudentCourseRepository) getQueryBuilder() *builder.QueryBuilder[models.TStudentCourse] {
	qb := builder.NewQueryBuilder[models.TStudentCourse](repo.db).
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

func (repo *tStudentCourseRepository) EnrollCourse(data *models.TStudentCourse) (*models.TStudentCourse, error) {
	if err := repo.getQueryBuilder().Create(data); err != nil {
		return nil, err
	}
	return data, nil
}

func (repo *tStudentCourseRepository) FindByID(id int64, userID int64) (*models.TStudentCourse, error) {
	return repo.getQueryBuilder().
		WithWhere(func(db *gorm.DB) *gorm.DB {
			return db.Where("id = ? AND user_id = ?", id, userID)
		}).
		WithPreloads("Course.Lessons.SubLessons.Materials", "Badge").
		FindOne()
}

func (repo *tStudentCourseRepository) FindMyCourse(userID int64) ([]models.TStudentCourse, error) {
	return repo.getQueryBuilder().
		WithWhere(func(db *gorm.DB) *gorm.DB {
			return db.Where("user_id = ?", userID) // ✅ FIX
		}).
		WithPreloads("Course", "Badge").
		FindAll()
}
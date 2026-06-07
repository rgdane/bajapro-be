package sql

import (
	"jk-api/internal/config"
	"jk-api/internal/database/models"
	adapter "jk-api/pkg/repository/adapter/sql"
	"jk-api/pkg/repository/query/sql/builder"

	"gorm.io/gorm"
)

type tStudentProgressRepository struct {
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

func NewTStudentProgressRepository() adapter.TStudentProgressRepository {
	return &tStudentProgressRepository{db: config.DB}
}

// --- 🔁 Chainable Configs ---

func (repo *tStudentProgressRepository) clone() *tStudentProgressRepository {
	clone := *repo
	return &clone
}

func (repo *tStudentProgressRepository) WithTx(tx *gorm.DB) adapter.TStudentProgressRepository {
	clone := repo.clone()
	clone.db = tx
	return clone
}

func (repo *tStudentProgressRepository) WithPreloads(preloads ...string) adapter.TStudentProgressRepository {
	clone := repo.clone()
	clone.preloads = append(clone.preloads, preloads...)
	return clone
}

func (repo *tStudentProgressRepository) WithAssociations(associations ...string) adapter.TStudentProgressRepository {
	clone := repo.clone()
	clone.associations = append(clone.associations, associations...)
	return clone
}

func (repo *tStudentProgressRepository) WithReplacements(replacements map[string]interface{}) adapter.TStudentProgressRepository {
	clone := repo.clone()
	clone.replacements = replacements
	return clone
}

func (repo *tStudentProgressRepository) WithJoins(joins ...string) adapter.TStudentProgressRepository {
	clone := repo.clone()
	clone.joins = append(clone.joins, joins...)
	return clone
}

func (repo *tStudentProgressRepository) WithWhere(query interface{}, args ...interface{}) adapter.TStudentProgressRepository {
	clone := repo.clone()
	clone.whereClauses = append(clone.whereClauses, func(db *gorm.DB) *gorm.DB {
		return db.Where(query, args...)
	})
	return clone
}

func (repo *tStudentProgressRepository) WithOrder(order string) adapter.TStudentProgressRepository {
	clone := repo.clone()
	clone.order = order
	return clone
}

func (repo *tStudentProgressRepository) WithLimit(limit int) adapter.TStudentProgressRepository {
	clone := repo.clone()
	clone.limit = &limit
	return clone
}

func (repo *tStudentProgressRepository) WithCursor(cursor int) adapter.TStudentProgressRepository {
	clone := repo.clone()
	clone.cursor = &cursor
	return clone
}

// --- 🔧 Query Builder Helper ---

func (repo *tStudentProgressRepository) getQueryBuilder() *builder.QueryBuilder[models.TStudentProgress] {
	qb := builder.NewQueryBuilder[models.TStudentProgress](repo.db).
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

func (repo *tStudentProgressRepository) CompleteTStudentProgress(data *models.TStudentProgress) (*models.TStudentProgress, error) {
	if err := repo.getQueryBuilder().Create(data); err != nil {
		return nil, err
	}
	return data, nil
}

func (repo *tStudentProgressRepository) FindByUserAndSubLesson(
	userID int64,
	subLessonID int64,
) (*models.TStudentProgress, error) {

	return repo.getQueryBuilder().
		WithWhere(func(db *gorm.DB) *gorm.DB {
			return db.
				Where("user_id = ?", userID).
				Where("sub_lesson_id = ?", subLessonID)
		}).
		FindFirst()
}

func (repo *tStudentProgressRepository) CountTotalSubLessonByCourse(
	courseID int64,
) (int64, error) {

	var total int64

	err := repo.db.
		Table("m_sub_lesson").
		Joins("JOIN m_lesson ON m_lesson.id = m_sub_lesson.lesson_id").
		Where("m_lesson.course_id = ?", courseID).
		Count(&total).
		Error

	if err != nil {
		return 0, err
	}

	return total, nil
}

func (repo *tStudentProgressRepository) CountCompletedByCourse(
	userID int64,
	courseID int64,
) (int64, error) {

	var total int64

	err := repo.db.
		Table("t_student_progress").
		Joins(`
			JOIN m_sub_lesson
				ON m_sub_lesson.id = t_student_progress.sub_lesson_id
		`).
		Joins(`
			JOIN m_lesson
				ON m_lesson.id = m_sub_lesson.lesson_id
		`).
		Where("t_student_progress.user_id = ?", userID).
		Where("m_lesson.course_id = ?", courseID).
		Where("t_student_progress.status = ?", "completed").
		Count(&total).
		Error

	if err != nil {
		return 0, err
	}

	return total, nil
}

func (repo *tStudentProgressRepository) FindCourseIDBySubLesson(
	subLessonID int64,
) (int64, error) {

	var courseID int64

	err := repo.db.
		Table("m_sub_lesson").
		Select("m_lesson.course_id").
		Joins(
			"JOIN m_lesson ON m_lesson.id = m_sub_lesson.lesson_id",
		).
		Where("m_sub_lesson.id = ?", subLessonID).
		Scan(&courseID).
		Error

	return courseID, err
}

func (repo *tStudentProgressRepository) UpdateStudentCourseProgress(
	userID int64,
	courseID int64,
	percentage float64,
) error {

	return repo.db.
		Model(&models.TStudentCourse{}).
		Where("user_id = ?", userID).
		Where("course_id = ?", courseID).
		Update("progress_percentage", percentage).
		Error
}
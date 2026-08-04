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

func (repo *tStudentCourseRepository) FindByUserIDAndCourseID(userID int64, courseID int64) (*models.TStudentCourse, error) {
	return repo.getQueryBuilder().
		WithWhere(func(db *gorm.DB) *gorm.DB {
			return db.Where("user_id = ? AND course_id = ?", userID, courseID)
		}).
		WithPreloads("Course", "Badge").
		FindOne()
}

func (repo *tStudentCourseRepository) UpdateTotalScoreByID(
	data *models.TStudentCourse,
) (*models.TStudentCourse, error) {

	err := repo.db.
		Model(&models.TStudentCourse{}).
		Where("id = ?", data.ID).
		Updates(map[string]interface{}{
			"total_score": data.TotalScore,
			"badge_id":    data.BadgeID,
		}).Error

	if err != nil {
		return nil, err
	}

	return data, nil
}
func (repo *tStudentCourseRepository) FindLeaderboard(
	courseID int64,
	classID *int64,
) ([]models.Leaderboard, error) {
	var leaderboard []models.Leaderboard

	query := `
		WITH enrolled_students AS (
			SELECT DISTINCT
				tc.user_id,
				tc.badge_id
			FROM t_student_course tc
			WHERE tc.course_id = ?
		),

		coding_scores AS (
			SELECT
				answer.user_id,
				SUM(answer.best_score) AS coding_score
			FROM (
				SELECT
					tca.user_id,
					tca.code_question_id,
					MAX(tca.exploring_score) AS best_score
				FROM t_code_answer tca
				INNER JOIN code_question cq
					ON cq.id = tca.code_question_id
				INNER JOIN m_sub_lesson sl
					ON sl.id = cq.sub_lesson_id
				INNER JOIN m_lesson l
					ON l.id = sl.lesson_id
				WHERE l.course_id = ?
				GROUP BY
					tca.user_id,
					tca.code_question_id
			) answer
			GROUP BY answer.user_id
		),

		essay_scores AS (
			SELECT
				answer.user_id,
				SUM(answer.best_score) AS essay_score
			FROM (
				SELECT
					tea.user_id,
					tea.essay_question_id,
					MAX(
						COALESCE(tea.konteks_penjelasan, 0)
						+ COALESCE(tea.keruntutan, 0)
						+ COALESCE(tea.kebenaran, 0)
					) AS best_score
				FROM t_essay_answer tea
				INNER JOIN essay_question eq
					ON eq.id = tea.essay_question_id
				INNER JOIN m_sub_lesson sl
					ON sl.id = eq.sub_lesson_id
				INNER JOIN m_lesson l
					ON l.id = sl.lesson_id
				WHERE l.course_id = ?
				GROUP BY
					tea.user_id,
					tea.essay_question_id
			) answer
			GROUP BY answer.user_id
		)

		SELECT
			u.id AS user_id,
			u.name,

			COALESCE(b.name, '') AS badge_name,
			COALESCE(b.image, '') AS badge_image,

			0 AS reading_score,
			COALESCE(cs.coding_score, 0) AS coding_score,
			COALESCE(es.essay_score, 0) AS essay_score,

			(
				COALESCE(cs.coding_score, 0)
				+ COALESCE(es.essay_score, 0)
			) AS total_score

		FROM enrolled_students enrolled

		INNER JOIN users u
			ON u.id = enrolled.user_id

		INNER JOIN user_has_roles uhr
			ON uhr.user_id = u.id

		INNER JOIN roles r
			ON r.id = uhr.role_id
			AND LOWER(r.name) = 'student'

		LEFT JOIN m_badge_settings b
			ON b.id = enrolled.badge_id

		LEFT JOIN coding_scores cs
			ON cs.user_id = u.id

		LEFT JOIN essay_scores es
			ON es.user_id = u.id

		WHERE 1 = 1
	`

	args := []any{
		courseID,
		courseID,
		courseID,
	}

	if classID != nil {
		query += ` AND u.class_id = ?`
		args = append(args, *classID)
	}

	query += `
		ORDER BY
			total_score DESC,
			u.name ASC
	`

	err := repo.db.
		Raw(query, args...).
		Scan(&leaderboard).
		Error

	if err != nil {
		return nil, err
	}

	return leaderboard, nil
}

func (repo *tStudentCourseRepository) FindAllScore() ([]models.TStudentCourse, error) {
	return repo.
		getQueryBuilder().
		FindAll()
}
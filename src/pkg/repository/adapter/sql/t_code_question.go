package sql

import (
	"jk-api/internal/database/models"

	"gorm.io/gorm"
)

type TCodeQuestionRepository interface {
	WithTx(tx *gorm.DB) TCodeQuestionRepository
	WithPreloads(preloads ...string) TCodeQuestionRepository
	WithAssociations(associations ...string) TCodeQuestionRepository
	WithReplacements(replacements map[string]interface{}) TCodeQuestionRepository
	WithJoins(joins ...string) TCodeQuestionRepository
	WithWhere(query interface{}, args ...interface{}) TCodeQuestionRepository
	WithOrder(order string) TCodeQuestionRepository
	WithLimit(limit int) TCodeQuestionRepository
	WithCursor(cursor int) TCodeQuestionRepository

	FindTCodeQuestionByID(id int64) (*models.CodeQuestion, error)
	FindTCodeQuestionsBySubLessonID(subLessonID int64) ([]models.CodeQuestion, error)
	CreateTCodeQuestion(data *models.CodeQuestion) (*models.CodeQuestion, error)
}

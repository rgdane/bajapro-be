package sql

import (
	"jk-api/internal/database/models"

	"gorm.io/gorm"
)

type TCodeAnswerRepository interface {
	WithTx(tx *gorm.DB) TCodeAnswerRepository
	WithPreloads(preloads ...string) TCodeAnswerRepository
	WithAssociations(associations ...string) TCodeAnswerRepository
	WithReplacements(replacements map[string]interface{}) TCodeAnswerRepository
	WithJoins(joins ...string) TCodeAnswerRepository
	WithWhere(query interface{}, args ...interface{}) TCodeAnswerRepository
	WithOrder(order string) TCodeAnswerRepository
	WithLimit(limit int) TCodeAnswerRepository
	WithCursor(cursor int) TCodeAnswerRepository

	FindTCodeAnswersByCodeQuestionID(codeQuestionID int64) ([]models.TCodeAnswer, error)
	CreateTCodeAnswer(data *models.TCodeAnswer) (*models.TCodeAnswer, error)
}

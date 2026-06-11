package sql

import (
	"jk-api/internal/database/models"

	"gorm.io/gorm"
)

type EssayQuestionRepository interface {
	WithTx(tx *gorm.DB) EssayQuestionRepository
	WithPreloads(preloads ...string) EssayQuestionRepository
	WithAssociations(associations ...string) EssayQuestionRepository
	WithReplacements(replacements map[string]interface{}) EssayQuestionRepository
	WithJoins(joins ...string) EssayQuestionRepository
	WithWhere(query interface{}, args ...interface{}) EssayQuestionRepository
	WithOrder(order string) EssayQuestionRepository
	WithLimit(limit int) EssayQuestionRepository
	WithCursor(cursor int) EssayQuestionRepository

	FindEssayQuestionByID(id int64) (*models.EssayQuestion, error)
	FindEssayQuestionsByCodeQuestionID(codeQuestionID int64) ([]models.EssayQuestion, error)
	CreateEssayQuestion(data *models.EssayQuestion) (*models.EssayQuestion, error)
}

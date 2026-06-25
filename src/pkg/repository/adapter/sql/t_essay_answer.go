package sql

import (
	"jk-api/internal/database/models"

	"gorm.io/gorm"
)

type TEssayAnswerRepository interface {
	WithTx(tx *gorm.DB) TEssayAnswerRepository
	WithPreloads(preloads ...string) TEssayAnswerRepository
	WithAssociations(associations ...string) TEssayAnswerRepository
	WithReplacements(replacements map[string]interface{}) TEssayAnswerRepository
	WithJoins(joins ...string) TEssayAnswerRepository
	WithWhere(query interface{}, args ...interface{}) TEssayAnswerRepository
	WithOrder(order string) TEssayAnswerRepository
	WithLimit(limit int) TEssayAnswerRepository
	WithCursor(cursor int) TEssayAnswerRepository

	FindTEssayAnswersByEssayQuestionIDAndUserID(essayQuestionID, userID int64) (*models.TEssayAnswer, error)
	CreateTEssayAnswer(data *models.TEssayAnswer) (*models.TEssayAnswer, error)
}

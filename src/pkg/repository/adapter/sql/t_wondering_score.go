package sql

import (
	"jk-api/internal/database/models"

	"gorm.io/gorm"
)

type TWonderingScoreRepository interface {
	WithTx(tx *gorm.DB) TWonderingScoreRepository
	WithPreloads(preloads ...string) TWonderingScoreRepository
	WithAssociations(associations ...string) TWonderingScoreRepository
	WithReplacements(replacements map[string]interface{}) TWonderingScoreRepository
	WithJoins(joins ...string) TWonderingScoreRepository
	WithWhere(query interface{}, args ...interface{}) TWonderingScoreRepository
	WithOrder(order string) TWonderingScoreRepository
	WithLimit(limit int) TWonderingScoreRepository
	WithCursor(cursor int) TWonderingScoreRepository

	FindTWonderingScoresBySubLessonIDAndUserID(subLessonID, userID int64) (*models.TWonderingScore, error)
	CreateTWonderingScore(data *models.TWonderingScore) (*models.TWonderingScore, error)
}

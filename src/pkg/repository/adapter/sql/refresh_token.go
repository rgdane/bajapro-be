package sql

import (
	"jk-api/internal/database/models"

	"gorm.io/gorm"
)

type RefreshTokenRepository interface {
	WithTx(tx *gorm.DB) RefreshTokenRepository
	WithWhere(query interface{}, args ...interface{}) RefreshTokenRepository
	WithOrder(order string) RefreshTokenRepository
	WithLimit(limit int) RefreshTokenRepository
	WithUnscoped() RefreshTokenRepository

	Insert(data *models.RefreshToken) (*models.RefreshToken, error)
	FindByToken(token string) (*models.RefreshToken, error)
	DeleteByToken(token string) error
	DeleteByUserID(userID int64) error
	DeleteExpired() error
}
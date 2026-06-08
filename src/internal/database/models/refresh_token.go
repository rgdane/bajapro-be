package models

import (
	"time"

)

type RefreshToken struct {
	ID        int64     `gorm:"primaryKey"`
	UserID    int64     `gorm:"index"`
	Token     string    `gorm:"type:text;unique"`
	ExpiresAt time.Time
	CreatedAt time.Time
}

func (*RefreshToken) TableName() string {
	return "refresh_tokens"
}

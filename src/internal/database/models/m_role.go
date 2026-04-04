package models

import "time"

type MRole struct {
	ID         int64      `gorm:"primaryKey;autoIncrement:true;type:serial" json:"id"`
	RoleName   string     `gorm:"size:50;not null" json:"role_name"`
	IsActive   bool       `gorm:"default:true" json:"isactive"`
	CreatedAt  time.Time  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	UserCreate int64      `json:"user_create"`
	UserUpdate int64      `json:"user_update"`
}

func (*MRole) TableName() string {
	return "m_role"
}

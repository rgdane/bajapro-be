package models

import "time"

type MBadgeSettings struct {
	ID         int64      `gorm:"primaryKey;autoIncrement:true" json:"id"`
	Name       string     `gorm:"size:100" json:"name"`
	Image      string     `gorm:"type:text" json:"image"`
	MinScore   int        `gorm:"column:min_score" json:"min_score"`
	MaxScore   int        `gorm:"column:max_score" json:"max_score"`
	IsActive   bool       `gorm:"column:isactive;default:true" json:"isactive"`
	CreatedAt  time.Time  `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt  time.Time  `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`

	
}

func (*MBadgeSettings) TableName() string {
	return "m_badge_settings"
}

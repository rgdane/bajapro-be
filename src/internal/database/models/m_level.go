package models

import "time"

type MLevel struct {
	ID          int64      `gorm:"primaryKey;autoIncrement:true;type:serial" json:"id"`
	LevelName   string     `gorm:"column:level_name;size:50" json:"level_name"`
	Description string     `gorm:"type:text" json:"description"`
	IsActive    bool       `gorm:"column:isactive;default:true" json:"isactive"`
	CreatedAt   time.Time  `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
 	DeletedAt   time.Time  `gorm:"column:deleted_at;index" json:"deleted_at,omitempty"`
}

func (*MLevel) TableName() string {
	return "m_levels"
}

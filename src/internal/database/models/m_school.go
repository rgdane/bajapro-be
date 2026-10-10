package models

import (
	"time"

	"gorm.io/gorm"
)

type MSchool struct {
	ID         int64          `gorm:"primaryKey;autoIncrement:true" json:"id"`
	SchoolName string         `gorm:"column:school_name;size:100;not null" json:"school_name"`
	IsActive   bool           `gorm:"column:isactive;default:true" json:"isactive"`
	CreatedAt  time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt  *time.Time     `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index:idx_school_deleted_at" json:"deleted_at"`

	// Relationships
	Classes []MClass `gorm:"foreignKey:SchoolID;references:ID" json:"classes"`
	Users   []User   `gorm:"foreignKey:SchoolID;references:ID" json:"users"`
}

func (*MSchool) TableName() string {
	return "m_school"
}

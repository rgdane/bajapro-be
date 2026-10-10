package models

import (
	"time"

	"gorm.io/gorm"
)

type MClass struct {
	ID        int64          `gorm:"primaryKey;autoIncrement:true" json:"id"`
	SchoolID  int64          `gorm:"column:school_id;not null" json:"school_id"`
	TeacherID *int64         `gorm:"column:teacher_id" json:"teacher_id"`
	ClassName string         `gorm:"column:class_name;size:100;not null" json:"class_name"`
	ClassCode string         `gorm:"column:class_code;size:50;uniqueIndex;not null" json:"class_code"`
	IsActive  bool           `gorm:"column:isactive;default:true" json:"isactive"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt *time.Time     `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index:idx_class_deleted_at" json:"deleted_at"`

	// Relationships
	School   *MSchool `gorm:"foreignKey:SchoolID;references:ID" json:"school,omitempty"`
	Teacher  *User    `gorm:"foreignKey:TeacherID;references:ID;-:migration" json:"teacher,omitempty"`
	Students []User   `gorm:"foreignKey:ClassID;-:migration" json:"students,omitempty"`
}

func (*MClass) TableName() string {
	return "m_class"
}

package models

import "time"

type MCourse struct {
	ID           int64      `gorm:"primaryKey;autoIncrement:true" json:"id"`
	CourseName   string     `gorm:"column:course_name;size:100" json:"course_name"`
	Description  string     `gorm:"type:text" json:"description"`
	ImgThumbnail string     `gorm:"column:img_thumbnail;type:text" json:"img_thumbnail"`
	Published    bool       `gorm:"default:false" json:"published"`
	IsActive     bool       `gorm:"column:isactive;default:true" json:"isactive"`
	CreatedAt    time.Time  `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt    time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
	DeletedAt    time.Time `gorm:"column:deleted_at;index" json:"deleted_at,omitempty"`

	// Relations
	Lessons *[]MLesson `gorm:"foreignKey:CourseID;references:ID" json:"lessons"`
}

func (*MCourse) TableName() string {
	return "m_course"
}

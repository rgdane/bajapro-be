package models

import "time"

type MLesson struct {
	ID           int64      `gorm:"primaryKey;autoIncrement:true;type:serial" json:"id"`
	CourseID     int64      `gorm:"column:course_id" json:"course_id"`
	LevelID      int64      `gorm:"column:level_id" json:"level_id"`
	Title        string     `gorm:"size:100" json:"title"`
	Description  string     `gorm:"type:text" json:"description"`
	Position     int        `json:"position"`
	ImgThumbnail string     `gorm:"column:img_thumbnail;type:text" json:"img_thumbnail"`
	Published    bool       `gorm:"default:false" json:"published"`
	IsActive     bool       `gorm:"column:isactive;default:true" json:"isactive"`
	CreatedAt    time.Time  `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt    *time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
	UserCreate   int64      `gorm:"column:user_create" json:"user_create"`
	UserUpdate   int64      `gorm:"column:user_update" json:"user_update"`

	// Foreign Key Relationships
	Course *MCourse `gorm:"foreignKey:CourseID;references:ID" json:"course"`
	Level  *MLevel  `gorm:"foreignKey:LevelID;references:ID" json:"level"`
}

func (*MLesson) TableName() string {
	return "m_lesson"
}

package models

import "time"

type TGenerationHistory struct {
	ID             int64      `gorm:"primaryKey;autoIncrement:true" json:"id"`
	SubLessonID    int64      `gorm:"column:sub_lesson_id" json:"sub_lesson_id"`
	TopicUsed      string     `gorm:"column:topic_used;type:text" json:"topic_used"`
	Result         string     `gorm:"type:text" json:"result"`
	GenerationTime int        `gorm:"column:generation_time" json:"generation_time"`
	IsActive       bool       `gorm:"column:isactive;default:true" json:"isactive"`
	CreatedAt      time.Time  `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt      *time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
	// Foreign Key Relationships
	SubLesson *MSubLesson `gorm:"foreignKey:SubLessonID;references:ID" json:"sub_lesson"`
}

func (*TGenerationHistory) TableName() string {
	return "t_generation_history"
}

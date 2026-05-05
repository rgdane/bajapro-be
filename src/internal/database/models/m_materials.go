package models

import "time"

type MMaterials struct {
	ID              int64      `gorm:"primaryKey;autoIncrement:true" json:"id"`
	SubLessonID    int64      `gorm:"column:sub_lesson_id" json:"sub_lesson_id"`
	Title           string     `gorm:"size:150" json:"title"`
	Materials       string     `gorm:"type:text" json:"materials"`
	URLVideo        string     `gorm:"column:url_video;type:text" json:"url_video"`
	ContentPosition int        `gorm:"column:content_position" json:"content_position"`
	PromptLLM       string     `gorm:"column:prompt_llm;type:text" json:"prompt_llm"`
	Published       bool       `gorm:"default:false" json:"published"`
	IsActive        bool       `gorm:"column:isactive;default:true" json:"isactive"`
	CreatedAt       time.Time  `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt       time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
	DeletedAt       time.Time  `gorm:"column:deleted_at;index" json:"deleted_at,omitempty"`

	SubLesson *MSubLesson `gorm:"foreignKey:SubLessonID;references:ID" json:"sub_lesson"`
}

func (*MMaterials) TableName() string {
	return "m_materials"
}

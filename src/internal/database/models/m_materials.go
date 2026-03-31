package models

import "time"

type MMaterials struct {
	ID              int64      `gorm:"primaryKey;autoIncrement:true;type:serial" json:"id"`
	SubSessionID    int64      `gorm:"column:sub_session_id" json:"sub_session_id"`
	Title           string     `gorm:"size:150" json:"title"`
	Materials       string     `gorm:"type:text" json:"materials"`
	URLVideo        string     `gorm:"column:url_video;type:text" json:"url_video"`
	ContentPosition int        `gorm:"column:content_position" json:"content_position"`
	PromptLLM       string     `gorm:"column:prompt_llm;type:text" json:"prompt_llm"`
	Published       bool       `gorm:"default:false" json:"published"`
	IsActive        bool       `gorm:"column:isactive;default:true" json:"isactive"`
	CreatedAt       time.Time  `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt       *time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
	UserCreate      int64      `gorm:"column:user_create" json:"user_create"`
	UserUpdate      int64      `gorm:"column:user_update" json:"user_update"`

	// Foreign Key Relationships
	SubLesson *MSubLesson `gorm:"foreignKey:SubSessionID;references:ID" json:"sub_lesson"`
}

func (*MMaterials) TableName() string {
	return "m_materials"
}

package models

import "time"

type MSubLesson struct {
	ID            int64      `gorm:"primaryKey;autoIncrement:true" json:"id"`
	LessonID      int64      `gorm:"column:lesson_id" json:"lesson_id"`
	Title         string     `gorm:"size:100" json:"title"`
	OrderPosition int        `gorm:"column:order_position" json:"order_position"`
	IsActive      bool       `gorm:"column:isactive;default:true" json:"isactive"`
	CreatedAt     time.Time  `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt     time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
	DeletedAt     time.Time  `gorm:"column:deleted_at;index" json:"deleted_at,omitempty"`

	// Foreign Key Relationships
	Lesson *MLesson `gorm:"foreignKey:LessonID;references:ID" json:"lesson"`
	Materials []MMaterials `gorm:"foreignKey:SubLessonID;references:ID" json:"materials"`
	CodeQuestions []CodeQuestion `gorm:"foreignKey:SubLessonID;references:ID" json:"code_questions"`
}

func (*MSubLesson) TableName() string {
	return "m_sub_lesson"
}

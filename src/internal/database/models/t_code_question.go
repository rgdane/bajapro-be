package models

import "time"

type TCodeQuestion struct {
	ID           int64      `gorm:"primaryKey;autoIncrement:true;type:serial" json:"id"`
	SubLessonID  int64      `gorm:"column:sub_lesson_id" json:"sub_lesson_id"`
	CodeQuestion string     `gorm:"column:code_question;type:text" json:"code_question"`
	Image        string     `gorm:"type:text" json:"image"`
	Score        int        `gorm:"column:score" json:"score"`
	Hint         string     `gorm:"type:text" json:"hint"`
	IsActive     bool       `gorm:"column:isactive;default:true" json:"isactive"`
	CreatedAt    time.Time  `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt    *time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
	UserCreate   int64      `gorm:"column:user_create" json:"user_create"`
	UserUpdate   int64      `gorm:"column:user_update" json:"user_update"`

	// Foreign Key Relationships
	SubLesson *MSubLesson `gorm:"foreignKey:SubLessonID;references:ID" json:"sub_lesson"`
}

func (*TCodeQuestion) TableName() string {
	return "t_code_question"
}

package models

import "time"

type EssayQuestion struct {
	ID             int64      `gorm:"primaryKey;autoIncrement:true" json:"id"`
	SubLessonID    int64     `gorm:"column:sub_lesson_id" json:"sub_lesson_id"`
	CodeQuestionID int64      `gorm:"column:code_question_id" json:"code_question_id"`
	EssayQuestion  string     `gorm:"column:essay_question;type:text" json:"essay_question"`
	Answer         string     `gorm:"column:answer;type:text" json:"answer"`
	Answer2        string     `gorm:"column:answer_2;type:text" json:"answer_2"`
	Answer3        string     `gorm:"column:answer_3;type:text" json:"answer_3"`
	Answer4        string     `gorm:"column:answer_4;type:text" json:"answer_4"`
	IsActive       bool       `gorm:"column:isactive;default:true" json:"isactive"`
	CreatedAt      time.Time  `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt      time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`

	// Foreign Key Relationships
	SubLesson *MSubLesson `gorm:"foreignKey:SubLessonID;references:ID" json:"sub_lesson"`
	CodeQuestion *CodeQuestion `gorm:"foreignKey:CodeQuestionID;references:ID" json:"code_question"`
	EssayAnswers []TEssayAnswer `gorm:"foreignKey:EssayQuestionID;references:ID" json:"essay_answers"`
}

func (*EssayQuestion) TableName() string {
	return "t_essay_question"
}

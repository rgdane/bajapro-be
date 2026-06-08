package models

import "time"

type TEssayQuestion struct {
	ID             int64      `gorm:"primaryKey;autoIncrement:true;type:serial" json:"id"`
	CodeQuestionID int64      `gorm:"column:code_question_id" json:"code_question_id"`
	EssayQuestion  string     `gorm:"column:essay_question;type:text" json:"essay_question"`
	Answer         string     `gorm:"type:text" json:"answer"`
	Answer1        string     `gorm:"column:answer_1;type:text" json:"answer_1"`
	Answer2        string     `gorm:"column:answer_2;type:text" json:"answer_2"`
	Answer3        string     `gorm:"column:answer_3;type:text" json:"answer_3"`
	Answer4        string     `gorm:"column:answer_4;type:text" json:"answer_4"`
	IsActive       bool       `gorm:"column:isactive;default:true" json:"isactive"`
	CreatedAt      time.Time  `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt      *time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
	UserCreate     int64      `gorm:"column:user_create" json:"user_create"`
	UserUpdate     int64      `gorm:"column:user_update" json:"user_update"`

	// Foreign Key Relationships
	CodeQuestion *TCodeQuestion `gorm:"foreignKey:CodeQuestionID;references:ID" json:"code_question"`
}

func (*TEssayQuestion) TableName() string {
	return "t_essay_question"
}

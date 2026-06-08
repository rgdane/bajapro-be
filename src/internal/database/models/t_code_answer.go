package models

import "time"

type TCodeAnswer struct {
	ID             int64      `gorm:"primaryKey;autoIncrement:true;type:serial" json:"id"`
	UserID         int64      `gorm:"column:user_id" json:"user_id"`
	CodeQuestionID int64      `gorm:"column:code_question_id" json:"code_question_id"`
	IsCodeRight    bool       `gorm:"column:is_code_right" json:"is_code_right"`
	ExploringScore int        `gorm:"column:exploring_score" json:"exploring_score"`
	CreatedAt      time.Time  `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt      *time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`

	// Foreign Key Relationships
	User         *User         `gorm:"foreignKey:UserID;references:ID" json:"user"`
	CodeQuestion *TCodeQuestion `gorm:"foreignKey:CodeQuestionID;references:ID" json:"code_question"`
}

func (*TCodeAnswer) TableName() string {
	return "t_code_answer"
}

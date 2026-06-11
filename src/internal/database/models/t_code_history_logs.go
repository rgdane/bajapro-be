package models

import "time"

type TCodeHistoryLogs struct {
	ID             int64      `gorm:"primaryKey;autoIncrement:true" json:"id"`
	UserID         int64      `gorm:"column:user_id" json:"user_id"`
	CodeQuestionID int64      `gorm:"column:code_question_id" json:"code_question_id"`
	TimeCount      int        `gorm:"column:time_count" json:"time_count"`
	Message        string     `gorm:"type:text" json:"message"`
	IsError        bool       `gorm:"column:is_error" json:"is_error"`
	CreatedAt      time.Time  `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt      *time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`

	// Foreign Key Relationships
	User         *User         `gorm:"foreignKey:UserID;references:ID" json:"user"`
	CodeQuestion *CodeQuestion `gorm:"foreignKey:CodeQuestionID;references:ID" json:"code_question"`
}

func (*TCodeHistoryLogs) TableName() string {
	return "t_code_history_logs"
}

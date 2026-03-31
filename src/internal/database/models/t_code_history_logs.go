package models

import "time"

type TCodeHistoryLogs struct {
	ID             int64      `gorm:"primaryKey;autoIncrement:true;type:serial" json:"id"`
	UserID         int64      `gorm:"column:user_id" json:"user_id"`
	CodeQuestionID int64      `gorm:"column:code_question_id" json:"code_question_id"`
	TimeCount      int        `gorm:"column:time_count" json:"time_count"`
	Message        string     `gorm:"type:text" json:"message"`
	IsError        bool       `gorm:"column:is_error" json:"is_error"`
	IsActive       bool       `gorm:"column:isactive;default:true" json:"isactive"`
	CreatedAt      time.Time  `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt      *time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
	UserCreate     int64      `gorm:"column:user_create" json:"user_create"`
	UserUpdate     int64      `gorm:"column:user_update" json:"user_update"`
	// Foreign Key Relationships
	User         *MUsers        `gorm:"foreignKey:UserID;references:ID" json:"user"`
	CodeQuestion *TCodeQuestion `gorm:"foreignKey:CodeQuestionID;references:ID" json:"code_question"`
}

func (*TCodeHistoryLogs) TableName() string {
	return "t_code_history_logs"
}

package models

import "time"

type TWonderingScore struct {
	ID          int64      `gorm:"primaryKey;autoIncrement:true;type:serial" json:"id"`
	UserID      int64      `gorm:"column:user_id" json:"user_id"`
	SubLessonID int64      `gorm:"column:sub_lesson_id" json:"sub_lesson_id"`
	Score       int        `gorm:"column:score" json:"score"`
	IsActive    bool       `gorm:"column:isactive;default:true" json:"isactive"`
	CreatedAt   time.Time  `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt   *time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
	UserCreate  int64      `gorm:"column:user_create" json:"user_create"`
	UserUpdate  int64      `gorm:"column:user_update" json:"user_update"`

	// Foreign Key Relationships
	User      *User      `gorm:"foreignKey:UserID;references:ID" json:"user"`
	SubLesson *MSubLesson `gorm:"foreignKey:SubLessonID;references:ID" json:"sub_lesson"`
}

func (*TWonderingScore) TableName() string {
	return "t_wondering_score"
}

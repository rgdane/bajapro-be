package models

import "time"

type MSubLesson struct {
	ID            int64      `gorm:"primaryKey;autoIncrement:true;type:serial" json:"id"`
	LessonID      int64      `gorm:"column:lesson_id" json:"lesson_id"`
	Title         string     `gorm:"size:100" json:"title"`
	OrderPosition int        `gorm:"column:order_position" json:"order_position"`
	IsActive      bool       `gorm:"column:isactive;default:true" json:"isactive"`
	CreatedAt     time.Time  `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt     *time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
	UserCreate    int64      `gorm:"column:user_create" json:"user_create"`
	UserUpdate    int64      `gorm:"column:user_update" json:"user_update"`

	// Foreign Key Relationships
	Lesson *MLesson `gorm:"foreignKey:LessonID;references:ID" json:"lesson"`
}

func (*MSubLesson) TableName() string {
	return "m_sub_lesson"
}

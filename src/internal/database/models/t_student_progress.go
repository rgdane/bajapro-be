package models

import "time"

type TStudentProgress struct {
	ID          int64      `gorm:"primaryKey;autoIncrement:true;type:serial" json:"id"`
	UserID      int64      `gorm:"column:user_id" json:"user_id"`
	SubLessonID int64      `gorm:"column:sub_lesson_id" json:"sub_lesson_id"`
	Status      string     `gorm:"size:50" json:"status"`
	CreatedAt   time.Time  `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt   *time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`

	// Foreign Key Relationships
	User      *User      `gorm:"foreignKey:UserID;references:ID" json:"user"`
	SubLesson *MSubLesson `gorm:"foreignKey:SubLessonID;references:ID" json:"sub_lesson"`
}

func (*TStudentProgress) TableName() string {
	return "t_student_progress"
}

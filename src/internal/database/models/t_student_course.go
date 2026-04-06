package models

import "time"

type TStudentCourse struct {
	ID         int64      `gorm:"primaryKey;autoIncrement:true;type:serial" json:"id"`
	UserID     int64      `gorm:"column:user_id" json:"user_id"`
	StudentID  int64      `gorm:"column:student_id" json:"student_id"`
	TotalScore int        `gorm:"column:total_score" json:"total_score"`
	BadgeID    int64      `gorm:"column:badge_id" json:"badge_id"`
	IsActive   bool       `gorm:"column:isactive;default:true" json:"isactive"`
	CreatedAt  time.Time  `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt  *time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
	UserCreate int64      `gorm:"column:user_create" json:"user_create"`
	UserUpdate int64      `gorm:"column:user_update" json:"user_update"`

	// Foreign Key Relationships
	User    *MUser          `gorm:"foreignKey:UserID;references:ID" json:"user"`
	Student *MUser          `gorm:"foreignKey:StudentID;references:ID" json:"student"`
	Badge   *MBadgeSettings `gorm:"foreignKey:BadgeID;references:ID" json:"badge"`
}

func (*TStudentCourse) TableName() string {
	return "t_student_course"
}

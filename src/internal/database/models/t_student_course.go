package models

import "time"

type TStudentCourse struct {
	ID                 int64      `gorm:"primaryKey;autoIncrement:true" json:"id"`
	UserID             int64      `gorm:"column:user_id" json:"user_id"`
	CourseID           int64      `gorm:"column:course_id" json:"course_id"`
	ProgressPercentage float64    `gorm:"column:progress_percentage" json:"progress_percentage"`
	TotalScore         int        `gorm:"column:total_score" json:"total_score"`
	BadgeID            int64      `gorm:"column:badge_id" json:"badge_id"`
	CreatedAt          time.Time  `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt          *time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
	DeletedAt          *time.Time `gorm:"column:deleted_at" json:"deleted_at"`

	// Foreign Key Relationships
	User   *User           `gorm:"foreignKey:UserID;references:ID" json:"user"`
	Course *MCourse        `gorm:"foreignKey:CourseID;references:ID" json:"course"`
	Badge  *MBadgeSettings `gorm:"foreignKey:BadgeID;references:ID" json:"badge"`
}

func (*TStudentCourse) TableName() string {
	return "t_student_course"
}

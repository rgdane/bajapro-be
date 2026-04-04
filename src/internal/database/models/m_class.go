package models

import "time"

type MClass struct {
	ID         int64      `gorm:"primaryKey;autoIncrement:true;type:serial" json:"id"`
	TeacherID  int64      `gorm:"column:teacher_id" json:"teacher_id"`
	ClassName  string     `gorm:"column:class_name;size:100" json:"class_name"`
	SchoolName string     `gorm:"column:school_name;size:100" json:"school_name"`
	ClassCode  string     `gorm:"column:class_code;size:50" json:"class_code"`
	IsActive   bool       `gorm:"column:isactive;default:true" json:"isactive"`
	CreatedAt  time.Time  `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt  time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
	UserCreate int64      `gorm:"column:user_create" json:"user_create"`
	UserUpdate int64      `gorm:"column:user_update" json:"user_update"`

	// Foreign Key Relationships
	Teacher *MUsers `gorm:"foreignKey:TeacherID;references:ID" json:"teacher"`
}

func (*MClass) TableName() string {
	return "m_class"
}

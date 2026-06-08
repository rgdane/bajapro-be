package models

import (
	"time"

	"gorm.io/gorm"
)

type MClass struct {
    ID         int64      `gorm:"primaryKey;autoIncrement:true" json:"id"`
    ClassName  string     `gorm:"column:class_name;size:100" json:"class_name"`
    SchoolName string     `gorm:"column:school_name;size:100" json:"school_name"`
    ClassCode  string     `gorm:"column:class_code;size:50;uniqueIndex" json:"class_code"`
    IsActive   bool       `gorm:"column:isactive;default:true" json:"isactive"`
    CreatedAt  time.Time  `gorm:"autoCreateTime" json:"created_at"`
    UpdatedAt  *time.Time `gorm:"autoUpdateTime" json:"updated_at"`
    DeletedAt  gorm.DeletedAt `gorm:"index:idx_class_deleted_at" json:"deleted_at"`

    // Relationships
    // Menampilkan daftar guru di kelas ini
    Teachers []User `gorm:"many2many:m_class_teachers;" json:"teachers"`
    // Menampilkan daftar siswa di kelas ini
    Students []User `gorm:"foreignKey:ClassID" json:"students"`
}

func (*MClass) TableName() string {
	return "m_class"
}

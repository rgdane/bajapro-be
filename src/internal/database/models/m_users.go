package models

import "time"

type MUser struct {
    ID                int64      `gorm:"primaryKey;autoIncrement:true" json:"id"`
    RoleID            int64      `gorm:"column:role_id" json:"role_id"`
    // ClassID digunakan untuk Siswa. Gunakan pointer *int64 agar bisa NULL untuk Admin/Guru
    ClassID           *int64     `gorm:"column:class_id" json:"class_id"`
    Name              string     `gorm:"size:100" json:"name"`
    Email             string     `gorm:"size:100;uniqueIndex" json:"email"`
    Password          string     `gorm:"type:text" json:"-"` // Sebaiknya password tidak muncul di JSON
    IsApprovedByAdmin bool       `gorm:"column:is_approved_by_admin;default:false" json:"is_approved_by_admin"`
    IsActive          bool       `gorm:"column:isactive;default:true" json:"isactive"`
    CreatedAt         time.Time  `gorm:"autoCreateTime" json:"created_at"`
    UpdatedAt         *time.Time `gorm:"autoUpdateTime" json:"updated_at"`

    // Relationships
    Role    *MRole   `gorm:"foreignKey:RoleID" json:"role"`
    Class   *MClass  `gorm:"foreignKey:ClassID" json:"class"` // Relasi milik Siswa
    // Relasi Many-to-Many untuk Guru
    TeachingClasses []MClass `gorm:"many2many:m_class_teachers;" json:"teaching_classes"`
}

func (*MUser) TableName() string {
	return "m_users"
}

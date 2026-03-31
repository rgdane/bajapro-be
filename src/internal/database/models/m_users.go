package models

import "time"

type MUsers struct {
	ID                int64      `gorm:"primaryKey;autoIncrement:true;type:serial" json:"id"`
	RoleID            int64      `gorm:"column:role_id" json:"role_id"`
	ClassID           int64      `gorm:"column:class_id" json:"class_id"`
	Name              string     `gorm:"size:100" json:"name"`
	Email             string     `gorm:"size:100;uniqueIndex" json:"email"`
	Password          string     `gorm:"type:text" json:"password"`
	IsApprovedByAdmin bool       `gorm:"column:is_approved_by_admin;default:false" json:"is_approved_by_admin"`
	InstansiSekolah   string     `gorm:"column:instansi_sekolah;size:100" json:"instansi_sekolah"`
	IsActive          bool       `gorm:"column:isactive;default:true" json:"isactive"`
	CreatedAt         time.Time  `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt         *time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
	UserCreate        int64      `gorm:"column:user_create" json:"user_create"`
	UserUpdate        int64      `gorm:"column:user_update" json:"user_update"`

	// Foreign Key Relationships
	Role  *MRole  `gorm:"foreignKey:RoleID;references:ID" json:"role"`
	Class *MClass `gorm:"foreignKey:ClassID;references:ID" json:"class"`
}

func (*MUsers) TableName() string {
	return "m_users"
}

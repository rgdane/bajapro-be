package dto

import (
	"jk-api/internal/database/models"
	"time"
)

// CreateMUserDto is used when creating a new MUser.
type CreateMUserDto struct {
	RoleID            int64  `json:"role_id"`
	ClassID           int64  `json:"class_id"`
	Name              string `json:"name"`
	Email             string `json:"email"`
	Password          string `json:"password"`
	IsApprovedByAdmin bool   `json:"is_approved_by_admin"`
	InstansiSekolah   string `json:"instansi_sekolah"`
	IsActive          bool   `json:"is_active"`
}

// UpdateMUserDto is used when updating an existing MUser.
type UpdateMUserDto struct {
	RoleID            *int64     `json:"role_id"`
	ClassID           *int64     `json:"class_id"`
	Name              *string    `json:"name"`
	Email             *string    `json:"email"`
	Password          *string    `json:"password"`
	OldPassword       *string    `json:"old_password"`
	NewPassword       *string    `json:"new_password"`
	IsApprovedByAdmin *bool      `json:"is_approved_by_admin"`
	InstansiSekolah   *string    `json:"instansi_sekolah"`
	IsActive          *bool      `json:"is_active"`
	DeletedAt         *time.Time `json:"deleted_at"`
}

// MUserResponseDto represents MUser output.
type MUserResponseDto struct {
	models.MUsers
	Role  *MRoleResponseDto  `json:"role,omitempty"`
	Class *MClassResponseDto `json:"class,omitempty"`
}

// BulkCreateMUserDto for batch creation.
type BulkCreateMUserDto struct {
	Data []CreateMUserDto `json:"m_users" binding:"required"`
}

// MUserFilterDto for query parameters.
type MUserFilterDto struct {
	RoleID            *int64  `form:"role_id"`
	ClassID           *int64  `form:"class_id"`
	Name              *string `form:"name"`
	Email             *string `form:"email"`
	IsApprovedByAdmin *bool   `form:"is_approved_by_admin"`
	InstansiSekolah   *string `form:"instansi_sekolah"`
	IsActive          *bool   `form:"is_active"`
	Preload           bool    `form:"preload"`
	Limit             int64   `form:"limit"`
	Cursor            int64   `form:"cursor"`
	Sort              string  `form:"sort"`
	Order             string  `form:"order"`
	ShowDeleted       bool    `form:"show_deleted"`
	Restore           bool    `form:"restore"`
}

// BulkUpdateMUserDto for batch update.
type BulkUpdateMUserDto struct {
	IDs      []int64         `json:"ids" binding:"required"`
	Data     *UpdateMUserDto `json:"data"`
	HasRoles *[]int64        `json:"has_roles"`
}

// BulkDeleteMUserDto for batch delete.
type BulkDeleteMUserDto struct {
	IDs []int64 `json:"ids" binding:"required"`
}

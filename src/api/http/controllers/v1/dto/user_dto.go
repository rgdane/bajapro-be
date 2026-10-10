package dto

import (
	"jk-api/internal/database/models"
	"time"
)

// CreateUserDto is used when creating a new User.
type CreateUserDto struct {
	Code              *string `json:"code"`
	Name              string  `json:"name"`
	Email             string  `json:"email"`
	Password          string  `json:"password"`
	IsPasswordDefault *bool   `json:"is_password_default"`
	Nip               *string `json:"nip"`
	Role              string  `json:"role"`
	SchoolID          *int64  `json:"school_id"`
	ClassID           *int64  `json:"class_id"`
}

// UpdateUserDto is used when updating an existing User.
type UpdateUserDto struct {
	Code        *string    `json:"code"`
	Name        *string    `json:"name"`
	Email       *string    `json:"email"`
	Password    *string    `json:"password"`
	Nip         *string    `json:"nip"`
	Role        *string    `json:"role"`
	OldPawword  *string    `json:"old_password"`
	NewPassword *string    `json:"new_password"`
	SchoolID    *int64     `json:"school_id"`
	ClassID     *int64     `json:"class_id"`
	Approval    *int8      `json:"approval" description:"0=Pending, 1=Disetujui, 2=Ditolak"`
	DeletedAt   *time.Time `json:"deleted_at"`
}

// UserResponseDto represents a detailed view of User with related data.
type UserResponseDto struct {
	models.User
}

type BulkCreateUserDto struct {
	Data []CreateUserDto `json:"users" binding:"required"`
}

type UserFilterDto struct {
	Role        string
	SchoolID    *int64
	ClassID     *int64
	Preload     bool
	Limit       int64
	Cursor      int64
	Name        string
	Sort        string
	Order       string
	ShowDeleted bool
	Restore     bool
}

type BulkUpdateUserDto struct {
	IDs  []int64        `json:"ids" binding:"required"`
	Data *UpdateUserDto `json:"data"`
}

type BulkDeleteUserDto struct {
	IDs []int64 `json:"ids" binding:"required"`
}

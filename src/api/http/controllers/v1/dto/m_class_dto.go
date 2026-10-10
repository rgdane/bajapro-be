package dto

import (
	"jk-api/internal/database/models"
	"time"
)

type CreateMClassDto struct {
	ClassName string `json:"class_name"`
	SchoolID  int64  `json:"school_id"`
	TeacherID *int64 `json:"teacher_id"`
	ClassCode string `json:"class_code"`
}

type UpdateMClassDto struct {
	ClassName *string    `json:"class_name"`
	SchoolID  *int64     `json:"school_id"`
	TeacherID *int64     `json:"teacher_id"`
	ClassCode *string    `json:"class_code"`
	IsActive  *bool      `json:"isactive"`
	DeletedAt *time.Time `json:"deleted_at"`
	Students  *[]int64   `json:"students"`
}

type BulkCreateMClassDto struct {
	Data []*CreateMClassDto `json:"data" binding:"required"`
}

type BulkUpdateMClassDto struct {
	IDs  []int64          `json:"ids" binding:"required"`
	Data *UpdateMClassDto `json:"data" binding:"required"`
}

type BulkDeleteMClassDto struct {
	IDs []int64 `json:"ids" binding:"required"`
}

// MClassResponseDto represents a detailed view of MClass with related data.
type MClassResponseDto struct {
	models.MClass
}

type MClassFilterDto struct {
	Preload     bool
	SchoolID    *int64
	TeacherID   *int64
	Sort        string
	Order       string
	Limit       int64
	Cursor      int64
	Name        string
	ShowDeleted bool
	Restore     bool
}

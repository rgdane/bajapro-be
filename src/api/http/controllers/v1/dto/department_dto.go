package dto

import (
	"jk-api/internal/database/models"
	"time"
)

// CreateDepartmentDto is used when creating a new Department.
type CreateDepartmentDto struct {
	Name string  `json:"name"`
	Code *string `json:"code"`
}

// UpdateDepartmentDto is used when updating an existing Department.
type UpdateDepartmentDto struct {
	Name      *string    `json:"name"`
	Code      *string    `json:"code"`
	DeletedAt *time.Time `json:"deleted_at"`
}

type BulkCreateDepartments struct {
	Data []*CreateDepartmentDto `json:"data" binding:"required"`
}

type BulkUpdateDepartmentDto struct {
	IDs  []int64              `json:"ids" binding:"required"`
	Data *UpdateDepartmentDto `json:"data" binding:"required"`
}

type BulkDeleteDepartmentDto struct {
	IDs []int64 `json:"ids" binding:"required"`
}

type DepartmentFilterDto struct {
	Preload     bool
	Sort        string
	Order       string
	Limit       int64
	Cursor      int64
	Name        string
	ShowDeleted bool
	Restore     bool
}

// DepartmentResponseDto represents a detailed view of Department with related data.
type DepartmentResponseDto struct {
	models.Department
}

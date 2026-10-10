package dto

import (
	"jk-api/internal/database/models"
	"time"
)

type CreateMSchoolDto struct {
	SchoolName string `json:"school_name"`
	IsActive   *bool  `json:"isactive"`
}

type UpdateMSchoolDto struct {
	SchoolName *string    `json:"school_name"`
	IsActive   *bool      `json:"isactive"`
	DeletedAt  *time.Time `json:"deleted_at"`
}

type MSchoolFilterDto struct {
	Preload     bool
	Name        string
	ShowDeleted bool
}

type MSchoolResponseDto struct {
	models.MSchool
}

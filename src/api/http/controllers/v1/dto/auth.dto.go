package dto

import "jk-api/internal/database/models"

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	User         UserResponse `json:"user"`
	AccessToken  string       `json:"access_token"`
	RefreshToken string       `json:"refresh_token"`
}

type UserResponse struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Role     string `json:"role"`
	SchoolID *int64 `json:"school_id,omitempty"`
	ClassID  *int64 `json:"class_id,omitempty"`
}

type RegisterRequest struct {
	Name               string  `json:"name"`
	Email              string  `json:"email"`
	Password           string  `json:"password"`
	Role               string  `json:"role"` // "admin" | "teacher" | "student"
	NIP                *string `json:"nip,omitempty"`
	ClassCode          *string `json:"class_code,omitempty"`           // For student registration
	SchoolID           *int64  `json:"school_id,omitempty"`            // For teacher registration with existing school
	ProposedSchoolName *string `json:"proposed_school_name,omitempty"` // For teacher registration proposing new school
}

type ProfileResponse struct {
	ID                int64           `json:"id"`
	Name              string          `json:"name"`
	Email             string          `json:"email"`
	IsPasswordDefault bool            `json:"is_password_default"`
	Role              string          `json:"role"`
	SchoolID          *int64          `json:"school_id,omitempty"`
	ClassID           *int64          `json:"class_id,omitempty"`
	School            *models.MSchool `json:"school,omitempty"`
	HasClass          *models.MClass  `json:"class,omitempty"`
}

type TeacherApprovalDecisionRequest struct {
	Action        string  `json:"action"` // "approve" | "reject"
	RejectionNote *string `json:"rejection_note,omitempty"`
}

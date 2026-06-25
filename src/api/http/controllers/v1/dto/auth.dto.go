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
	ID       int64         `json:"id"`
	Name     string        `json:"name"`
	Email    string        `json:"email"`
	HasRoles []models.Role `json:"has_roles"`
}

type RegisterRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"` // "student" | "teacher"
}

type ProfileResponse struct {
	ID        int64          `json:"id"`
	Name      string         `json:"name"`
	Email     string         `json:"email"`
	IsPasswordDefault bool `json:"is_password_default"`
	HasRoles  []models.Role  `json:"has_roles"`
	HasClass  *models.MClass `json:"class,omitempty"`
}

package dto

import "jk-api/internal/database/models"

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	ID    int64          `json:"id"`
	Name  string         `json:"name"`
	Email string         `json:"email"`
	Token string         `json:"token"`
	Role  *models.MRole  `json:"role"`
	Class *models.MClass `json:"class"`
}

type ProfileResponse struct {
	ID       int64          `json:"id"`
	Name     string         `json:"name"`
	Email    string         `json:"email"`
	IsActive bool           `json:"is_active"`
	Role     *models.MRole  `json:"role"`
	Class    *models.MClass `json:"class"`
}

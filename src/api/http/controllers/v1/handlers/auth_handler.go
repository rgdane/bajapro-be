package handlers

import (
	"fmt"
	"jk-api/api/http/controllers/v1/dto"
	"jk-api/api/http/controllers/v1/mapper"
	"jk-api/pkg/services/v1"
)

type AuthHandler struct {
	AuthService services.AuthService
}

func NewAuthHandler(service services.AuthService) *AuthHandler {
	return &AuthHandler{AuthService: service}
}

func (h *AuthHandler) GetProfileHandler(token string) (*dto.ProfileResponse, error) {
	user, err := h.AuthService.GetProfile(token)
	fmt.Println(user)
	if err != nil {
		return nil, err
	}
	data, _ := mapper.AuthModelToProfile(user)
	return data, nil
}

func (h *AuthHandler) Login(req *dto.LoginRequest) (*dto.LoginResponse, string, error) {
	user, err := h.AuthService.Login(req.Email, req.Password)
	if err != nil {
		return nil, "", err
	}

	accessToken, err := h.AuthService.GenerateAccessToken(user)
	if err != nil {
		return nil, "", err
	}

	refreshToken, err := h.AuthService.GenerateRefreshToken(user)
	if err != nil {
		return nil, "", err
	}

	// 🔥 kirim ke response
	data, err := mapper.AuthModelToDto(user, accessToken, refreshToken)
	if err != nil {
		return nil, "", err
	}

	// 👉 tambahin refresh token ke DTO
	data.RefreshToken = refreshToken

	return data, accessToken, nil
}

func (h *AuthHandler) Logout(token string) error {
	return h.AuthService.Logout(token)
}

func (h *AuthHandler) RefreshToken(Token string) (string, error) {
	return h.AuthService.RefreshToken(Token)
}

func (h *AuthHandler) Register(req *dto.RegisterRequest) (*dto.LoginResponse, string, error) {
	user, err := h.AuthService.Register(req)
	if err != nil {
		return nil, "", err
	}

	// ❗ kalau teacher belum approve → jangan kasih token
	if user.RoleID == 2 && !user.IsApprovedByAdmin {
		return nil, "", fmt.Errorf("akun teacher menunggu approval admin")
	}

	accessToken, err := h.AuthService.GenerateAccessToken(user)
	if err != nil {
		return nil, "", err
	}

	refreshToken, err := h.AuthService.GenerateRefreshToken(user)
	if err != nil {
		return nil, "", err
	}

	data, err := mapper.AuthModelToDto(user, accessToken, refreshToken)
	if err != nil {
		return nil, "", err
	}

	return data, accessToken, nil
}

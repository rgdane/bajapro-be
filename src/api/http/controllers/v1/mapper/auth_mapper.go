package mapper

import (
	"jk-api/api/http/controllers/v1/dto"
	"jk-api/internal/database/models"
)

func AuthModelToDto(data *models.MUsers, token string) (*dto.LoginResponse, error) {
	if data == nil {
		return nil, nil
	}

	response := &dto.LoginResponse{
		ID:    data.ID,
		Name:  data.Name,
		Email: data.Email,
		Token: token,
		Role:  data.Role,
		Class: data.Class,
	}
	return response, nil
}

func AuthModelToProfile(data *models.MUsers) (*dto.ProfileResponse, error) {
	if data == nil {
		return nil, nil
	}

	response := &dto.ProfileResponse{
		ID:       data.ID,
		Name:     data.Name,
		Email:    data.Email,
		Role:     data.Role,
		Class:    data.Class,
		IsActive: data.IsActive,
	}
	return response, nil
}

package mapper

import (
	"jk-api/api/http/controllers/v1/dto"
	"jk-api/internal/database/models"
)

func AuthModelToDto(data *models.User, accessToken, refreshToken string) (*dto.LoginResponse, error) {
	return &dto.LoginResponse{
		User: dto.UserResponse{
			ID:       data.ID,
			Name:     data.Name,
			Email:    data.Email,
			Role:     string(data.Role),
			SchoolID: data.SchoolID,
			ClassID:  data.ClassID,
		},
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func AuthModelToProfile(data *models.User) (*dto.ProfileResponse, error) {
	if data == nil {
		return nil, nil
	}

	response := &dto.ProfileResponse{
		ID:                data.ID,
		Name:              data.Name,
		Email:             data.Email,
		Role:              string(data.Role),
		SchoolID:          data.SchoolID,
		ClassID:           data.ClassID,
		School:            data.School,
		HasClass:          data.HasClass,
		IsPasswordDefault: data.IsPasswordDefault,
	}
	return response, nil
}

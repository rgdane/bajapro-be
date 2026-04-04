package mapper

import (
	"jk-api/api/http/controllers/v1/dto"
	"jk-api/internal/database/models"
)

func MRoleModelToResponseDto(data *models.MRole) (*dto.MRoleResponseDto, error) {
	if data == nil {
		return nil, nil
	}

	return &dto.MRoleResponseDto{
		ID:       data.ID,
		RoleName: data.RoleName,
		IsActive: data.IsActive,
	}, nil
}

func MClassModelToResponseDto(data *models.MClass) (*dto.MClassResponseDto, error) {
	if data == nil {
		return nil, nil
	}

	return &dto.MClassResponseDto{
		ID:         data.ID,
		TeacherID:  data.TeacherID,
		ClassName:  data.ClassName,
		SchoolName: data.SchoolName,
		ClassCode:  data.ClassCode,
		IsActive:   data.IsActive,
	}, nil
}

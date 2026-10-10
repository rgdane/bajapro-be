package mapper

import (
	"jk-api/api/http/controllers/v1/dto"
	"jk-api/internal/database/models"
)

func CreateMSchoolDtoToModel(dto *dto.CreateMSchoolDto) (*models.MSchool, error) {
	if dto == nil {
		return nil, nil
	}

	isActive := true
	if dto.IsActive != nil {
		isActive = *dto.IsActive
	}

	data := &models.MSchool{
		SchoolName: dto.SchoolName,
		IsActive:   isActive,
	}

	return data, nil
}

func UpdateMSchoolDtoToModel(dto *dto.UpdateMSchoolDto) map[string]interface{} {
	payload := make(map[string]interface{})

	if dto.SchoolName != nil {
		payload["school_name"] = *dto.SchoolName
	}
	if dto.IsActive != nil {
		payload["isactive"] = *dto.IsActive
	}
	if dto.DeletedAt != nil {
		payload["deleted_at"] = dto.DeletedAt
	}

	return payload
}

func MSchoolModelToResponseDto(data *models.MSchool) (*dto.MSchoolResponseDto, error) {
	if data == nil {
		return nil, nil
	}

	return &dto.MSchoolResponseDto{
		MSchool: *data,
	}, nil
}

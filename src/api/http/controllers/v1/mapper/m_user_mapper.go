package mapper

import (
	"jk-api/api/http/controllers/v1/dto"
	"jk-api/internal/database/models"
)

func CreateMUserDtoToModel(input *dto.CreateMUserDto) (*models.MUsers, error) {
	return &models.MUsers{
		RoleID:            input.RoleID,
		ClassID:           input.ClassID,
		Name:              input.Name,
		Email:             input.Email,
		Password:          input.Password,
		IsApprovedByAdmin: input.IsApprovedByAdmin,
		InstansiSekolah:   input.InstansiSekolah,
		IsActive:          input.IsActive,
	}, nil
}

func UpdateMUserDtoToMap(dto *dto.UpdateMUserDto) (
	payload map[string]interface{},
	associations map[string]interface{},
) {
	payload = make(map[string]interface{})
	associations = make(map[string]interface{})

	if dto.RoleID != nil {
		payload["role_id"] = *dto.RoleID
	}
	if dto.ClassID != nil {
		payload["class_id"] = *dto.ClassID
	}
	if dto.Name != nil {
		payload["name"] = *dto.Name
	}
	if dto.Email != nil {
		payload["email"] = *dto.Email
	}
	if dto.Password != nil {
		payload["password"] = *dto.Password
	}
	if dto.IsApprovedByAdmin != nil {
		payload["is_approved_by_admin"] = *dto.IsApprovedByAdmin
	}
	if dto.InstansiSekolah != nil {
		payload["instansi_sekolah"] = *dto.InstansiSekolah
	}
	if dto.IsActive != nil {
		payload["is_active"] = *dto.IsActive
	}

	return payload, associations
}

func MUserModelToResponseDto(data *models.MUsers) (*dto.MUserResponseDto, error) {
	if data == nil {
		return nil, nil
	}

	response := &dto.MUserResponseDto{
		MUsers: *data,
	}

	if data.Role != nil {
		roleDto, err := MRoleModelToResponseDto(data.Role)
		if err != nil {
			return nil, err
		}
		response.Role = roleDto
	}

	if data.Class != nil {
		classDto, err := MClassModelToResponseDto(data.Class)
		if err != nil {
			return nil, err
		}
		response.Class = classDto
	}

	return response, nil
}

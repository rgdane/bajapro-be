package mapper

import (
	"jk-api/api/http/controllers/v1/dto"
	"jk-api/internal/database/models"
)

func CreateMClassDtoToModel(dto *dto.CreateMClassDto) (*models.MClass, error) {
	if dto == nil {
		return nil, nil
	}

	data := &models.MClass{
		ClassName: dto.ClassName,
		SchoolID:  dto.SchoolID,
		TeacherID: dto.TeacherID,
		ClassCode: dto.ClassCode,
	}

	return data, nil
}

func UpdateMClassDtoToModel(dto *dto.UpdateMClassDto) (
	payload map[string]interface{},
	associations map[string]interface{},
	err error,
) {
	payload = make(map[string]interface{})
	associations = make(map[string]interface{})

	if dto.ClassName != nil {
		payload["class_name"] = *dto.ClassName
	}
	if dto.SchoolID != nil {
		payload["school_id"] = *dto.SchoolID
	}
	if dto.TeacherID != nil {
		payload["teacher_id"] = *dto.TeacherID
	}
	if dto.ClassCode != nil {
		payload["class_code"] = *dto.ClassCode
	}
	if dto.IsActive != nil {
		payload["isactive"] = *dto.IsActive
	}

	if dto.Students != nil {
		var users []models.User
		for _, id := range *dto.Students {
			users = append(users, models.User{ID: id})
		}
		associations["Students"] = users
	}

	return payload, associations, nil
}

func MClassModelToResponseDto(data *models.MClass) (*dto.MClassResponseDto, error) {
	if data == nil {
		return nil, nil
	}

	responseDto := &dto.MClassResponseDto{
		MClass: *data,
	}

	return responseDto, nil
}

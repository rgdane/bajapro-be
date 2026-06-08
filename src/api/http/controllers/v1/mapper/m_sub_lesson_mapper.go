package mapper

import (
	"jk-api/api/http/controllers/v1/dto"
	"jk-api/internal/database/models"
)

func CreateMSubLessonDtoToModel(dto *dto.CreateMSubLessonDto) (*models.MSubLesson, error) {
	if dto == nil {
		return nil, nil
	}

	data := &models.MSubLesson{
		LessonID:      dto.LessonID,
		Title:           dto.Title,
		OrderPosition: dto.OrderPosition,
	}

	return data, nil
}

func UpdateMSubLessonDtoToModel(dto *dto.UpdateMSubLessonDto) (map[string]interface{}, error) {
	if dto == nil {
		return nil, nil
	}

	updates := map[string]interface{}{}

	if dto.LessonID != nil {
		updates["lesson_id"] = *dto.LessonID
	}
	if dto.Title != nil {
		updates["title"] = *dto.Title
	}
	if dto.OrderPosition != nil {
		updates["order_position"] = *dto.OrderPosition
	}
	if dto.IsActive != nil {
		updates["is_active"] = *dto.IsActive
	}


	return updates, nil
}

func MSubLessonModelToResponseDto(data *models.MSubLesson) (*dto.MSubLessonResponseDto, error) {
	if data == nil {
		return nil, nil
	}

	responseDto := &dto.MSubLessonResponseDto{
		MSubLesson: *data,
	}

	return responseDto, nil
}

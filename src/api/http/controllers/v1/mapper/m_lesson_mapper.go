package mapper

import (
	"jk-api/api/http/controllers/v1/dto"
	"jk-api/internal/database/models"
)

func CreateMLessonDtoToModel(dto *dto.CreateMLessonDto) (*models.MLesson, error) {
	if dto == nil {
		return nil, nil
	}

	data := &models.MLesson{
		CourseID:    dto.CourseID,
		LevelID:     dto.LevelID,
		Title:       dto.Title,
		Description: dto.Description,
		Position:    dto.Position,
	}

	return data, nil
}

func UpdateMLessonDtoToModel(dto *dto.UpdateMLessonDto) (map[string]interface{}, error) {
	if dto == nil {
		return nil, nil
	}

	updates := map[string]interface{}{}

	if dto.CourseID != nil {
		updates["course_id"] = *dto.CourseID
	}
	if dto.LevelID != nil {
		updates["level_id"] = *dto.LevelID
	}
	if dto.Title != nil {
		updates["title"] = *dto.Title
	}
	if dto.Description != nil {
		updates["description"] = *dto.Description
	}
	if dto.Position != nil {
		updates["position"] = *dto.Position
	}
	if dto.ImgThumbnail != nil {
		updates["img_thumbnail"] = *dto.ImgThumbnail
	}
	if dto.Published != nil {
		updates["published"] = *dto.Published
	}
	if dto.IsActive != nil {
		updates["is_active"] = *dto.IsActive
	}


	return updates, nil
}

func MLessonModelToResponseDto(data *models.MLesson) (*dto.MLessonResponseDto, error) {
	if data == nil {
		return nil, nil
	}

	responseDto := &dto.MLessonResponseDto{
		MLesson: *data,
	}

	return responseDto, nil
}

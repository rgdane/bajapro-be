package mapper

import (
	"jk-api/api/http/controllers/v1/dto"
	"jk-api/internal/database/models"
)

func CreateMMaterialDtoToModel(dto *dto.CreateMMaterialDto) (*models.MMaterials, error) {
	if dto == nil {
		return nil, nil
	}

	data := &models.MMaterials{
		SubLessonID: dto.SubLessonID,
		Title:       dto.Title,
		Materials:  dto.Materials,
		URLVideo:  dto.URLVideo,
		ContentPosition: dto.ContentPosition,
		PromptLLM:       dto.PromptLLM,
	}

	return data, nil
}

func UpdateMMaterialDtoToModel(dto *dto.UpdateMMaterialDto) (map[string]interface{}, error) {
	if dto == nil {
		return nil, nil
	}

	updates := map[string]interface{}{}

	if dto.SubLessonID != nil {
		updates["sub_lesson_id"] = *dto.SubLessonID
	}
	if dto.Title != nil {
		updates["title"] = *dto.Title
	}
	if dto.Materials != nil {
		updates["materials"] = *dto.Materials
	}
	if dto.URLVideo != nil {
		updates["url_video"] = *dto.URLVideo
	}
	if dto.PromptLLM != nil {
		updates["prompt_llm"] = *dto.PromptLLM
	}
	if dto.ContentPosition != nil {
		updates["content_position"] = *dto.ContentPosition
	}
	if dto.Published != nil {
		updates["published"] = *dto.Published
	}
	if dto.IsActive != nil {
		updates["is_active"] = *dto.IsActive
	}


	return updates, nil
}

func MMaterialModelToResponseDto(data *models.MMaterials) (*dto.MMaterialResponseDto, error) {
	if data == nil {
		return nil, nil
	}

	responseDto := &dto.MMaterialResponseDto{
		MMaterials: *data,
	}

	return responseDto, nil
}

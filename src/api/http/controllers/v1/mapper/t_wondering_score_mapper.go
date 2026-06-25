package mapper

import (
	"jk-api/api/http/controllers/v1/dto"
	"jk-api/internal/database/models"
)

func CreateTWonderingScoreDtoToModel(
	dto *dto.TWonderingScoreCreateDto,
	userID int64,
) (*models.TWonderingScore, error) {

	data := &models.TWonderingScore{
		UserID:      userID,
		SubLessonID: dto.SubLessonID,

		// Wondering point otomatis
		Score:    10,
	}

	return data, nil
}

func TWonderingScoreModelToResponseDto(
	data *models.TWonderingScore,
) (*dto.TWonderingScoreResponseDto, error) {

	if data == nil {
		return nil, nil
	}

	response := &dto.TWonderingScoreResponseDto{
		TWonderingScore: *data,
	}

	return response, nil
}
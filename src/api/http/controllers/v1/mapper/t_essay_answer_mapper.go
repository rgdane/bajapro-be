package mapper

import (
	"jk-api/api/http/controllers/v1/dto"
	"jk-api/internal/database/models"
)

func CreateTEssayAnswerDtoToModel(
	dto *dto.TEssayAnswerCreateDto,
	userID int64,
) (*models.TEssayAnswer, error) {

	data := &models.TEssayAnswer{
		UserID:      userID,
		EssayQuestionID: dto.EssayQuestionID,
		Answer:      dto.Answer,
		KonteksPenjelasan: dto.KonteksPenjelasan,
		Keruntutan:   dto.Keruntutan,
		Kebenaran:    dto.Kebenaran,
		TeacherNotes: dto.TeacherNotes,
		IsApprovedByTeacher: dto.IsApprovedByTeacher,	
	}

	return data, nil
}

func TEssayAnswerModelToResponseDto(
	data *models.TEssayAnswer,
) (*dto.TEssayAnswerResponseDto, error) {

	if data == nil {
		return nil, nil
	}

	response := &dto.TEssayAnswerResponseDto{
		TEssayAnswer: *data,
	}

	return response, nil
}
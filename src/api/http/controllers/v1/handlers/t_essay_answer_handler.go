package handlers

import (
	"jk-api/api/http/controllers/v1/dto"
	"jk-api/api/http/controllers/v1/mapper"
	"jk-api/pkg/services/v1"
)

type TEssayAnswerHandler struct {
	Service services.TEssayAnswerService
}

func NewTEssayAnswerHandler(service services.TEssayAnswerService) *TEssayAnswerHandler {
	return &TEssayAnswerHandler{Service: service}
}

func (h *TEssayAnswerHandler) CreateTEssayAnswerHandler(
	input *dto.TEssayAnswerCreateDto,
	userID int64,
) (*dto.TEssayAnswerResponseDto, error) {

	db := h.Service.GetDB().Begin()
	committed := false

	defer func() {
		if r := recover(); r != nil {
			db.Rollback()
			panic(r)
		}

		if !committed {
			db.Rollback()
		}
	}()

	service := h.Service.WithTx(db)

	payload, err := mapper.CreateTEssayAnswerDtoToModel(
		input,
		userID,
	)
	if err != nil {
		return nil, err
	}

	data, err := service.CreateTEssayAnswer(
		payload,
		userID,
	)
	if err != nil {
		return nil, err
	}

	if err := db.Commit().Error; err != nil {
		return nil, err
	}

	committed = true

	return mapper.TEssayAnswerModelToResponseDto(data)
}

func (h *TEssayAnswerHandler) GetTEssayAnswersByEssayQuestionIDAndUserIDHandler(
	filter dto.TEssayAnswerFilterDto,
	essayQuestionID int64,
	userID int64,
) (*dto.TEssayAnswerResponseDto, error) {

	data, err := h.Service.GetTEssayAnswersByEssayQuestionIDAndUserID(
		essayQuestionID,
		userID,
	)
	if err != nil {
		return nil, err
	}

	return mapper.TEssayAnswerModelToResponseDto(data)
}
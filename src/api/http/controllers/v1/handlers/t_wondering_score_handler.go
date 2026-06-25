package handlers

import (
	"jk-api/api/http/controllers/v1/dto"
	"jk-api/api/http/controllers/v1/mapper"
	"jk-api/pkg/services/v1"
)

type TWonderingScoreHandler struct {
	Service services.TWonderingScoreService
}

func NewTWonderingScoreHandler(service services.TWonderingScoreService) *TWonderingScoreHandler {
	return &TWonderingScoreHandler{Service: service}
}

func (h *TWonderingScoreHandler) CreateTWonderingScoreHandler(
	input *dto.TWonderingScoreCreateDto,
	userID int64,
) (*dto.TWonderingScoreResponseDto, error) {

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

	payload, err := mapper.CreateTWonderingScoreDtoToModel(
		input,
		userID,
	)
	if err != nil {
		return nil, err
	}

	data, err := service.CreateTWonderingScore(
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

	return mapper.TWonderingScoreModelToResponseDto(data)
}

func (h *TWonderingScoreHandler) GetTWonderingScoresBySubLessonIDAndUserIDHandler(
	filter dto.TWonderingScoreFilterDto,
	subLessonID int64,
	userID int64,
) (*dto.TWonderingScoreResponseDto, error) {

	data, err := h.Service.GetTWonderingScoresBySubLessonIDAndUserID(
		subLessonID,
		userID,
	)
	if err != nil {
		return nil, err
	}

	return mapper.TWonderingScoreModelToResponseDto(data)
}
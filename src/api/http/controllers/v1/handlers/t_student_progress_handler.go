
package handlers

import (
	"jk-api/api/http/controllers/v1/dto"
	"jk-api/api/http/controllers/v1/mapper"
	"jk-api/pkg/services/v1"
)

type TStudentProgressHandler struct {
	Service services.TStudentProgressService
}

func NewTStudentProgressHandler(service services.TStudentProgressService) *TStudentProgressHandler {
	return &TStudentProgressHandler{Service: service}
}

func (h *TStudentProgressHandler) GetStudentProgressHandler(
	userID int64,
	courseID int64,
) ([]dto.TStudentProgressResponseDto, error) {

	data, err := h.Service.GetStudentProgress(
		userID,
		courseID,
	)
	if err != nil {
		return nil, err
	}

	response := make(
		[]dto.TStudentProgressResponseDto,
		0,
		len(data),
	)

	for _, item := range data {
		dtoItem, err := mapper.TStudentProgressModelToResponseDto(&item)
		if err != nil {
			return nil, err
		}

		response = append(response, *dtoItem)
	}

	return response, nil
}

func (h *TStudentProgressHandler) CompleteTStudentProgressHandler(input *dto.CompleteTStudentProgressDto, userID int64) (*dto.TStudentProgressResponseDto, error) {
	db := h.Service.GetDB().Begin()
	committed := false
	defer func() {
		if r := recover(); r != nil {
			db.Rollback()
			panic(r) // terusin panic biar gak ketelen
		}
		if !committed {
			db.Rollback()
		}
	}()

	TStudentProgressService := h.Service.WithTx(db)

	payload, err := mapper.CompleteTStudentProgressDtoToModel(input, userID)
	if err != nil {
		return nil, err
	}

	createdData, err := TStudentProgressService.CompleteTStudentProgress(payload)
	if err != nil {
		return nil, err
	}

	if err := db.Commit().Error; err != nil {
		return nil, err
	}
	committed = true

	return mapper.TStudentProgressModelToResponseDto(createdData)
}




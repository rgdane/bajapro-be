
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


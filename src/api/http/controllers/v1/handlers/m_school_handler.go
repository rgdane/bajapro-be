package handlers

import (
	"jk-api/api/http/controllers/v1/dto"
	"jk-api/api/http/controllers/v1/mapper"
	"jk-api/internal/database/models"
	"jk-api/pkg/services/v1"
)

type MSchoolHandler struct {
	Service services.MSchoolService
}

func NewMSchoolHandler(service services.MSchoolService) *MSchoolHandler {
	return &MSchoolHandler{Service: service}
}

func (h *MSchoolHandler) CreateMSchoolHandler(input *dto.CreateMSchoolDto) (*dto.MSchoolResponseDto, error) {
	payload, err := mapper.CreateMSchoolDtoToModel(input)
	if err != nil {
		return nil, err
	}

	createdData, err := h.Service.CreateMSchool(payload)
	if err != nil {
		return nil, err
	}

	return mapper.MSchoolModelToResponseDto(createdData)
}

func (h *MSchoolHandler) UpdateMSchoolHandler(id int64, input *dto.UpdateMSchoolDto) (*models.MSchool, error) {
	payload := mapper.UpdateMSchoolDtoToModel(input)
	return h.Service.UpdateMSchool(id, payload)
}

func (h *MSchoolHandler) DeleteMSchoolHandler(id int64) error {
	return h.Service.DeleteMSchool(id)
}

func (h *MSchoolHandler) GetMSchoolByIDHandler(id int64, filter dto.MSchoolFilterDto) (*models.MSchool, error) {
	return h.Service.GetMSchoolByID(id, filter)
}

func (h *MSchoolHandler) GetAllMSchoolsHandler(filter dto.MSchoolFilterDto) ([]models.MSchool, error) {
	return h.Service.GetAllMSchools(filter)
}

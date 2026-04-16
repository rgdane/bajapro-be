package handlers

import (
	"fmt"
	"jk-api/api/http/controllers/v1/dto"
	"jk-api/api/http/controllers/v1/mapper"
	"jk-api/internal/database/models"
	"jk-api/pkg/services/v1"
)

type MMaterialHandler struct {
	Service services.MMaterialService
}

func NewMMaterialHandler(service services.MMaterialService) *MMaterialHandler {
	return &MMaterialHandler{Service: service}
}

func (h *MMaterialHandler) CreateMMaterialHandler(input *dto.CreateMMaterialDto) (*dto.MMaterialResponseDto, error) {
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

	mMaterialService := h.Service.WithTx(db)

	payload, err := mapper.CreateMMaterialDtoToModel(input)
	if err != nil {
		return nil, err
	}

	createdData, err := mMaterialService.CreateMMaterial(payload)
	if err != nil {
		return nil, err
	}

	if err := db.Commit().Error; err != nil {
		return nil, err
	}
	committed = true

	return mapper.MMaterialModelToResponseDto(createdData)
}

func (h *MMaterialHandler) UpdateMMaterialHandler(id int64, input *dto.UpdateMMaterialDto) (*models.MMaterials, error) {
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

	mMaterialService := h.Service.WithTx(db)

	payload, err := mapper.UpdateMMaterialDtoToModel(input)
	if err != nil {
		return nil, err
	}

	updatedData, err := mMaterialService.UpdateMMaterial(id, payload)
	if err != nil {
		return nil, err
	}

	if err := db.Commit().Error; err != nil {
		return nil, err
	}
	committed = true

	return updatedData, nil
}

func (h *MMaterialHandler) DeleteMMaterialHandler(id int64) error {
	return h.Service.DeleteMMaterial(id)
}

func (h *MMaterialHandler) GetMMaterialByIDHandler(id int64, filter dto.MMaterialFilterDto) (*models.MMaterials, error) {
	return h.Service.GetMMaterialByID(id, filter)
}

func (h *MMaterialHandler) GetAllMMaterialsHandler(filter dto.MMaterialFilterDto) ([]models.MMaterials, int64, error) {
	data, err  := h.Service.GetAllMMaterials(filter)
	if err != nil {
		return nil, 0, err
	}
	var total int64
	db := h.Service.GetDB()
	if filter.Name != "" {
		db = db.Where("name ILIKE ?", "%"+filter.Name+"%")
	}
	if filter.ShowDeleted {
		db = db.Unscoped().Where("deleted_at IS NOT NULL")
	}
	if err := db.Model(&models.MMaterials{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	return data, total, nil
}

func (h *MMaterialHandler) BulkCreateMMaterialsHandler(input *dto.BulkCreateMMaterialsDto) ([]*models.MMaterials, error) {
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

	mMaterialService := h.Service.WithTx(db)

	var mMaterials []*models.MMaterials
	for _, createDto := range input.Data {
		MMaterial, err := mapper.CreateMMaterialDtoToModel(createDto)
		if err != nil {
			return nil, err
		}
		if MMaterial != nil {
			mMaterials = append(mMaterials, MMaterial)
		}
	}

	createdMMaterials, err := mMaterialService.BulkCreateMMaterials(mMaterials)
	if err != nil {
		return nil, err
	}

	if err := db.Commit().Error; err != nil {
		return nil, err
	}
	committed = true

	return createdMMaterials, nil
}

func (h *MMaterialHandler) BulkUpdateMMaterialsHandler(input *dto.BulkUpdateMMaterialDto) ([]*models.MMaterials, error) {
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

	mMaterialService := h.Service.WithTx(db)

	updates, err := mapper.UpdateMMaterialDtoToModel(input.Data)
	if err != nil {
		return nil, fmt.Errorf("failed to map update data: %w", err)
	}

	if len(updates) == 0 {
		return nil, fmt.Errorf("update data cannot be empty")
	}

	err = mMaterialService.BulkUpdateMMaterials(input.IDs, updates)
	if err != nil {
		return nil, fmt.Errorf("failed to bulk update MMaterials: %w", err)
	}

	updatedMMaterials, err := mMaterialService.GetMMaterialsByIDs(input.IDs)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve updated MMaterials: %w", err)
	}

	if err := db.Commit().Error; err != nil {
		return nil, err
	}
	committed = true

	return updatedMMaterials, nil
}

func (h *MMaterialHandler) BulkDeleteMMaterialsHandler(input *dto.BulkDeleteMMaterialDto) error {
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

	mMaterialService := h.Service.WithTx(db)

	if err := mMaterialService.BulkDeleteMMaterials(input.IDs); err != nil {
		return err
	}

	if err := db.Commit().Error; err != nil {
		return err
	}
	committed = true

	return nil
}
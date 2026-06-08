package handlers

import (
	"fmt"
	"jk-api/api/http/controllers/v1/dto"
	"jk-api/api/http/controllers/v1/mapper"
	"jk-api/internal/database/models"
	"jk-api/pkg/services/v1"
)

type MLessonHandler struct {
	Service services.MLessonService
}

func NewMLessonHandler(service services.MLessonService) *MLessonHandler {
	return &MLessonHandler{Service: service}
}

func (h *MLessonHandler) CreateMLessonHandler(input *dto.CreateMLessonDto) (*dto.MLessonResponseDto, error) {
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

	mLessonService := h.Service.WithTx(db)

	payload, err := mapper.CreateMLessonDtoToModel(input)
	if err != nil {
		return nil, err
	}

	createdData, err := mLessonService.CreateMLesson(payload)
	if err != nil {
		return nil, err
	}

	if err := db.Commit().Error; err != nil {
		return nil, err
	}
	committed = true

	return mapper.MLessonModelToResponseDto(createdData)
}

func (h *MLessonHandler) UpdateMLessonHandler(id int64, input *dto.UpdateMLessonDto) (*models.MLesson, error) {
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

	mLessonService := h.Service.WithTx(db)

	payload, err := mapper.UpdateMLessonDtoToModel(input)
	if err != nil {
		return nil, err
	}

	updatedData, err := mLessonService.UpdateMLesson(id, payload)
	if err != nil {
		return nil, err
	}

	if err := db.Commit().Error; err != nil {
		return nil, err
	}
	committed = true

	return updatedData, nil
}

func (h *MLessonHandler) DeleteMLessonHandler(id int64) error {
	return h.Service.DeleteMLesson(id)
}

func (h *MLessonHandler) GetMLessonByIDHandler(id int64, filter dto.MLessonFilterDto) (*models.MLesson, error) {
	return h.Service.GetMLessonByID(id, filter)
}

func (h *MLessonHandler) GetAllMLessonsHandler(filter dto.MLessonFilterDto) ([]models.MLesson, int64, error) {
	data, err  := h.Service.GetAllMLessons(filter)
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
	if err := db.Model(&models.MLesson{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	return data, total, nil
}

func (h *MLessonHandler) BulkCreateMLessonsHandler(input *dto.BulkCreateMLessonsDto) ([]*models.MLesson, error) {
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

	mLessonService := h.Service.WithTx(db)

	var mLessons []*models.MLesson
	for _, createDto := range input.Data {
		MLesson, err := mapper.CreateMLessonDtoToModel(createDto)
		if err != nil {
			return nil, err
		}
		if MLesson != nil {
			mLessons = append(mLessons, MLesson)
		}
	}

	createdMLessons, err := mLessonService.BulkCreateMLessons(mLessons)
	if err != nil {
		return nil, err
	}

	if err := db.Commit().Error; err != nil {
		return nil, err
	}
	committed = true

	return createdMLessons, nil
}

func (h *MLessonHandler) BulkUpdateMLessonsHandler(input *dto.BulkUpdateMLessonDto) ([]*models.MLesson, error) {
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

	mLessonService := h.Service.WithTx(db)

	updates, err := mapper.UpdateMLessonDtoToModel(input.Data)
	if err != nil {
		return nil, fmt.Errorf("failed to map update data: %w", err)
	}

	if len(updates) == 0 {
		return nil, fmt.Errorf("update data cannot be empty")
	}

	err = mLessonService.BulkUpdateMLessons(input.IDs, updates)
	if err != nil {
		return nil, fmt.Errorf("failed to bulk update MLessons: %w", err)
	}

	updatedMLessons, err := mLessonService.GetMLessonsByIDs(input.IDs)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve updated MLessons: %w", err)
	}

	if err := db.Commit().Error; err != nil {
		return nil, err
	}
	committed = true

	return updatedMLessons, nil
}

func (h *MLessonHandler) BulkDeleteMLessonsHandler(input *dto.BulkDeleteMLessonDto) error {
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

	mLessonService := h.Service.WithTx(db)

	if err := mLessonService.BulkDeleteMLessons(input.IDs); err != nil {
		return err
	}

	if err := db.Commit().Error; err != nil {
		return err
	}
	committed = true

	return nil
}
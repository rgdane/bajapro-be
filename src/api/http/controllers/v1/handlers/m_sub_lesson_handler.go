package handlers

import (
	"fmt"
	"jk-api/api/http/controllers/v1/dto"
	"jk-api/api/http/controllers/v1/mapper"
	"jk-api/internal/database/models"
	"jk-api/pkg/services/v1"
)

type MSubLessonHandler struct {
	Service services.MSubLessonService
}

func NewMSubLessonHandler(service services.MSubLessonService) *MSubLessonHandler {
	return &MSubLessonHandler{Service: service}
}

func (h *MSubLessonHandler) CreateMSubLessonHandler(input *dto.CreateMSubLessonDto) (*dto.MSubLessonResponseDto, error) {
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

	mSubLessonService := h.Service.WithTx(db)

	payload, err := mapper.CreateMSubLessonDtoToModel(input)
	if err != nil {
		return nil, err
	}

	createdData, err := mSubLessonService.CreateMSubLesson(payload)
	if err != nil {
		return nil, err
	}

	if err := db.Commit().Error; err != nil {
		return nil, err
	}
	committed = true

	return mapper.MSubLessonModelToResponseDto(createdData)
}

func (h *MSubLessonHandler) UpdateMSubLessonHandler(id int64, input *dto.UpdateMSubLessonDto) (*models.MSubLesson, error) {
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

	mSubLessonService := h.Service.WithTx(db)

	payload, err := mapper.UpdateMSubLessonDtoToModel(input)
	if err != nil {
		return nil, err
	}

	updatedData, err := mSubLessonService.UpdateMSubLesson(id, payload)
	if err != nil {
		return nil, err
	}

	if err := db.Commit().Error; err != nil {
		return nil, err
	}
	committed = true

	return updatedData, nil
}

func (h *MSubLessonHandler) DeleteMSubLessonHandler(id int64) error {
	return h.Service.DeleteMSubLesson(id)
}

func (h *MSubLessonHandler) GetMSubLessonByIDHandler(id int64, filter dto.MSubLessonFilterDto) (*models.MSubLesson, error) {
	return h.Service.GetMSubLessonByID(id, filter)
}

func (h *MSubLessonHandler) GetAllMSubLessonsHandler(filter dto.MSubLessonFilterDto) ([]models.MSubLesson, int64, error) {
	data, err  := h.Service.GetAllMSubLessons(filter)
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
	if err := db.Model(&models.MSubLesson{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	return data, total, nil
}

func (h *MSubLessonHandler) BulkCreateMSubLessonsHandler(input *dto.BulkCreateMSubLessonsDto) ([]*models.MSubLesson, error) {
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

	mSubLessonService := h.Service.WithTx(db)

	var mSubLessons []*models.MSubLesson
	for _, createDto := range input.Data {
		MSubLesson, err := mapper.CreateMSubLessonDtoToModel(createDto)
		if err != nil {
			return nil, err
		}
		if MSubLesson != nil {
			mSubLessons = append(mSubLessons, MSubLesson)
		}
	}

	createdMSubLessons, err := mSubLessonService.BulkCreateMSubLessons(mSubLessons)
	if err != nil {
		return nil, err
	}

	if err := db.Commit().Error; err != nil {
		return nil, err
	}
	committed = true

	return createdMSubLessons, nil
}

func (h *MSubLessonHandler) BulkUpdateMSubLessonsHandler(input *dto.BulkUpdateMSubLessonDto) ([]*models.MSubLesson, error) {
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

	mSubLessonService := h.Service.WithTx(db)

	updates, err := mapper.UpdateMSubLessonDtoToModel(input.Data)
	if err != nil {
		return nil, fmt.Errorf("failed to map update data: %w", err)
	}

	if len(updates) == 0 {
		return nil, fmt.Errorf("update data cannot be empty")
	}

	err = mSubLessonService.BulkUpdateMSubLessons(input.IDs, updates)
	if err != nil {
		return nil, fmt.Errorf("failed to bulk update MSubLessons: %w", err)
	}

	updatedMSubLessons, err := mSubLessonService.GetMSubLessonsByIDs(input.IDs)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve updated MLessons: %w", err)
	}

	if err := db.Commit().Error; err != nil {
		return nil, err
	}
	committed = true

	return updatedMSubLessons, nil
}

func (h *MSubLessonHandler) BulkDeleteMSubLessonsHandler(input *dto.BulkDeleteMSubLessonDto) error {
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

	mSubLessonService := h.Service.WithTx(db)

	if err := mSubLessonService.BulkDeleteMSubLessons(input.IDs); err != nil {
		return err
	}

	if err := db.Commit().Error; err != nil {
		return err
	}
	committed = true

	return nil
}
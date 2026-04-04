package handlers

import (
	"fmt"
	"jk-api/api/http/controllers/v1/dto"
	"jk-api/api/http/controllers/v1/mapper"
	"jk-api/internal/database/models"
	"jk-api/pkg/services/v1"
)

type MUserHandler struct {
	Service services.MUserService
}

func NewMUserHandler(service services.MUserService) *MUserHandler {
	return &MUserHandler{Service: service}
}

func (h *MUserHandler) CreateMUserHandler(input *dto.CreateMUserDto) (*dto.MUserResponseDto, error) {
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

	mUserService := h.Service.WithTx(db)

	payload, err := mapper.CreateMUserDtoToModel(input)
	if err != nil {
		return nil, err
	}

	createdData, err := mUserService.CreateMUser(payload)
	if err != nil {
		return nil, err
	}

	if err := db.Commit().Error; err != nil {
		return nil, err
	}
	committed = true

	return mapper.MUserModelToResponseDto(createdData)
}

func (h *MUserHandler) GetAllMUsersHandler(filter dto.MUserFilterDto) ([]dto.MUserResponseDto, error) {
	mUsers, err := h.Service.GetAllMUsers(filter)
	if err != nil {
		return nil, err
	}

	var output []dto.MUserResponseDto
	for _, m := range mUsers {
		item, err := mapper.MUserModelToResponseDto(&m)
		if err != nil {
			return nil, err
		}
		output = append(output, *item)
	}

	return output, nil
}

func (h *MUserHandler) GetMUserByIDHandler(id int64, filter dto.MUserFilterDto) (*dto.MUserResponseDto, error) {
	mUser, err := h.Service.GetMUserByID(id, filter)
	if err != nil {
		return nil, err
	}

	return mapper.MUserModelToResponseDto(mUser)
}

func (h *MUserHandler) UpdateMUserHandler(id int64, input *dto.UpdateMUserDto) (*dto.MUserResponseDto, error) {
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

	mUserService := h.Service.WithTx(db)

	updates, associations := mapper.UpdateMUserDtoToMap(input)

	updatedData, err := mUserService.UpdateMUser(id, updates, associations)
	if err != nil {
		return nil, err
	}

	if err := db.Commit().Error; err != nil {
		return nil, err
	}
	committed = true

	return mapper.MUserModelToResponseDto(updatedData)
}

func (h *MUserHandler) BulkCreateMUsersHandler(input *dto.BulkCreateMUserDto) ([]dto.MUserResponseDto, error) {
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

	mUserService := h.Service.WithTx(db)

	var users []*models.MUsers
	for _, createDto := range input.Data {
		mUser, err := mapper.CreateMUserDtoToModel(&createDto)
		if err != nil {
			return nil, err
		}
		users = append(users, mUser)
	}

	created, err := mUserService.BulkCreateMUsers(users)
	if err != nil {
		return nil, err
	}

	if err := db.Commit().Error; err != nil {
		return nil, err
	}
	committed = true

	var output []dto.MUserResponseDto
	for _, m := range created {
		item, err := mapper.MUserModelToResponseDto(m)
		if err != nil {
			return nil, err
		}
		output = append(output, *item)
	}

	return output, nil
}

func (h *MUserHandler) BulkUpdateMUsersHandler(input *dto.BulkUpdateMUserDto) error {
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

	mUserService := h.Service.WithTx(db)

	if len(input.IDs) == 0 {
		return fmt.Errorf("ids cannot be empty")
	}

	if input.Data == nil {
		return fmt.Errorf("data cannot be nil")
	}

	payload, associations := mapper.UpdateMUserDtoToMap(input.Data)
	if len(payload) == 0 {
		return fmt.Errorf("update data cannot be empty")
	}

	if err := mUserService.BulkUpdateMUsers(input.IDs, payload, associations); err != nil {
		return fmt.Errorf("failed to bulk update m_users: %w", err)
	}

	if err := db.Commit().Error; err != nil {
		return err
	}
	committed = true

	return nil
}

func (h *MUserHandler) BulkDeleteMUsersHandler(input *dto.BulkDeleteMUserDto, isPermanent bool) error {
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

	mUserService := h.Service.WithTx(db)

	if err := mUserService.BulkDeleteMUsers(input.IDs, isPermanent); err != nil {
		return err
	}

	if err := db.Commit().Error; err != nil {
		return err
	}
	committed = true

	return nil
}

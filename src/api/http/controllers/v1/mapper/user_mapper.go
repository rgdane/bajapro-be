package mapper

import (
	"crypto/rand"
	"jk-api/api/http/controllers/v1/dto"
	"jk-api/internal/database/models"
	"math/big"
)

func CreateUserDtoToModel(dto *dto.CreateUserDto) (*models.User, error) {
	if dto == nil {
		return nil, nil
	}

	chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()"
	generatedPassword := make([]byte, 8)
	for i := range generatedPassword {
		num, _ := rand.Int(rand.Reader, big.NewInt(int64(len(chars))))
		generatedPassword[i] = chars[num.Int64()]
	}

	role := models.UserRole(dto.Role)
	if !role.Valid() {
		role = models.RoleStudent
	}

	data := &models.User{
		Code:     dto.Code,
		Name:     dto.Name,
		Email:    dto.Email,
		Role:     role,
		Nip:      dto.Nip,
		SchoolID: dto.SchoolID,
		ClassID:  dto.ClassID,
		Password: func() string {
			if dto.Password != "" {
				return dto.Password
			}
			return string(generatedPassword)
		}(),
		IsPasswordDefault: func() bool {
			if dto.IsPasswordDefault != nil {
				return *dto.IsPasswordDefault
			}
			return true
		}(),
	}

	return data, nil
}

func UpdateUserDtoToModel(dto *dto.UpdateUserDto) (
	payload map[string]interface{},
	associations map[string]interface{},
) {
	payload = make(map[string]interface{})
	associations = make(map[string]interface{})

	if dto.Code != nil {
		payload["code"] = *dto.Code
	}
	if dto.Name != nil {
		payload["name"] = *dto.Name
	}
	if dto.Email != nil {
		payload["email"] = *dto.Email
	}
	if dto.Password != nil {
		payload["password"] = *dto.Password
	}
	if dto.Nip != nil {
		payload["nip"] = *dto.Nip
	}
	if dto.Role != nil {
		payload["role"] = *dto.Role
	}
	if dto.SchoolID != nil {
		payload["school_id"] = *dto.SchoolID
	}
	if dto.ClassID != nil {
		payload["class_id"] = *dto.ClassID
	}

	if dto.NewPassword != nil && dto.OldPawword != nil {
		payload["new_password"] = *dto.NewPassword
		payload["old_password"] = *dto.OldPawword
	}
	if dto.Approval != nil {
		payload["approval"] = *dto.Approval
	}
	payload["deleted_at"] = dto.DeletedAt

	return payload, associations
}

func UserModelToResponseDto(data *models.User) (*dto.UserResponseDto, error) {
	if data == nil {
		return nil, nil
	}

	responseDto := &dto.UserResponseDto{
		User: *data,
	}

	return responseDto, nil
}

package seeders

import (
	"jk-api/internal/database/models"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func SeedUsers(db *gorm.DB) error {
	var teacherRole, studentRole models.Role

	if err := db.Where("name = ?", "teacher").First(&teacherRole).Error; err != nil {
		return err
	}

	if err := db.Where("name = ?", "student").First(&studentRole).Error; err != nil {
		return err
	}

	password, err := bcrypt.GenerateFromPassword(
		[]byte("123456"),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return err
	}

	studentClassID := int64(1)

	users := []struct {
		Name     string
		Email    string
		Role     models.Role
		ClassID  *int64
		Approval int8
	}{
		{
			Name:     "Teacher 1",
			Email:    "teacher@mail.com",
			Role:     teacherRole,
			ClassID:  nil,
			Approval: 0,
		},
		{
			Name:     "Student 1",
			Email:    "student@mail.com",
			Role:     studentRole,
			ClassID:  &studentClassID,
			Approval: 1,
		},
	}

	for _, item := range users {
		var user models.User

		err := db.Where("email = ?", item.Email).First(&user).Error

		switch {
		case err == gorm.ErrRecordNotFound:
			user = models.User{
				Name:              item.Name,
				Email:             item.Email,
				ClassID:           item.ClassID,
				Password:          string(password),
				Approval:          item.Approval,
				IsActive:          true,
				CreatedAt:         time.Now(),
				UpdatedAt:         time.Now(),
			}

			if err := db.Create(&user).Error; err != nil {
				return err
			}

		case err != nil:
			return err

		default:
			// Pastikan user lama juga mendapatkan class_id terbaru.
			if err := db.Model(&user).Updates(map[string]any{
				"name":                   item.Name,
				"class_id":               item.ClassID,
				"approval":               item.Approval,
				"isactive":               true,
				"updated_at":             time.Now(),
			}).Error; err != nil {
				return err
			}
		}

		// Replace membuat role tetap konsisten dan tidak duplikat.
		if err := db.Model(&user).
			Association("HasRoles").
			Replace(&item.Role); err != nil {
			return err
		}
	}

	return nil
}
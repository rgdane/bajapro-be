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

	password, _ := bcrypt.GenerateFromPassword([]byte("123456"), bcrypt.DefaultCost)

	users := []struct {
		Name     string
		Email    string
		Role     models.Role
		Approved bool
	}{
		{"Teacher 1", "teacher@mail.com", teacherRole, false}, // pending
		{"Student 1", "student@mail.com", studentRole, true},  // langsung aktif
	}

	for _, u := range users {
		var existing models.User
		err := db.Where("email = ?", u.Email).First(&existing).Error

		if err == gorm.ErrRecordNotFound {
			user := models.User{
				Name:              u.Name,
				Email:             u.Email,
				Password:          string(password),
				IsApprovedByAdmin: u.Approved,
				IsActive:          true,
				CreatedAt:         time.Now(),
				UpdatedAt:         time.Now(),
			}

			if err := db.Create(&user).Error; err != nil {
				return err
			}

			if err := db.Model(&user).Association("HasRoles").Append(&u.Role); err != nil {
				return err
			}
		} else if err != nil {
			return err
		}
	}

	return nil
}
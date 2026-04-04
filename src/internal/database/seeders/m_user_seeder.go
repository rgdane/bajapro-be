package seeders

import (
	"jk-api/internal/database/models"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func SeedMUsers(db *gorm.DB) error {
	var studentRole models.MRole
	if err := db.Where("name = ?", "student").First(&studentRole).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			studentRole = models.MRole{
				RoleName:  "student",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			}
			if err := db.Create(&studentRole).Error; err != nil {
				return err
			}
		} else {
			return err
		}
	}

	var studentClass models.MClass
	if err := db.Where("name = ?", "class_a").First(&studentClass).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			studentClass = models.MClass{
				ClassName: "class_a",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			}
			if err := db.Create(&studentClass).Error; err != nil {
				return err
			}
		} else {
			return err
		}
	}

	var existing models.MUsers
	if err := db.Where("email = ?", "pelajar@gmail.com").First(&existing).Error; err == nil {
		return nil
	} else if err != gorm.ErrRecordNotFound {
		return err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("123456"), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	student := models.MUsers{
		RoleID:            studentRole.ID,
		ClassID:           studentClass.ID,
		Name:              "User Pelajar",
		Email:             "pelajar@gmail.com",
		Password:          string(hashedPassword),
		IsApprovedByAdmin: true,
		InstansiSekolah:   "School A",
		IsActive:          true,
		CreatedAt:         time.Now(),
		UpdatedAt:         time.Now(),
	}

	if err := db.Create(&student).Error; err != nil {
		return err
	}

	return nil
}

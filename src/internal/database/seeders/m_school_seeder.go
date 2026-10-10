package seeders

import (
	"jk-api/internal/database/models"

	"gorm.io/gorm"
)

func SeedSchools(db *gorm.DB) error {
	var count int64
	db.Model(&models.MSchool{}).Count(&count)
	if count > 0 {
		return nil
	}

	schools := []models.MSchool{
		{SchoolName: "SMK Negeri 1 Jakarta", IsActive: true},
		{SchoolName: "SMK Negeri 2 Bandung", IsActive: true},
	}

	return db.Create(&schools).Error
}

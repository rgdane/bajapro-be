package seeders

import (
	"jk-api/internal/database/models"

	"gorm.io/gorm"
)

func SeedCourses(db *gorm.DB) error {
	var count int64
	db.Model(&models.MCourse{}).Count(&count)

	if count > 0 {
		return nil // biar tidak duplicate
	}

	courses := []models.MCourse{
		{
			CourseName:   "Java",
			Description:  "Java is a high-level, class-based, object-oriented programming language that is designed to have as few implementation dependencies as possible",
			ImgThumbnail: "java.png",
			Published:    true,
			IsActive:     true,
		},
	}

	return db.Create(&courses).Error
}
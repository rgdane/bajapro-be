package seeders

import (
	"jk-api/internal/database/models"

	"gorm.io/gorm"
)

func SeedLevels(db *gorm.DB) error {
	var count int64
	db.Model(&models.MLevel{}).Count(&count)

	if count > 0 {
		return nil // biar tidak duplicate
	}

	levels := []models.MLevel{
		{
			LevelName:   "Easy",
			Description: "Level untuk pemula yang baru memulai belajar.",
			IsActive:    true,
		},
		{
			LevelName:   "Medium",
			Description: "Level untuk yang sudah memahami dasar dan ingin meningkatkan skill.",
			IsActive:    true,
		},
		{
			LevelName:   "Hard",
			Description: "Level untuk yang sudah mahir dan siap ke materi kompleks.",
			IsActive:    true,
		},
		{
			LevelName:   "Advanced",
			Description: "Level tertinggi untuk yang sudah profesional.",
			IsActive:    true,
		},
	}

	return db.Create(&levels).Error
}
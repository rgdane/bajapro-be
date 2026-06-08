package seeders

import (
	"jk-api/internal/database/models"

	"gorm.io/gorm"
)

func SeedBadgeSettings(db *gorm.DB) error {
	var count int64
	db.Model(&models.MBadgeSettings{}).Count(&count)

	if count > 0 {
		return nil // biar tidak duplicate
	}

	badges := []models.MBadgeSettings{
		{
			Name:     "Warrior",
			Image:    "warrior.png",
			MinScore: 0,
			MaxScore: 50,
			IsActive: true,
		},
		{
			Name:     "Elite",
			Image:    "elite.png",
			MinScore: 51,
			MaxScore: 100,
			IsActive: true,
		},
		{
			Name:     "Master",
			Image:    "master.png",
			MinScore: 101,
			MaxScore: 150,
			IsActive: true,
		},
		{
			Name:     "Grandmaster",
			Image:    "grandmaster.png",
			MinScore: 151,
			MaxScore: 200,
			IsActive: true,
		},
		{
			Name:     "Epic",
			Image:    "epic.png",
			MinScore: 201,
			MaxScore: 300,
			IsActive: true,
		},
		{
			Name:     "Legend",
			Image:    "legend.png",
			MinScore: 301,
			MaxScore: 500,
			IsActive: true,
		},
		{
			Name:     "Mythic",
			Image:    "mythic.png",
			MinScore: 501,
			MaxScore: 800,
			IsActive: true,
		},
		{
			Name:     "Diamond",
			Image:    "diamond.png",
			MinScore: 801,
			MaxScore: 9999,
			IsActive: true,
		},
	}

	return db.Create(&badges).Error
}

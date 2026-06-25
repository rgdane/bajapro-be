package seeders

import (
	"jk-api/internal/database/models"

	"gorm.io/gorm"
)

func SeedClasses(db *gorm.DB) error {
	var count int64
	db.Model(&models.MClass{}).Count(&count)

	if count > 0 {
		return nil // biar tidak duplicate
	}

	classes := []models.MClass{
		{
			ClassName:  "X RPL 1",
			SchoolName: "SMK Negeri 1 Jakarta",
			ClassCode:  "RPL-X-001",
		},
		{
			ClassName:  "X RPL 2",
			SchoolName: "SMK Negeri 1 Jakarta",
			ClassCode:  "RPL-X-002",
		},
		{
			ClassName:  "XI RPL 1",
			SchoolName: "SMK Negeri 1 Jakarta",
			ClassCode:  "RPL-XI-001",
		},
		{
			ClassName:  "XII RPL 1",
			SchoolName: "SMK Negeri 1 Jakarta",
			ClassCode:  "RPL-XII-001",
		},
		{
			ClassName:  "X TKJ 1",
			SchoolName: "SMK Negeri 2 Bandung",
			ClassCode:  "TKJ-X-001",
		},
	}

	return db.Create(&classes).Error
}

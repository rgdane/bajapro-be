package seeders

import (
	"jk-api/internal/database/models"

	"gorm.io/gorm"
)

func SeedClasses(db *gorm.DB) error {
	var count int64
	db.Model(&models.MClass{}).Count(&count)

	if count > 0 {
		return nil
	}

	var school1, school2 models.MSchool
	db.Where("LOWER(school_name) = ?", "smk negeri 1 jakarta").First(&school1)
	db.Where("LOWER(school_name) = ?", "smk negeri 2 bandung").First(&school2)

	if school1.ID == 0 || school2.ID == 0 {
		return nil
	}

	classes := []models.MClass{
		{
			ClassName: "X RPL 1",
			ClassCode: "RPL-X-001",
			SchoolID:  school1.ID,
		},
		{
			ClassName: "X RPL 2",
			ClassCode: "RPL-X-002",
			SchoolID:  school1.ID,
		},
		{
			ClassName: "XI RPL 1",
			ClassCode: "RPL-XI-001",
			SchoolID:  school1.ID,
		},
		{
			ClassName: "XII RPL 1",
			ClassCode: "RPL-XII-001",
			SchoolID:  school1.ID,
		},
		{
			ClassName: "X TKJ 1",
			ClassCode: "TKJ-X-001",
			SchoolID:  school2.ID,
		},
	}

	return db.Create(&classes).Error
}

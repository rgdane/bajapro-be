package seeders

import (
	"jk-api/internal/database/models"

	"gorm.io/gorm"
)

func SeedLessons(db *gorm.DB) error {
	var count int64
	db.Model(&models.MLesson{}).Count(&count)

	if count > 0 {
		return nil // biar tidak duplicate
	}
	
	lessons := []models.MLesson{
		{
			CourseID:     1,
			LevelID:      1,
			Title:        "Pengenalan Dasar",
			Description:  "Materi pengenalan dasar untuk pemula",
			Position:     1,
			ImgThumbnail: "https://example.com/img1.png",
		},
		{
			CourseID:     1,
			LevelID:      1,
			Title:        "Dasar Lanjutan",
			Description:  "Materi lanjutan dari pengenalan dasar",
			Position:     2,
			ImgThumbnail: "https://example.com/img2.png",
		},
		{
			CourseID:     2,
			LevelID:      2,
			Title:        "Intermediate Lesson",
			Description:  "Materi tingkat menengah",
			Position:     1,
			ImgThumbnail: "https://example.com/img3.png",
		},
	}

	for _, lesson := range lessons {
		if err := db.Create(&lesson).Error; err != nil {
			return err
		}
	}

	return nil
}

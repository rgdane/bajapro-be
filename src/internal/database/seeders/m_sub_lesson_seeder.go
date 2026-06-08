package seeders

import (
	"jk-api/internal/database/models"

	"gorm.io/gorm"
)

func SeedSubLessons(db *gorm.DB) error {
	subLessons := []models.MSubLesson{
		{
			LessonID:      1,
			Title:         "Pengenalan Sub Materi 1",
			OrderPosition: 1,
		},
		{
			LessonID:      1,
			Title:         "Pengenalan Sub Materi 2",
			OrderPosition: 2,
		},
		{
			LessonID:      2,
			Title:         "Sub Materi Lanjutan 1",
			OrderPosition: 1,
		},
		{
			LessonID:      2,
			Title:         "Sub Materi Lanjutan 2",
			OrderPosition: 2,
		},
	}

	for _, sub := range subLessons {
		if err := db.Create(&sub).Error; err != nil {
			return err
		}
	}

	return nil
}

package seeders

import (
	"jk-api/internal/database/models"

	"gorm.io/gorm"
)

func SeedSubLessons(db *gorm.DB) error {
	var count int64
	db.Model(&models.MSubLesson{}).Count(&count)

	if count > 0 {
		return nil // biar tidak duplicate
	}
	subLessons := []models.MSubLesson{
		{
			LessonID:      1,
			Title:         "Tipe Data, Variabel dan Operator",
			OrderPosition: 1,
		},
		{
			LessonID:      1,
			Title:         "Sintaks Pemilihan IF-ELSE",
			OrderPosition: 2,
		},
		{
			LessonID:      2,
			Title:         "Sintaks Pemilihan IF-ELSE",
			OrderPosition: 1,
		},
		{
			LessonID:      2,
			Title:         "Sintaks Pemilihan Switch-Case",
			OrderPosition: 2,
		},
		{
			LessonID:      3,
			Title:         "Sintaks Pemilihan Bersarang",
			OrderPosition: 1,
		},
		{
			LessonID:      4,
			Title:         "Perulangan dengan For",
			OrderPosition: 1,
		},
		{
			LessonID:      4,
			Title:         "Perulangan dengan While",
			OrderPosition: 2,
		},
		{
			LessonID:      5,
			Title:         "Perulangan dengan Do-While",
			OrderPosition: 1,
		},
		{
			LessonID:      5,
			Title:         "Perulangan Bersarang",
			OrderPosition: 2,
		},
		{
			LessonID:      6,
			Title:         "Array Satu Dimensi",
			OrderPosition: 1,
		},
		{
			LessonID:      7,
			Title:         "Array Multidimensi",
			OrderPosition: 1,
		},
		{
			LessonID:      8,
			Title:         "Fungsi Static",
			OrderPosition: 1,
		},
		{
			LessonID:      9,
			Title:         "Fungsi Rekursif",
			OrderPosition: 1,
		},
	}

	for _, sub := range subLessons {
		if err := db.Create(&sub).Error; err != nil {
			return err
		}
	}

	return nil
}

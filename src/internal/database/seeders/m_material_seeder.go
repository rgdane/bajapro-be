package seeders

import (
	"jk-api/internal/database/models"

	"gorm.io/gorm"
)

func SeedMaterials(db *gorm.DB) error {
	var count int64
	db.Model(&models.MMaterials{}).Count(&count)

	if count > 0 {
		return nil // biar tidak duplicate
	}

	materials := []models.MMaterials{
		{
			SubLessonID:     1,
			Title:           "Pengantar Materi",
			Materials:       "Materi dasar untuk memahami konsep awal.",
			URLVideo:        "https://www.youtube.com/watch?v=video1",
			ContentPosition: 1,
			PromptLLM:       "Jelaskan materi ini dengan bahasa sederhana",
		},
		{
			SubLessonID:     1,
			Title:           "Pendalaman Materi",
			Materials:       "Penjelasan lebih dalam terkait materi sebelumnya.",
			URLVideo:        "https://www.youtube.com/watch?v=video2",
			ContentPosition: 2,
			PromptLLM:       "Buatkan contoh kasus dari materi ini",
		},
		{
			SubLessonID:     2,
			Title:           "Latihan Pemahaman",
			Materials:       "Latihan soal untuk menguji pemahaman.",
			URLVideo:        "",
			ContentPosition: 1,
			PromptLLM:       "Buatkan soal latihan dari materi ini",
		},
	}

	for _, m := range materials {
		var existing models.MMaterials

		err := db.Where("title = ? AND sub_lesson_id = ?", m.Title, m.SubLessonID).
			First(&existing).Error

		if err == gorm.ErrRecordNotFound {
			if err := db.Create(&m).Error; err != nil {
				return err
			}
		} else if err != nil {
			return err
		}
	}

	return nil
}
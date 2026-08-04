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
			Title:        "Tipe Data, Variabel dan Operator",
			Description:  "User dapat memahami dan mampu menjelaskan tentang Tipe Data, Variabel, Input-Output, Sequence dan operator pada pemrograman Java.",
			Position:     1,
			ImgThumbnail: "https://example.com/img1.png",
		},
		{
			CourseID:     1,
			LevelID:      1,
			Title:        "Sintaks Pemilihan 1",
			Description:  "Materi lanjutan dari pengenalan dasar sintaks pemilihan.",
			Position:     2,
			ImgThumbnail: "https://example.com/img2.png",
		},
		{
			CourseID:     1,
			LevelID:      1,
			Title:        "Sintaks Pemilihan 2",
			Description:  "Materi lanjutan dari sintaks pemilihan 1.",
			Position:     3,
			ImgThumbnail: "https://example.com/img3.png",
		},
		{
			CourseID:     1,
			LevelID:      2,
			Title:        "Sintaks Perulangan 1",
			Description:  "Materi tentang penggunaan sintaks perulangan.",
			Position:     4,
			ImgThumbnail: "https://example.com/img4.png",
		},
		{
			CourseID:     1,
			LevelID:      2,
			Title:        "Sintaks Perulangan 2",
			Description:  "Materi lanjutan mengenai sintaks perulangan.",
			Position:     5,
			ImgThumbnail: "https://example.com/img5.png",
		},
		{
			CourseID:     1,
			LevelID:      2,
			Title:        "Array 1",
			Description:  "Materi pengenalan array satu dimensi.",
			Position:     6,
			ImgThumbnail: "https://example.com/img6.png",
		},
		{
			CourseID:     1,
			LevelID:      3,
			Title:        "Array Multidimensi",
			Description:  "Materi mengenai penggunaan array multidimensi.",
			Position:     7,
			ImgThumbnail: "https://example.com/img7.png",
		},
		{
			CourseID:     1,
			LevelID:      3,
			Title:        "Fungsi 1",
			Description:  "Materi dasar mengenai fungsi atau method pada Java.",
			Position:     8,
			ImgThumbnail: "https://example.com/img8.png",
		},
		{
			CourseID:     1,
			LevelID:      3,
			Title:        "Fungsi 2",
			Description:  "Materi lanjutan mengenai fungsi atau method pada Java.",
			Position:     9,
			ImgThumbnail: "https://example.com/img9.png",
		},
	}

	for _, lesson := range lessons {
		if err := db.Create(&lesson).Error; err != nil {
			return err
		}
	}

	return nil
}

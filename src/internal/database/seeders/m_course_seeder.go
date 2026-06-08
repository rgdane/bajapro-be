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
			CourseName:   "Basic Programming",
			Description:  "Belajar dasar-dasar pemrograman dari nol.",
			ImgThumbnail: "basic_programming.png",
			Published:    true,
			IsActive:     true,
		},
		{
			CourseName:   "Web Development",
			Description:  "Mempelajari HTML, CSS, dan JavaScript untuk membuat website.",
			ImgThumbnail: "web_dev.png",
			Published:    true,
			IsActive:     true,
		},
		{
			CourseName:   "Backend Development (Golang)",
			Description:  "Belajar membuat API menggunakan Golang dan Fiber.",
			ImgThumbnail: "golang_backend.png",
			Published:    true,
			IsActive:     true,
		},
		{
			CourseName:   "Mobile Development (Flutter)",
			Description:  "Membuat aplikasi mobile menggunakan Flutter.",
			ImgThumbnail: "flutter.png",
			Published:    false,
			IsActive:     true,
		},
		{
			CourseName:   "Database Design",
			Description:  "Mempelajari perancangan database menggunakan PostgreSQL.",
			ImgThumbnail: "database.png",
			Published:    true,
			IsActive:     true,
		},
	}

	return db.Create(&courses).Error
}
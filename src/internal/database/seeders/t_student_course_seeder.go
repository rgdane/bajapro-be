package seeders

import (
	"jk-api/internal/database/models"
	"gorm.io/gorm"
)

func SeedStudentCourses(db *gorm.DB) error {
	var count int64

	// cek apakah sudah ada data
	if err := db.Model(&models.TStudentCourse{}).Count(&count).Error; err != nil {
		return err
	}

	// kalau sudah ada, skip (anti duplicate global)
	if count > 0 {
		return nil
	}

	data := []models.TStudentCourse{
		{
			UserID:     1,
			CourseID:   1,
			TotalScore: 80,
			BadgeID:    1,
		},
		{
			UserID:     2,
			CourseID:   1,
			TotalScore: 90,
			BadgeID:    2,
		},
		{
			UserID:     2,
			CourseID:   2,
			TotalScore: 75,
			BadgeID:    1,
		},
	}

	for _, item := range data {
		var existing models.TStudentCourse

		err := db.
			Where("user_id = ? AND course_id = ?", item.UserID, item.CourseID).
			First(&existing).Error

		if err == gorm.ErrRecordNotFound {
			if err := db.Create(&item).Error; err != nil {
				return err
			}
		} else if err != nil {
			return err
		}
	}

	return nil
}
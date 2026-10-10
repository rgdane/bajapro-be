package seeders

import (
	"jk-api/internal/database/models"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func SeedUsers(db *gorm.DB) error {
	password, err := bcrypt.GenerateFromPassword(
		[]byte("123456"),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return err
	}

	studentClassID := int64(1)

	var school models.MSchool
	db.Where("LOWER(school_name) = ?", "smk negeri 1 jakarta").First(&school)
	var schoolID *int64
	if school.ID != 0 {
		schoolID = &school.ID
	}

	users := []struct {
		Name     string
		Email    string
		Role     models.UserRole
		SchoolID *int64
		ClassID  *int64
		Approval int8
	}{
		{
			Name:     "Admin 1",
			Email:    "admin@mail.com",
			Role:     models.RoleAdmin,
			SchoolID: nil,
			ClassID:  nil,
			Approval: 1,
		},
		{
			Name:     "Teacher 1",
			Email:    "teacher@mail.com",
			Role:     models.RoleTeacher,
			SchoolID: schoolID,
			ClassID:  nil,
			Approval: 0,
		},
		{
			Name:     "Student 1",
			Email:    "student@mail.com",
			Role:     models.RoleStudent,
			SchoolID: schoolID,
			ClassID:  &studentClassID,
			Approval: 1,
		},
	}

	for _, item := range users {
		var user models.User

		err := db.Where("email = ?", item.Email).First(&user).Error

		switch {
		case err == gorm.ErrRecordNotFound:
			user = models.User{
				Name:      item.Name,
				Email:     item.Email,
				Role:      item.Role,
				SchoolID:  item.SchoolID,
				ClassID:   item.ClassID,
				Password:  string(password),
				Approval:  item.Approval,
				IsActive:  true,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			}

			if err := db.Create(&user).Error; err != nil {
				return err
			}

		case err != nil:
			return err

		default:
			// Pastikan user lama juga mendapatkan class_id terbaru.
			if err := db.Model(&user).Updates(map[string]any{
				"name":       item.Name,
				"school_id":  item.SchoolID,
				"class_id":   item.ClassID,
				"role":       item.Role,
				"approval":   item.Approval,
				"isactive":   true,
				"updated_at": time.Now(),
			}).Error; err != nil {
				return err
			}
		}
	}

	// Tautkan teacher seeder sebagai pengajar kelas di sekolahnya,
	// supaya registrasi student via kode kelas punya teacher valid.
	var teacher models.User
	if err := db.Where("email = ?", "teacher@mail.com").First(&teacher).Error; err == nil && teacher.SchoolID != nil {
		if err := db.Model(&models.MClass{}).
			Where("school_id = ? AND teacher_id IS NULL", *teacher.SchoolID).
			Update("teacher_id", teacher.ID).Error; err != nil {
			return err
		}
	}

	return nil
}

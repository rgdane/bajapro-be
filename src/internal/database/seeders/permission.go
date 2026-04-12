package seeders

import (
	"jk-api/internal/database/models"
	"time"

	"gorm.io/gorm"
)

func SeedPermissions(db *gorm.DB) error {
	permissionsMap := map[string][]string{
		"roles":               {"create", "update", "delete", "view", "viewOwn"},
		"permissions":         {"create", "update", "delete", "view", "viewOwn"},
		"users":               {"create", "update", "delete", "view", "viewOwn"},
		"classes":             {"create", "update", "delete", "view", "viewOwn"},
		"m_badge_settings":      {"create", "update", "delete", "view", "viewOwn"},
	}

	for module, actions := range permissionsMap {
		for _, action := range actions {
			permName := module + "." + action

			var count int64
			db.Model(&models.Permission{}).Where("name = ?", permName).Count(&count)
			if count == 0 {
				db.Create(&models.Permission{
					Name:      permName,
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				})
			}
		}
	}

	return nil
}

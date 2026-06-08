package migrations

import (
	"log"

	"gorm.io/gorm"
)

func SetupSequenceTable(db *gorm.DB) error {
	sequences := map[string]string{
		"users_seq":           "users",
		"squads_seq":          "squads",
		"roles_seq":           "roles",
		"permissions_seq":     "permissions",
		"m_classes_seq":        "m_class",
		"m_badge_settings_seq": "m_badge_settings",
		"m_courses_seq":        "m_course",
		"m_levels_seq":         "m_level",
		"m_lessons_seq":        "m_lesson",
		"m_sub_lessons_seq":    "m_sub_lesson",
		"m_materials_seq":      "m_material",
	}

	for seqName, tableName := range sequences {
		var existingSeq string

		err := db.Raw(`
			SELECT sequence_name 
			FROM information_schema.sequences 
			WHERE sequence_name = ? AND sequence_schema = 'public'
		`, seqName).Scan(&existingSeq).Error

		if err != nil {
			log.Printf("Error checking sequence %s: %v", seqName, err)
			continue
		}

		if existingSeq == "" {
			createSeqSQL := `
				CREATE SEQUENCE IF NOT EXISTS ` + seqName + `
					START 1
					INCREMENT 1
					MINVALUE 1
					MAXVALUE 9007199254740991;
			`

			if err := db.Exec(createSeqSQL).Error; err != nil {
				log.Printf("Failed to create sequence %s for table %s: %v", seqName, tableName, err)
				continue
			}
		} else {
			log.Printf("📝 Sequence already exists: %s", seqName)
		}
	}

	return nil
}

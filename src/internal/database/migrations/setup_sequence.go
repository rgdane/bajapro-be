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
		"projects_seq":        "projects",
		"departments_seq":     "departments",
		"divisions_seq":       "divisions",
		"levels_seq":          "levels",
		"positions_seq":       "positions",
		"titles_seq":          "titles",
		"statuses_seq":        "statuses",
		"holidays_seq":        "holidays",
		"leaves_seq":          "leaves",
		"backlog_items_seq":   "backlog_items",
		"sprints_seq":         "sprints",
		"backlogs_seq":        "backlogs",
		"comments_seq":        "comments",
		"activity_logs_seq":   "activity_logs",
		"notifications_seq":   "notifications",
		"cms_articles_seq":    "cms_articles",
		"cms_categories_seq":  "cms_categories",
		"cms_tags_seq":        "cms_tags",
		"documents_seq":       "documents",
		"sops_seq":            "sops",
		"spks_seq":            "spks",
		"spk_jobs_seq":        "spk_jobs",
		"spk_txs_seq":         "spk_txs",
		"spk_tx_results_seq":  "spk_tx_results",
		"spk_tx_versions_seq": "spk_tx_versions",
		"sop_jobs_seq":        "sop_jobs",
		"products_seq":        "products",
		"epics_seq":           "epics",
		"jobs_seq":            "jobs",
		"flowcharts_seq":      "flowcharts",
		"sprint_goals_seq":    "sprint_goals",
		"sprint_retros_seq":   "sprint_retros",
		"retro_items_seq":     "retro_items",
		"sprint_dailies_seq":  "sprint_dailies",
		"languages_seq":      "languages",
		"color_palettes_seq":  "color_palettes",
		"sop_menus_seq":        "sop_menus",
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

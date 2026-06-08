package migrations

import (
	"jk-api/internal/config"
	"jk-api/internal/database/models"
	"log"
)

func Migrate() {
	db := config.DB

	if err := SetupSequenceTable(db); err != nil {
		log.Fatalf("❌ Failed to create sequences: %v", err)
	}

	// if err := SetupJoinTable(db); err != nil {
	// 	log.Fatalf("❌ Failed to setup join table: %v", err)
	// }

	err := db.AutoMigrate(
		&models.User{},
		&models.MLevel{},
		&models.Role{},
		&models.MClass{},
		&models.MLesson{},
		&models.MSubLesson{},
		&models.MBadgeSettings{},
		&models.MCourse{},
		&models.MMaterials{},
		&models.TStudentCourse{},
		&models.TStudentProgress{},
		&models.TWonderingScore{},
		&models.TCodeQuestion{},
		&models.TCodeAnswer{},
		&models.TCodeHistoryLogs{},
		&models.TEssayQuestion{},
		&models.TEssayAnswer{},
		&models.TGenerationHistory{},
		&models.Permission{},
	)

	// if err := NotificationPriorityEnum(db); err != nil {
	// 	log.Fatal("❌ Migration failed:", err)
	// }

	// if err := AddBacklogPriorityCheck(db); err != nil {
	// 	log.Fatal("❌ Migration failed:", err)
	// }

	if err != nil {
		log.Fatalf("❌ Migration failed: %v", err)
	}

	log.Println("✅ Migration complete")
}

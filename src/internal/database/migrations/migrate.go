package migrations

import (
	"fmt"
	"jk-api/internal/config"
	"jk-api/internal/database/models"
	"log"

	"gorm.io/gorm"
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
		&models.RefreshToken{},
		&models.MSchool{},
		&models.RegistrationRequest{},
		&models.RegistrationRequestSchool{},
		&models.MLevel{},
		&models.MClass{},
		&models.MLesson{},
		&models.MSubLesson{},
		&models.MBadgeSettings{},
		&models.MCourse{},
		&models.MMaterials{},
		&models.TStudentCourse{},
		&models.TStudentProgress{},
		&models.TWonderingScore{},
		&models.CodeQuestion{},
		&models.TCodeAnswer{},
		&models.TCodeHistoryLogs{},
		&models.EssayQuestion{},
		&models.TEssayAnswer{},
		&models.Leaderboard{},
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

	MigrateData(db)
	EnsureManualFks(db)

	log.Println("✅ Migration complete")
}

// addForeignKeyIfMissing menambahkan FK constraint jika belum ada di database.
func addForeignKeyIfMissing(db *gorm.DB, constraint, table, keys, refs, onDelete string) {
	var count int64
	if err := db.Raw(
		`SELECT count(*) FROM pg_constraint WHERE conname = ?`,
		constraint,
	).Scan(&count).Error; err != nil || count > 0 {
		return
	}

	db.Exec(fmt.Sprintf(
		"ALTER TABLE %s ADD CONSTRAINT %s FOREIGN KEY (%s) REFERENCES %s (%s) ON UPDATE CASCADE ON DELETE %s",
		table, constraint, keys, refs, "id", onDelete,
	))
}

// EnsureManualFks menambahkan constraint FK untuk relasi yang sengaja
// ditandai `-:migration` agar AutoMigrate GORM tidak membuat FK saling-silang
// (users<->m_class) yang membuat ReorderModels menaruh m_class sebelum users.
func EnsureManualFks(db *gorm.DB) {
	addForeignKeyIfMissing(db, "fk_users_class", "users", "class_id", "m_class", "SET NULL")
	addForeignKeyIfMissing(db, "fk_m_class_teacher", "m_class", "teacher_id", "users", "SET NULL")
}

func MigrateData(db *gorm.DB) {
	// 1. Migrasi school_instance dari users ke m_school jika kolom masih ada
	if db.Migrator().HasColumn("users", "school_instance") {
		db.Exec(`
			INSERT INTO m_school (school_name, isactive, created_at, updated_at)
			SELECT DISTINCT school_instance, true, NOW(), NOW()
			FROM users
			WHERE school_instance IS NOT NULL AND school_instance != ''
			  AND NOT EXISTS (
				SELECT 1 FROM m_school s
				WHERE LOWER(TRIM(s.school_name)) = LOWER(TRIM(users.school_instance))
			  );
		`)

		db.Exec(`
			UPDATE users u
			SET school_id = s.id
			FROM m_school s
			WHERE u.school_instance IS NOT NULL AND LOWER(TRIM(u.school_instance)) = LOWER(TRIM(s.school_name))
			  AND u.school_id IS NULL;
		`)
	}

	// 2. Migrasi school_name dari m_class ke m_school jika kolom masih ada
	if db.Migrator().HasColumn("m_class", "school_name") {
		db.Exec(`
			INSERT INTO m_school (school_name, isactive, created_at, updated_at)
			SELECT DISTINCT school_name, true, NOW(), NOW()
			FROM m_class
			WHERE school_name IS NOT NULL AND school_name != ''
			  AND NOT EXISTS (
				SELECT 1 FROM m_school s
				WHERE LOWER(TRIM(s.school_name)) = LOWER(TRIM(m_class.school_name))
			  );
		`)

		db.Exec(`
			UPDATE m_class c
			SET school_id = s.id
			FROM m_school s
			WHERE c.school_name IS NOT NULL AND LOWER(TRIM(c.school_name)) = LOWER(TRIM(s.school_name))
			  AND (c.school_id IS NULL OR c.school_id = 0);
		`)
	}

	// 3. Migrasi pivot m_class_teachers ke m_class.teacher_id jika pivot masih ada
	if db.Migrator().HasTable("m_class_teachers") {
		db.Exec(`
			UPDATE m_class c
			SET teacher_id = ct.user_id
			FROM m_class_teachers ct
			WHERE c.id = ct.m_class_id AND c.teacher_id IS NULL;
		`)
	}
}

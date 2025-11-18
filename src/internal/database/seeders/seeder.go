package seeders

import "gorm.io/gorm"

func InitSeeder(db *gorm.DB) {
	SeedPermissions(db)
	SeedRoles(db)
	SeedAdmin(db)
	// SeedFlowcharts(db)
	// SeedDepartments(db)
	// SeedDivisions(db)
	// SeedLevels(db)
	// SeedPositions(db)
	// SeedTitles(db)
	// SeedProjects(db)
	// SeedSquads(db)
	// SeedSops(db)

	//TODO: Buat error check setiap seeder karyawan memiliki return error
}

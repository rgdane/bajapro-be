package bootstrap

import (
	"jk-api/internal/config"
	"jk-api/internal/container"

	"github.com/joho/godotenv"
)

func InitContainer() *container.AppContainer {
	config.InitPostgres()
	config.InitFirebaseApp()
	config.InitBucketApp()
	config.InitNeo4j()

	// migrations.Migrate()
	// seeders.InitSeeder(config.DB)

	services := container.NewAppContainer()
	return services
}

func InitEnv() {
	if err := godotenv.Load(); err != nil {
		Log.Warn("No .env file found, using environment variables")
	}
}

package bootstrap

import (
	"os"

	"jk-api/api/http/routes/v1"
	"jk-api/internal/config"
	"jk-api/internal/container"
)

func InitLogger() {
	config.StartLogger()
}

func LoadConfig() {
	config.LoadConfig()
}

func InitNeo4j() {
	if err := config.Neo4jApp(); err != nil {
		config.Logger.Fatalf("❌ Failed to initialize Neo4j: %v", err)
		return
	}
	config.Logger.Info("✅ Neo4j initialized")
}

func InitPostgres() {
	if err := config.PostgresApp(); err != nil {
		config.Logger.Fatalf("❌ Failed to initialize Postgres: %v", err)
		return
	}
	config.Logger.Info("✅ Postgres initialized")
}

func InitFiber() {
	app := config.InitFiberApp()

	// Sajikan file hasil upload (gambar) secara statis di /uploads.
	uploadDir := os.Getenv("UPLOAD_DIR")
	if uploadDir == "" {
		uploadDir = "./uploads"
	}
	_ = os.MkdirAll(uploadDir, 0o755)
	app.Static("/uploads", uploadDir)

	routes.Setup(app, container.NewAppContainer())
	config.Logger.Infof("✅ REST API started on port %s", config.AppConfig.AppPort)

	if err := app.Listen(":" + config.AppConfig.AppPort); err != nil {
		config.Logger.Fatalf("❌ Failed to start Fiber: %v", err)
		return
	}
}

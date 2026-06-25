package container

import (
	"jk-api/api/http/controllers/v1/handlers"
	"jk-api/pkg/repository/query/sql"
	"jk-api/pkg/services/v1"
)

func InitAuthContainer() *handlers.AuthHandler {
	userRepo := sql.NewUserRepository()
	refreshRepo := sql.NewRefreshTokenRepository()
	service := services.NewAuthService(userRepo, refreshRepo)
	return handlers.NewAuthHandler(service)
}

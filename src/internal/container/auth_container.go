package container

import (
	"jk-api/api/http/controllers/v1/handlers"
	"jk-api/pkg/repository/query/sql"
	"jk-api/pkg/services/v1"
)

func InitAuthContainer() *handlers.AuthHandler {
	repo := sql.NewUserRepository()
	service := services.NewAuthService(repo)
	return handlers.NewAuthHandler(service)
}

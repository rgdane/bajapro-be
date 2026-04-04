// m_user_container.go

package container

import (
	"jk-api/api/http/controllers/v1/handlers"
	"jk-api/pkg/repository/query/sql"
	"jk-api/pkg/services/v1"
)

func InitMUserContainer() *handlers.MUserHandler {
	repo := sql.NewMUsersRepository()
	service := services.NewMUserService(repo)
	return handlers.NewMUserHandler(service)
}

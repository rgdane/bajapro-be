package container

import (
	"jk-api/api/http/controllers/v1/handlers"
	"jk-api/pkg/repository/query/sql"
	"jk-api/pkg/services/v1"
)

func InitMSchoolContainer() *handlers.MSchoolHandler {
	repo := sql.NewMSchoolRepository()
	service := services.NewMSchoolService(repo)
	return handlers.NewMSchoolHandler(service)
}

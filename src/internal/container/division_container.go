package container

import (
	"jk-api/api/http/controllers/v1/handlers"
	"jk-api/pkg/repository/query/sql"
	"jk-api/pkg/services/v1"
)

func InitDivisionContainer() *handlers.DivisionHandler {
	repo := sql.NewDivisionRepository()
	service := services.NewDivisionService(repo)
	return handlers.NewDivisionHandler(service)
}

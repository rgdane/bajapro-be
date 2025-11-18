package container

import (
	"jk-api/api/http/controllers/v1/handlers"
)

type AppContainer struct {
	AuthHandler       *handlers.AuthHandler
	LevelHandler      *handlers.LevelHandler
	PermissionHandler *handlers.PermissionHandler
	PositionHandler   *handlers.PositionHandler
	RoleHandler       *handlers.RoleHandler
	TitleHandler      *handlers.TitleHandler
	UserHandler       *handlers.UserHandler
	DivisionHandler   *handlers.DivisionHandler
}

func NewAppContainer() *AppContainer {
	return &AppContainer{
		AuthHandler:       InitAuthContainer(),
		LevelHandler:      InitLevelContainer(),
		PermissionHandler: InitPermissionContainer(),
		PositionHandler:   InitPositionContainer(),
		RoleHandler:       InitRoleContainer(),
		TitleHandler:      InitTitleContainer(),
		UserHandler:       InitUserContainer(),
		DivisionHandler:   InitDivisionContainer(),
	}
}

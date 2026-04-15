package container

import (
	"jk-api/api/http/controllers/v1/handlers"
)

type AppContainer struct {
	AuthHandler       *handlers.AuthHandler
	MLevelHandler      *handlers.MLevelHandler
	PermissionHandler *handlers.PermissionHandler
	RoleHandler       *handlers.RoleHandler
	UserHandler       *handlers.UserHandler
	MClassHandler     *handlers.MClassHandler
	MBadgeSettingsHandler *handlers.MBadgeSettingsHandler
	MCourseHandler    *handlers.MCourseHandler
}

func NewAppContainer() *AppContainer {
	return &AppContainer{
		AuthHandler:       InitAuthContainer(),
		MLevelHandler:      InitMLevelContainer(),
		PermissionHandler: InitPermissionContainer(),
		RoleHandler:       InitRoleContainer(),
		UserHandler:       InitUserContainer(),
		MClassHandler:     InitMClassContainer(),
		MBadgeSettingsHandler: InitMBadgeSettingsContainer(),
		MCourseHandler:    InitMCourseContainer(),
	}
}

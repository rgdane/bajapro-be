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
	MLessonHandler    *handlers.MLessonHandler
	MSubLessonHandler *handlers.MSubLessonHandler
	MMaterialHandler  *handlers.MMaterialHandler
	TStudentCourseHandler *handlers.TStudentCourseHandler
	TStudentProgressHandler *handlers.TStudentProgressHandler
	TCodeQuestionHandler *handlers.TCodeQuestionHandler
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
		MLessonHandler:    InitMLessonContainer(),
		MSubLessonHandler: InitMSubLessonContainer(),
		MMaterialHandler:  InitMMaterialContainer(),
		TStudentCourseHandler: InitTStudentCourseContainer(),
		TStudentProgressHandler: InitTStudentProgressContainer(),
		TCodeQuestionHandler: InitTCodeQuestionContainer(),
	}
}

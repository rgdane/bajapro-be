package container

import (
	"jk-api/api/http/controllers/v1/handlers"
)

type AppContainer struct {
	AuthHandler             *handlers.AuthHandler
	MSchoolHandler          *handlers.MSchoolHandler
	MLevelHandler           *handlers.MLevelHandler
	UserHandler             *handlers.UserHandler
	MClassHandler           *handlers.MClassHandler
	MBadgeSettingsHandler   *handlers.MBadgeSettingsHandler
	MCourseHandler          *handlers.MCourseHandler
	MLessonHandler          *handlers.MLessonHandler
	MSubLessonHandler       *handlers.MSubLessonHandler
	MMaterialHandler        *handlers.MMaterialHandler
	TStudentCourseHandler   *handlers.TStudentCourseHandler
	TStudentProgressHandler *handlers.TStudentProgressHandler
	TCodeQuestionHandler    *handlers.TCodeQuestionHandler
	EssayQuestionHandler    *handlers.EssayQuestionHandler
	TCodeAnswerHandler      *handlers.TCodeAnswerHandler
	TWonderingScoreHandler  *handlers.TWonderingScoreHandler
	TEssayAnswerHandler     *handlers.TEssayAnswerHandler
}

func NewAppContainer() *AppContainer {
	return &AppContainer{
		AuthHandler:             InitAuthContainer(),
		MSchoolHandler:          InitMSchoolContainer(),
		MLevelHandler:           InitMLevelContainer(),
		UserHandler:             InitUserContainer(),
		MClassHandler:           InitMClassContainer(),
		MBadgeSettingsHandler:   InitMBadgeSettingsContainer(),
		MCourseHandler:          InitMCourseContainer(),
		MLessonHandler:          InitMLessonContainer(),
		MSubLessonHandler:       InitMSubLessonContainer(),
		MMaterialHandler:        InitMMaterialContainer(),
		TStudentCourseHandler:   InitTStudentCourseContainer(),
		TStudentProgressHandler: InitTStudentProgressContainer(),
		TCodeQuestionHandler:    InitTCodeQuestionContainer(),
		EssayQuestionHandler:    InitEssayQuestionContainer(),
		TCodeAnswerHandler:      InitTCodeAnswerContainer(),
		TWonderingScoreHandler:  InitTWonderingScoreContainer(),
		TEssayAnswerHandler:     InitTEssayAnswerContainer(),
	}
}

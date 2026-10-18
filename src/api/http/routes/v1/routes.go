package routes

import (
	"jk-api/internal/container"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App, c *container.AppContainer) {
	api := app.Group("/api/v1")

	AuthRoutes(api, c)
	MLevelRoutes(api, c)
	UserRoutes(api, c)
	MClassRoutes(api, c)
	MSchoolRoutes(api, c)
	MBadgeSettings(api, c)
	MCourses(api, c)
	TStudentCourseRoute(api, c)
	TStudentProgressRoutes(api, c)
	TCodeQuestionRoute(api, c)
	MLessonRoutes(api, c)
	MSubLessonRoutes(api, c)
	MMaterialRoutes(api, c)
	EssayQuestionRoute(api, c)
	TCodeAnswerRoute(api, c)
	TWonderingScoreRoute(api, c)
	TEssayAnswerRoute(api, c)
	CodeRoutes(api, c)
	UploadRoutes(api, c)
}

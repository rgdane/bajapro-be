

package routes

import (
	"jk-api/api/http/controllers/v1"
	"jk-api/api/http/middleware"
	"jk-api/internal/container"

	"github.com/gofiber/fiber/v2"
)

func TWonderingScoreRoute(router fiber.Router, c *container.AppContainer) {
	app := router.Group("t_wondering_score", middleware.JWTMiddleware())
	app.Post("/", controllers.CreateTWonderingScore(c))
	app.Get("/sub_lesson/:subLessonID", controllers.GetTWonderingScoresBySubLessonID(c))
}

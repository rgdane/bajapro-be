

package routes

import (
	"jk-api/api/http/controllers/v1"
	"jk-api/api/http/middleware"
	"jk-api/internal/container"

	"github.com/gofiber/fiber/v2"
)

func TEssayAnswerRoute(router fiber.Router, c *container.AppContainer) {
	app := router.Group("t_essay_answer", middleware.JWTMiddleware())
	app.Get("/essay_questions/:essayQuestionID", controllers.GetTEssayAnswersByEssayQuestionIDAndUserID(c))
	app.Post("/", controllers.CreateTEssayAnswer(c))
	
}

package routes

import (
	"jk-api/api/http/controllers/v1"
	"jk-api/api/http/middleware"
	"jk-api/internal/container"

	"github.com/gofiber/fiber/v2"
)

func CodeRoutes(router fiber.Router, c *container.AppContainer) {
	app := router.Group("code", middleware.JWTMiddleware())
	app.Post("/run", controllers.RunCode)
}

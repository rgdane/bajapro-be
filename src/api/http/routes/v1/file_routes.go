package routes

import (
	"jk-api/api/http/controllers/v1"
	"jk-api/api/http/middleware"

	"github.com/gofiber/fiber/v2"
)

func FileRoutes(router fiber.Router) {
	app := router.Group("/files")

	app.Post("/", controllers.UploadFiles(), middleware.JWTMiddleware())
	app.Get("/:name", controllers.GetFileByName())
}

package routes

import (
	"jk-api/api/http/controllers/v1"
	"jk-api/api/http/middleware"
	"jk-api/internal/container"

	"github.com/gofiber/fiber/v2"
)

func UploadRoutes(router fiber.Router, c *container.AppContainer) {
	router.Post(
		"upload",
		middleware.JWTMiddleware(),
		middleware.RequireRole("admin"),
		controllers.UploadFile,
	)
}

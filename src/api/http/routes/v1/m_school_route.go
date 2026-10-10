package routes

import (
	"jk-api/api/http/controllers/v1"
	"jk-api/api/http/middleware"
	"jk-api/internal/container"

	"github.com/gofiber/fiber/v2"
)

func MSchoolRoutes(router fiber.Router, c *container.AppContainer) {
	app := router.Group("school", middleware.JWTMiddleware())
	app.Get("/", controllers.GetMSchools(c))
	app.Post("/", middleware.RequireRole("admin"), controllers.CreateMSchools(c))
	app.Get("/:id", controllers.GetMSchoolByID(c))
	app.Put("/:id", middleware.RequireRole("admin"), controllers.UpdateMSchools(c))
	app.Delete("/:id", middleware.RequireRole("admin"), controllers.DeleteMSchools(c))
}

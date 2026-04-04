// m_user_routes.go

package routes

import (
	"jk-api/api/http/controllers/v1"
	"jk-api/api/http/middleware"
	"jk-api/internal/container"

	"github.com/gofiber/fiber/v2"
)

func MUserRoutes(router fiber.Router, c *container.AppContainer) {
	app := router.Group("m-users", middleware.JWTMiddleware())

	app.Get("/", controllers.GetMUsers(c))
	app.Get("/:id", controllers.GetMUserByID(c))
	app.Post("/bulk-create", controllers.BulkCreateMUsers(c))
	app.Put("/bulk-update", controllers.BulkUpdateMUsers(c))
	app.Delete("/bulk-delete", controllers.BulkDeleteMUsers(c))
	app.Post("/", controllers.CreateMUsers(c))
	app.Put("/:id", controllers.UpdateMUsers(c))
	app.Delete("/:id", controllers.DeleteMUsers(c))
}

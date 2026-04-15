package routes

import (
	"jk-api/internal/container"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App, c *container.AppContainer) {
	api := app.Group("/api/v1")

	AuthRoutes(api, c)
	MLevelRoutes(api, c)
	PermissionRoutes(api, c)
	RoleRoutes(api, c)
	UserRoutes(api, c)
	MClassRoute(api, c)
	MBadgeSettings(api, c)
	MCourse(api, c)
}

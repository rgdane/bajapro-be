package routes

import (
	"jk-api/internal/container"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App, c *container.AppContainer) {
	api := app.Group("/api/v1")

	AuthRoutes(api, c)
	LevelRoutes(api, c)
	PermissionRoutes(api, c)
	PositionRoutes(api, c)
	RoleRoutes(api, c)
	TitleRoutes(api, c)
	UserRoutes(api, c)
	DivisionRoutes(api, c)
	MClassRoute(api, c)
}

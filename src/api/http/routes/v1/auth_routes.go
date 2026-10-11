package routes

import (
	"jk-api/api/http/controllers/v1"
	"jk-api/api/http/middleware"
	"jk-api/internal/container"

	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(router fiber.Router, c *container.AppContainer) {
	app := router.Group("auth")

	app.Get("/profile",
		middleware.JWTMiddleware(),
		controllers.GetProfile(c),
	)
	app.Post("/login", controllers.Login(c))
	app.Post("/register", controllers.Register(c))
	// Daftar sekolah untuk dropdown registrasi guru (publik, tanpa JWT).
	app.Get("/schools", controllers.GetMSchools(c))
	app.Post("/refresh", controllers.RefreshToken(c))
	app.Post("/logout", middleware.JWTMiddleware(), controllers.Logout(c))

	// Admin teacher registration approval routes
	admin := router.Group("admin/approval-pengajar", middleware.JWTMiddleware(), middleware.RequireRole("admin"))
	admin.Get("/", controllers.GetPendingTeacherRequests(c))
	admin.Post("/:id", controllers.ProcessTeacherRequest(c))
}

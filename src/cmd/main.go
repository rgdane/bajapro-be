package main

import (
	"os"

	"jk-api/cmd/bootstrap"

	"github.com/gofiber/fiber/v2"
)

var log = bootstrap.Log

func main() {
	bootstrap.Setup()

	services := bootstrap.Services
	app := bootstrap.NewFiber()

	bootstrap.InitFiber(app, services)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "🚀 Welcome to JK API!",
			"status":  "healthy",
			"version": "1.0.0",
		})
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Infof("🚀 Server running at http://localhost:%s", port)
	app.Listen(":" + port)
}

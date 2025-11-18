package bootstrap

import (
	httpRoutes "jk-api/api/http/routes/v1"
	"jk-api/internal/container"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/cors"
	fiberLogger "github.com/gofiber/fiber/v2/middleware/logger"
)

func NewFiber() *fiber.App {
	return fiber.New(fiber.Config{
		AppName:         "JalanKerja API",
		ServerHeader:    "JK API Server with WebSocket Support",
		ReadBufferSize:  16 * 1024,
		WriteBufferSize: 16 * 1024,
		BodyLimit:       10 * 1024 * 1024, // 10 MB
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}

			if c.Get("Upgrade") == "websocket" {
				log.Errorf("WebSocket Error [%s %s]: %v", c.Method(), c.Path(), err)
			} else {
				log.Errorf("HTTP Error [%s %s]: %v", c.Method(), c.Path(), err)
			}

			return c.Status(code).JSON(fiber.Map{
				"error":   true,
				"message": err.Error(),
			})
		},
	})
}
func InitFiber(app *fiber.App, services *container.AppContainer) {
	app.Use(fiberLogger.New(fiberLogger.Config{
		Format:     "[${time}] [${status}] ${method} ${path} - ${latency}",
		TimeFormat: "2006-01-02 15:04:05",
		TimeZone:   "Asia/Jakarta",
	}))

	// Use Fiber built-in CORS middleware
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowHeaders:     "Origin,Content-Type,Accept,Authorization",
		AllowCredentials: false,
		MaxAge:           86400,
	}))

	httpRoutes.Setup(app, services)
}

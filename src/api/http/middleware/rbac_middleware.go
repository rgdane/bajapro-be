package middleware

import (
	"github.com/gofiber/fiber/v2"
)

func RequireRole(roles ...string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userRole, ok := c.Locals("role").(string)
		if !ok || userRole == "" {
			return fiber.ErrForbidden
		}

		for _, r := range roles {
			if userRole == r {
				return c.Next()
			}
		}

		return fiber.ErrForbidden
	}
}

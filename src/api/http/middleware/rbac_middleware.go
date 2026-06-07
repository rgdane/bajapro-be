package middleware

import (
	"github.com/gofiber/fiber/v2"
)

func RequireRole(roles ...string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userRoles := c.Locals("roles").([]string)

		for _, ur := range userRoles {
			for _, r := range roles {
				if ur == r {
					return c.Next()
				}
			}
		}

		return fiber.ErrForbidden
	}
}

func RequirePermission(perms ...string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userPerms := c.Locals("permissions").([]string)

		for _, up := range userPerms {
			for _, p := range perms {
				if up == p {
					return c.Next()
				}
			}
		}

		return fiber.ErrForbidden
	}
}
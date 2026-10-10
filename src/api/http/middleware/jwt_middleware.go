package middleware

import (
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func JWTMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var tokenStr string

		// 1. Prioritas ambil dari Cookie
		tokenStr = c.Cookies("access_token")

		// 2. Fallback ke Authorization Header (untuk Postman/API)
		if tokenStr == "" {
			authHeader := c.Get("Authorization")

			if strings.HasPrefix(authHeader, "Bearer ") {
				tokenStr = strings.TrimPrefix(authHeader, "Bearer ")
			}
		}

		if tokenStr == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"success": false,
				"error":   "missing access token",
			})
		}

		token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil || !token.Valid {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"success": false,
				"error":   "invalid token",
			})
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"success": false,
				"error":   "invalid claims",
			})
		}

		if userID, ok := claims["user_id"].(float64); ok {
			c.Locals("user_id", int64(userID))
		}

		c.Locals("name", claims["name"])
		c.Locals("role", claims["role"])

		return c.Next()
	}
}
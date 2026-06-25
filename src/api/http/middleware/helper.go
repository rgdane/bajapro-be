package middleware

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func extractToken(c *fiber.Ctx) (string, error) {
	authHeader := c.Get("Authorization")

	if authHeader == "" {
		return "", fmt.Errorf("missing authorization header")
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return "", fmt.Errorf("invalid authorization format")
	}

	return parts[1], nil
}

func toStringSlice(data interface{}) []string {
	var result []string

	if arr, ok := data.([]interface{}); ok {
		for _, v := range arr {
			if str, ok := v.(string); ok {
				result = append(result, str)
			}
		}
	}

	return result
}
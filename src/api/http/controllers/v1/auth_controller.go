package controllers

import (
	"errors"
	"jk-api/api/http/controllers/v1/dto"
	"jk-api/api/http/controllers/v1/handlers"
	"jk-api/api/http/presenters"
	"jk-api/internal/container"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type AuthController struct {
	Handler *handlers.AuthHandler
}

func NewAuthController(h *handlers.AuthHandler) *AuthController {
	return &AuthController{Handler: h}
}

func GetProfile(cn *container.AppContainer) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return presenters.ErrorResponse(c, fiber.StatusUnauthorized, errors.New("missing authorization header"))
		}

		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			return presenters.ErrorResponse(c, fiber.StatusUnauthorized, errors.New("invalid authorization header format"))
		}

		tokenString := tokenParts[1]
		data, err := cn.AuthHandler.GetProfileHandler(tokenString)
		if err != nil {
			return presenters.ErrorResponse(c, fiber.StatusInternalServerError, err)
		}
		return presenters.SuccessResponse(c, data)
	}
}

func Login(cn *container.AppContainer) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var input dto.LoginRequest
		if err := c.BodyParser(&input); err != nil {
			return presenters.ErrorResponseWithMessage(c, fiber.StatusBadRequest, "Invalid request")
		}

		dto, _, err := cn.AuthHandler.Login(&input)
		if err != nil {
			return presenters.ErrorResponse(c, fiber.StatusInternalServerError, err)
		}

		return presenters.SuccessResponse(c, dto)
	}
}

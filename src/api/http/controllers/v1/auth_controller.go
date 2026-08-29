package controllers

import (
	"errors"
	"jk-api/api/http/controllers/v1/dto"
	"jk-api/api/http/controllers/v1/handlers"
	"jk-api/api/http/presenters"
	"jk-api/internal/container"

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
		token := c.Cookies("access_token")

		if token == "" {
			return presenters.ErrorResponse(c,
				fiber.StatusUnauthorized,
				errors.New("missing access token"))
		}

		data, err := cn.AuthHandler.GetProfileHandler(token)
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

		result, _, err := cn.AuthHandler.Login(&input)
		if err != nil {
			return presenters.ErrorResponse(c, fiber.StatusInternalServerError, err)
		}

		// Access Token Cookie
		c.Cookie(&fiber.Cookie{
			Name:     "access_token",
			Value:    result.AccessToken,
			HTTPOnly: true,
			Secure:   false, // ubah menjadi true jika sudah HTTPS
			SameSite: "Lax",
			Path:     "/",
			MaxAge:   60 * 60, // 1 jam
		})

		// Refresh Token Cookie
		c.Cookie(&fiber.Cookie{
			Name:     "refresh_token",
			Value:    result.RefreshToken,
			HTTPOnly: true,
			Secure:   false,
			SameSite: "Lax",
			Path:     "/",
			MaxAge:   60 * 60 * 24 * 7, // 7 hari
		})

		return presenters.SuccessResponse(c, result)
	}
}

func Logout(cn *container.AppContainer) fiber.Handler {
	return func(c *fiber.Ctx) error {
		refreshToken := c.Cookies("refresh_token")
		if refreshToken != "" {
			_ = cn.AuthHandler.Logout(refreshToken)
		}

		c.Cookie(&fiber.Cookie{
			Name:     "access_token",
			Value:    "",
			MaxAge:   -1,
			Path:     "/",
			HTTPOnly: true,
			SameSite: "Lax",
		})

		c.Cookie(&fiber.Cookie{
			Name:     "refresh_token",
			Value:    "",
			MaxAge:   -1,
			Path:     "/",
			HTTPOnly: true,
			SameSite: "Lax",
		})

		return presenters.SuccessResponse(c, "Logged out successfully")
	}
}

func RefreshToken(cn *container.AppContainer) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var input struct {
			RefreshToken string `json:"refresh_token"`
		}
		if err := c.BodyParser(&input); err != nil {
			return presenters.ErrorResponseWithMessage(c, fiber.StatusBadRequest, "Invalid request")
		}

		newToken, err := cn.AuthHandler.RefreshToken(input.RefreshToken)
		if err != nil {
			return presenters.ErrorResponse(c, fiber.StatusInternalServerError, err)
		}
		return presenters.SuccessResponse(c, map[string]string{"token": newToken})
	}
}

func Register(cn *container.AppContainer) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var input dto.RegisterRequest
		if err := c.BodyParser(&input); err != nil {
			return presenters.ErrorResponseWithMessage(c, fiber.StatusBadRequest, "Invalid request")
		}
		dto, _, err := cn.AuthHandler.Register(&input)
		if err != nil {
			return presenters.ErrorResponse(c, fiber.StatusInternalServerError, err)
		}

		return presenters.SuccessResponse(c, dto)
	}
}

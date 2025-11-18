package presenters

import "github.com/gofiber/fiber/v2"

func SuccessResponse(c *fiber.Ctx, data any, total ...int64) error {
	resp := fiber.Map{
		"success": true,
		"data":    data,
	}

	if len(total) > 0 && total[0] > 0 {
		resp["total"] = total[0]
	}

	return c.Status(fiber.StatusOK).JSON(resp)
}

func SuccessCreatedResponse(c *fiber.Ctx, data any) error {
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"data":    data,
	})
}

func ErrorResponse(c *fiber.Ctx, status int, err error) error {
	return c.Status(status).JSON(fiber.Map{
		"success": false,
		"error":   err.Error(),
	})
}

func SuccessResponseWithMessage(c *fiber.Ctx, message string, data any) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": message,
		"data":    data,
	})
}

func ErrorResponseWithMessage(c *fiber.Ctx, status int, message string) error {
	return c.Status(status).JSON(fiber.Map{
		"success": false,
		"error":   message,
	})
}

// func SuccessLogin(c *fiber.Ctx, data any, token string) error {
// 	return c.Status(fiber.StatusOK).JSON(fiber.Map{
// 		"success": true,
// 		"data":   data,
// 		"token": token,
// 	})
// }

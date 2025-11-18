package controllers

import (
	"jk-api/api/http/controllers/v1/handlers"
	"jk-api/api/http/presenters"
	"net/url"

	"github.com/gofiber/fiber/v2"
)

func UploadFiles() fiber.Handler {
	return func(c *fiber.Ctx) error {
		form, err := c.MultipartForm()
		if err != nil {
			return presenters.ErrorResponseWithMessage(c, fiber.StatusBadRequest, err.Error())
		}

		files := form.File["file"]
		if len(files) == 0 {
			return presenters.ErrorResponseWithMessage(c, fiber.StatusBadRequest, "No files provided")
		}

		for _, file := range files {
			err = handlers.UploadFileHandler(file, c)
			if err != nil {
				return presenters.ErrorResponse(c, fiber.StatusInternalServerError, err)
			}
		}

		return nil
	}
}

func GetFileByName() fiber.Handler {
	return func(c *fiber.Ctx) error {
		rawName := c.Params("name")
		name, err := url.QueryUnescape(rawName)
		if err != nil {
			return presenters.ErrorResponse(c, fiber.StatusBadRequest, err)
		}

		err = handlers.GetFileByNameHandler(c, name)
		if err != nil {
			return presenters.ErrorResponse(c, fiber.StatusInternalServerError, err)
		}
		return nil
	}
}

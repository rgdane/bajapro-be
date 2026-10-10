package controllers

import (
	"jk-api/api/http/controllers/v1/dto"
	"jk-api/api/http/presenters"
	"jk-api/internal/container"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetMSchools(cn *container.AppContainer) fiber.Handler {
	return func(c *fiber.Ctx) error {
		filter := dto.MSchoolFilterDto{
			Preload: c.Query("preload", "false") == "true",
			Name:    c.Query("name"),
		}

		data, err := cn.MSchoolHandler.GetAllMSchoolsHandler(filter)
		if err != nil {
			return presenters.ErrorResponse(c, fiber.StatusInternalServerError, err)
		}
		return presenters.SuccessResponse(c, data)
	}
}

func GetMSchoolByID(cn *container.AppContainer) fiber.Handler {
	return func(c *fiber.Ctx) error {
		filter := dto.MSchoolFilterDto{
			Preload: c.Query("preload", "false") == "true",
		}

		id, err := strconv.ParseInt(c.Params("id"), 10, 64)
		if err != nil {
			return presenters.ErrorResponseWithMessage(c, fiber.StatusBadRequest, "Invalid ID")
		}

		data, err := cn.MSchoolHandler.GetMSchoolByIDHandler(id, filter)
		if err != nil {
			return presenters.ErrorResponse(c, fiber.StatusInternalServerError, err)
		}
		return presenters.SuccessResponse(c, data)
	}
}

func CreateMSchools(cn *container.AppContainer) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var input dto.CreateMSchoolDto
		if err := c.BodyParser(&input); err != nil {
			return presenters.ErrorResponseWithMessage(c, fiber.StatusBadRequest, "Invalid request")
		}

		result, err := cn.MSchoolHandler.CreateMSchoolHandler(&input)
		if err != nil {
			return presenters.ErrorResponse(c, fiber.StatusInternalServerError, err)
		}
		return presenters.SuccessResponse(c, result)
	}
}

func UpdateMSchools(cn *container.AppContainer) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := strconv.ParseInt(c.Params("id"), 10, 64)
		if err != nil {
			return presenters.ErrorResponseWithMessage(c, fiber.StatusBadRequest, "Invalid ID")
		}

		var input dto.UpdateMSchoolDto
		if err := c.BodyParser(&input); err != nil {
			return presenters.ErrorResponseWithMessage(c, fiber.StatusBadRequest, "Invalid input")
		}

		updated, err := cn.MSchoolHandler.UpdateMSchoolHandler(id, &input)
		if err != nil {
			return presenters.ErrorResponse(c, fiber.StatusInternalServerError, err)
		}
		return presenters.SuccessResponse(c, updated)
	}
}

func DeleteMSchools(cn *container.AppContainer) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := strconv.ParseInt(c.Params("id"), 10, 64)
		if err != nil {
			return presenters.ErrorResponseWithMessage(c, fiber.StatusBadRequest, "Invalid ID")
		}

		if err := cn.MSchoolHandler.DeleteMSchoolHandler(id); err != nil {
			return presenters.ErrorResponse(c, fiber.StatusInternalServerError, err)
		}
		return presenters.SuccessResponseWithMessage(c, "MSchool deleted successfully", nil)
	}
}

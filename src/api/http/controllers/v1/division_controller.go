package controllers

import (
	"fmt"
	"jk-api/api/http/controllers/v1/dto"
	"jk-api/api/http/presenters"
	"jk-api/internal/container"
	"jk-api/internal/helper"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetDivisions(cn *container.AppContainer) fiber.Handler {
	return func(c *fiber.Ctx) error {
		departmentId, _ := helper.ParseQueryInt64(c, "department_id")
		sopId, _ := helper.ParseQueryInt64(c, "sop_id")
		sort := c.Query("sort")
		order := c.Query("order")
		cursor, _ := helper.ParseQueryInt64(c, "cursor")
		limit, _ := helper.ParseQueryInt64(c, "limit")
		name := c.Query("name")

		filter := dto.DivisionFilterDto{
			DepartmentID: departmentId,
			SopId:        sopId,
			Preload:      c.Query("preload", "false") == "true",
			Sort:         sort,
			Order:        order,
			Cursor:       cursor,
			Limit:        limit,
			Name:         name,
			ShowDeleted:  c.Query("show_deleted", "false") == "true",
		}

		data, total, err := cn.DivisionHandler.GetAllDivisionsHandler(filter)
		if err != nil {
			return presenters.ErrorResponse(c, fiber.StatusInternalServerError, err)
		}
		return presenters.SuccessResponse(c, data, total)
	}
}

func GetDivisionByID(cn *container.AppContainer) fiber.Handler {
	return func(c *fiber.Ctx) error {
		filter := dto.DivisionFilterDto{
			Preload: c.Query("preload", "false") == "true",
		}
		id, err := strconv.ParseInt(c.Params("id"), 10, 64)
		if err != nil {
			return presenters.ErrorResponseWithMessage(c, fiber.StatusBadRequest, "Invalid ID")
		}

		data, err := cn.DivisionHandler.GetDivisionByIDHandler(id, filter)
		if err != nil {
			return presenters.ErrorResponse(c, fiber.StatusInternalServerError, err)
		}
		return presenters.SuccessResponse(c, data)
	}
}

func CreateDivisions(cn *container.AppContainer) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var input dto.CreateDivisionDto
		if err := c.BodyParser(&input); err != nil {
			return presenters.ErrorResponseWithMessage(c, fiber.StatusBadRequest, "Invalid request")
		}

		result, err := cn.DivisionHandler.CreateDivisionHandler(&input)
		if err != nil {
			return presenters.ErrorResponse(c, fiber.StatusInternalServerError, err)
		}
		return presenters.SuccessResponse(c, result)
	}
}

func UpdateDivisions(cn *container.AppContainer) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := strconv.ParseInt(c.Params("id"), 10, 64)
		if err != nil {
			return presenters.ErrorResponseWithMessage(c, fiber.StatusBadRequest, "Invalid ID")
		}

		var input dto.UpdateDivisionDto
		if err := c.BodyParser(&input); err != nil {
			return presenters.ErrorResponseWithMessage(c, fiber.StatusBadRequest, "Invalid input")
		}

		updated, err := cn.DivisionHandler.UpdateDivisionHandler(id, &input)
		if err != nil {
			return presenters.ErrorResponse(c, fiber.StatusInternalServerError, err)
		}
		return presenters.SuccessResponse(c, updated)
	}
}

func DeleteDivisions(cn *container.AppContainer) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := strconv.ParseInt(c.Params("id"), 10, 64)
		if err != nil {
			return presenters.ErrorResponseWithMessage(c, fiber.StatusBadRequest, "Invalid ID")
		}

		if err := cn.DivisionHandler.DeleteDivisionHandler(id); err != nil {
			return presenters.ErrorResponse(c, fiber.StatusInternalServerError, err)
		}
		return presenters.SuccessResponseWithMessage(c, "Division deleted successfully", nil)
	}
}

func BulkCreateDivisions(cn *container.AppContainer) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var input dto.BulkCreateDivisionDto
		if err := c.BodyParser(&input); err != nil {
			return presenters.ErrorResponseWithMessage(c, fiber.StatusBadRequest, "Invalid request body for bulk create")
		}
		if len(input.Data) == 0 {
			return presenters.ErrorResponseWithMessage(c, fiber.StatusBadRequest, "No data provided")
		}

		createdDivisions, err := cn.DivisionHandler.BulkCreateHandler(&input)
		if err != nil {
			return presenters.ErrorResponse(c, fiber.StatusInternalServerError, err)
		}
		return presenters.SuccessResponse(c, createdDivisions)
	}
}

func BulkUpdateDivisions(cn *container.AppContainer) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var input dto.BulkUpdateDivisionDto
		if err := c.BodyParser(&input); err != nil {
			return presenters.ErrorResponseWithMessage(c, fiber.StatusBadRequest, "Invalid request body for bulk update")
		}
		if len(input.IDs) == 0 {
			return presenters.ErrorResponseWithMessage(c, fiber.StatusBadRequest, "No Division IDs provided")
		}
		if input.Data == nil {
			return presenters.ErrorResponseWithMessage(c, fiber.StatusBadRequest, "No update data provided")
		}

		updatedDivisions, err := cn.DivisionHandler.BulkUpdateHandler(&input)
		if err != nil {
			return presenters.ErrorResponse(c, fiber.StatusInternalServerError, err)
		}
		return presenters.SuccessResponse(c, updatedDivisions)
	}
}

func BulkDeleteDivisions(cn *container.AppContainer) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var input dto.BulkDeleteDivisionDto
		if err := c.BodyParser(&input); err != nil {
			return presenters.ErrorResponseWithMessage(c, fiber.StatusBadRequest, "Invalid request body for bulk delete")
		}

		if len(input.IDs) == 0 {
			return presenters.ErrorResponseWithMessage(c, fiber.StatusBadRequest, "No Division IDs provided")
		}

		err := cn.DivisionHandler.BulkDeleteHandler(&input)
		if err != nil {
			return presenters.ErrorResponse(c, fiber.StatusInternalServerError, err)
		}

		return presenters.SuccessResponseWithMessage(c, fmt.Sprintf("Successfully deleted %d Divisions", len(input.IDs)), nil)
	}
}

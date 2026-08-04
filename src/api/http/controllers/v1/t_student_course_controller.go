package controllers

import (
	"errors"
	"jk-api/api/http/controllers/v1/dto"
	"jk-api/api/http/presenters"
	"jk-api/internal/container"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetMyCourse(cn *container.AppContainer) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userID := c.Locals("user_id").(int64)
		filter := dto.TStudentCourseFilterDto{
			Preload: c.Query("preload", "false") == "true",
		}

		data, err := cn.TStudentCourseHandler.GetMyCoursesHandler(userID, filter)
		if err != nil {
			return presenters.ErrorResponse(c, fiber.StatusInternalServerError, err)
		}
		return presenters.SuccessResponse(c, data)
	}
}

func GetTStudentCourseByID(cn *container.AppContainer) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userID := c.Locals("user_id").(int64)
		id, err := strconv.ParseInt(c.Params("id"), 10, 64)
		if err != nil {
			return presenters.ErrorResponseWithMessage(c, fiber.StatusBadRequest, "Invalid course ID")
		}

		filter := dto.TStudentCourseFilterDto{
			Preload: c.Query("preload", "false") == "true",
		}

		data, err := cn.TStudentCourseHandler.GetTStudentCourseByIDHandler(filter, userID, id)
		if err != nil {
			return presenters.ErrorResponse(c, fiber.StatusInternalServerError, err)
		}
		return presenters.SuccessResponse(c, data)
	}
}

func EnrollCourse(cn *container.AppContainer) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userID := c.Locals("user_id").(int64)

		courseID, err := strconv.ParseInt(c.Params("id"), 10, 64)
		if err != nil {
			return presenters.ErrorResponseWithMessage(
				c,
				fiber.StatusBadRequest,
				"Invalid course id",
			)
		}

		result, err := cn.TStudentCourseHandler.
			EnrollTStudentCourseHandler(userID, courseID)

		if err != nil {
			return presenters.ErrorResponse(
				c,
				fiber.StatusInternalServerError,
				err,
			)
		}

		return presenters.SuccessResponse(c, result)
	}
}

func GetLeaderboard(
	cn *container.AppContainer,
) fiber.Handler {
	return func(c *fiber.Ctx) error {
		courseID, err := strconv.ParseInt(
			c.Query("course_id"),
			10,
			64,
		)
		if err != nil || courseID <= 0 {
			return presenters.ErrorResponse(
				c,
				fiber.StatusBadRequest,
				errors.New("course_id tidak valid"),
			)
		}

		var classID *int64

		if rawClassID := c.Query("class_id"); rawClassID != "" {
			parsedClassID, parseErr := strconv.ParseInt(
				rawClassID,
				10,
				64,
			)
			if parseErr != nil || parsedClassID <= 0 {
				return presenters.ErrorResponse(
					c,
					fiber.StatusBadRequest,
					errors.New("class_id tidak valid"),
				)
			}

			classID = &parsedClassID
		}

		result, err :=
			cn.TStudentCourseHandler.GetLeaderboardHandler(
				courseID,
				classID,
			)

		if err != nil {
			return presenters.ErrorResponse(
				c,
				fiber.StatusInternalServerError,
				err,
			)
		}

		return presenters.SuccessResponse(c, result)
	}
}
func GetDashboard(cn *container.AppContainer) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userID := c.Locals("user_id").(int64)
		
		result, err := cn.TStudentCourseHandler.GetDashboardHandler(userID)

		if err != nil {
			return presenters.ErrorResponse(
				c,
				fiber.StatusInternalServerError,
				err,
			)
		}

		return presenters.SuccessResponse(c, result)
	}
}

func GetCourseProgress(cn *container.AppContainer) fiber.Handler {
	return func(c *fiber.Ctx) error {
		studentID := c.Locals("user_id").(int64)

		courseID, err := strconv.ParseInt(
			c.Query("courseId"),
			10,
			64,
		)
		if err != nil || courseID <= 0 {
			return presenters.ErrorResponseWithMessage(
				c,
				fiber.StatusBadRequest,
				"Invalid course id",
			)
		}

		result, err := cn.TStudentCourseHandler.
			GetCourseProgressHandler(studentID, courseID)

		if err != nil {
			return presenters.ErrorResponse(
				c,
				fiber.StatusInternalServerError,
				err,
			)
		}

		if result == nil {
			return c.Status(fiber.StatusOK).JSON(nil)
		}

		return presenters.SuccessResponse(c, result)
	}
}


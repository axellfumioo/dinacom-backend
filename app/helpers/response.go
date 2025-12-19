package helpers

import (
	"backend-dinakom/app/dto/response"
	"math"
	"github.com/gofiber/fiber/v2"
)

func SuccessResponse(c *fiber.Ctx, message string, data interface{}) error {
	return c.Status(fiber.StatusOK).JSON(response.BaseResponse{
		Success: true,
		Message: message,
		Data:    data,
	})
}

func CreatedResponse(c *fiber.Ctx, message string, data interface{}) error {
	return c.Status(fiber.StatusCreated).JSON(response.BaseResponse{
		Success: true,
		Message: message,
		Data:    data,
	})
}

func ErrorResponse(c *fiber.Ctx, statusCode int, message string, err interface{}) error {
	return c.Status(statusCode).JSON(response.BaseResponse{
		Success: false,
		Message: message,
		Error:   err,
	})
}

func ValidationErrorResponse(c *fiber.Ctx, err interface{}) error {
	return c.Status(fiber.StatusBadRequest).JSON(response.BaseResponse{
		Success: false,
		Message: "Validation error",
		Error:   err,
	})
}

func PaginationResponse(c *fiber.Ctx, message string, data interface{}, page, pageSize int, totalRows int64) error {
	totalPages := int(math.Ceil(float64(totalRows) / float64(pageSize)))

	return c.JSON(response.PaginationResponse{
		Success: true,
		Message: message,
		Data:    data,
		Pagination: response.Pagination{
			Page:       page,
			PageSize:   pageSize,
			TotalRows:  totalRows,
			TotalPages: totalPages,
		},
	})
}

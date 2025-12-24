package controllers

import (
	"backend-dinakom/app/helpers"
	"backend-dinakom/app/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type FoodScanResultController interface {
	GetAllResults(c *fiber.Ctx) error
	GetAllUserResults(c *fiber.Ctx) error
	GetResultByID(c *fiber.Ctx) error
}

type foodScanResultController struct {
	foodScanResultService services.FoodScanResultService
}

func NewFoodScanResultController(foodScanResultService services.FoodScanResultService) FoodScanResultController {
	return &foodScanResultController{foodScanResultService: foodScanResultService}
}

func (ctrl *foodScanResultController) GetAllResults(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	pageSize, _ := strconv.Atoi(c.Query("page_size", "10"))

	// Validate pagination
	if page < 1 {
		page = 1
	}
	// Validate pagination
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}

	data, totalRows, err := ctrl.foodScanResultService.GetAllResults(page, pageSize)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return helpers.PaginationResponse(c, "foodscan results retrieved successfully", data, page, pageSize, totalRows)
}

func (ctrl *foodScanResultController) GetAllUserResults(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(string)
	page, _ := strconv.Atoi(c.Query("page", "1"))
	pageSize, _ := strconv.Atoi(c.Query("page_size", "10"))
	// Validate pagination
	if page < 1 {
		page = 1
	}
	// Validate pagination
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}

	data, totalRows, err := ctrl.foodScanResultService.GetAllUserResults(userID, page, pageSize)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return helpers.PaginationResponse(c, "foodscan results retrieved successfully", data, page, pageSize, totalRows)
}

func (ctrl *foodScanResultController) GetResultByID(c *fiber.Ctx) error {
	id := c.Params("id")

	result, err := ctrl.foodScanResultService.GetResultByID(id)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return helpers.SuccessResponse(c, "result retrieved successfully", result)
}

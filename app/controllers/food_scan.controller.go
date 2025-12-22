package controllers

import (
	"backend-dinakom/app/dto/request"
	"backend-dinakom/app/helpers"
	"backend-dinakom/app/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type FoodScanController interface {
	GetAllFoodScans(c *fiber.Ctx) error
	GetUserFoodScans(c *fiber.Ctx) error
	ScanFood(c *fiber.Ctx) error
}

type foodScanController struct {
	foodScanService services.FoodScanService
}

func NewFoodScanController(foodScanService services.FoodScanService) FoodScanController {
	return &foodScanController{foodScanService: foodScanService}
}

func (ctrl *foodScanController) GetAllFoodScans(c *fiber.Ctx) error {
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

	foodScans, totalRows, err := ctrl.foodScanService.GetAllFoodScans(page, pageSize)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return helpers.PaginationResponse(c, "foodscans retrieved successfully", foodScans, page, pageSize, totalRows)
}

func (ctrl *foodScanController) GetUserFoodScans(c *fiber.Ctx) error {
	userId := c.Locals("user_id").(string)

	foodScans, err := ctrl.foodScanService.GetUserFoodScans(userId)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return helpers.SuccessResponse(c, "foodscans retrieved successfully", foodScans)
}

func (ctrl *foodScanController) ScanFood(c *fiber.Ctx) error {
	userId := c.Locals("user_id").(string)

	image, err := c.FormFile("image")
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "This field must be image file")
	}

	foodScan, err := ctrl.foodScanService.ScanFood(userId, request.ScanFoodRequest{Image: *image})
	return helpers.CreatedResponse(c, "food_scan created successfully", foodScan)
}

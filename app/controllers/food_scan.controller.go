package controllers

import (
	"backend-dinakom/app/dto/request"
	"backend-dinakom/app/helpers"
	"backend-dinakom/app/services"

	"github.com/gofiber/fiber/v2"
)

type FoodScanController interface {
	GetUserFoodScans(c *fiber.Ctx) error
	ScanFood(c *fiber.Ctx) error
}

type foodScanController struct {
	foodScanService services.FoodScanService
}

func NewFoodScanController(foodScanService services.FoodScanService) FoodScanController {
	return &foodScanController{foodScanService: foodScanService}
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

package controllers

import (
	"backend-dinakom/app/helpers"
	"backend-dinakom/app/services"

	"github.com/gofiber/fiber/v2"
)

type FoodScanResultController interface {
	GetResultByID(c *fiber.Ctx) error
}

type foodScanResultController struct {
	foodScanResultService services.FoodScanResultService
}

func NewFoodScanResultController(foodScanResultService services.FoodScanResultService) FoodScanResultController {
	return &foodScanResultController{foodScanResultService: foodScanResultService}
}

func (ctrl *foodScanResultController) GetResultByID(c *fiber.Ctx) error {
	id := c.Params("id")

	result, err := ctrl.foodScanResultService.GetResultByID(id)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return helpers.SuccessResponse(c, "result retrieved successfully", result)
}

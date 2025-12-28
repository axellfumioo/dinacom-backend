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

// GetAllFoodScanResults(Admin) godoc
// @Summary GetAllFoodScanResults
// @Description Endpoint to get all FoodScans results data
// @Tags Results
// @Produce  json
// @Param page query int false "Page number" default(1)
// @Param page_size query int false "Page size" default(10) minimum(1) maximum(100)
// @Success 200 {object} response.PaginationResponse
// @Failure 400 {object} response.ErrorResponse
// @Failure 404 {object} response.ErrorResponse
// @Router /results [get]
// @Security BearerAuth
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

// GetUserFoodScanResults godoc
// @Summary GetUserFoodScanResult
// @Description Endpoint to get all user foodscan results data
// @Tags Results
// @Produce  json
// @Param page query int false "Page number" default(1)
// @Param page_size query int false "Page size" default(10) minimum(1) maximum(100)
// @Success 200 {object} response.PaginationResponse
// @Failure 400 {object} response.ErrorResponse
// @Failure 404 {object} response.ErrorResponse
// @Router /results/user [get]
// @Security BearerAuth
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

// GetFoodScanResultByID godoc
// @Summary GetFoodScanResultByID
// @Description Endpoint to get all user foodscan results data by ID
// @Tags Results
// @Produce  json
// @Param id path string true "ID"
// @Success 200 {object} response.SuccessResponse
// @Failure 400 {object} response.ErrorResponse
// @Failure 404 {object} response.ErrorResponse
// @Router /results/{id}/get [get]
// @Security BearerAuth
func (ctrl *foodScanResultController) GetResultByID(c *fiber.Ctx) error {
	id := c.Params("id")

	result, err := ctrl.foodScanResultService.GetResultByID(id)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return helpers.SuccessResponse(c, "result retrieved successfully", result)
}

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
	GetFoodScanByID(c *fiber.Ctx) error
	ScanFood(c *fiber.Ctx) error
}

type foodScanController struct {
	foodScanService services.FoodScanService
}

func NewFoodScanController(foodScanService services.FoodScanService) FoodScanController {
	return &foodScanController{foodScanService: foodScanService}
}

// GetAllFoodScans godoc
// @Summary GetAllFoodScans
// @Description Endpoint to GetAllFoodScans data
// @Tags Foodscans
// @Produce  json
// @Param page query int false "Page number" default(1)
// @Param page_size query int false "Page size" default(10) minimum(1) maximum(100)
// @Success 200 {object} response.PaginationResponse
// @Failure 400 {object} response.ErrorResponse
// @Failure 404 {object} response.ErrorResponse
// @Router /foodscans [get]
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

// GetUserFoodScans godoc
// @Summary GetUserFoodScans
// @Description Endpoint to get all user FoodScans data
// @Tags Foodscans
// @Produce  json
// @Param page query int false "Page number" default(1)
// @Param page_size query int false "Page size" default(10) minimum(1) maximum(100)
// @Success 200 {object} response.PaginationResponse
// @Failure 400 {object} response.ErrorResponse
// @Failure 404 {object} response.ErrorResponse
// @Router /foodscans/user [get]
// @Security BearerAuth
func (ctrl *foodScanController) GetUserFoodScans(c *fiber.Ctx) error {
	userId := c.Locals("user_id").(string)

	foodScans, err := ctrl.foodScanService.GetUserFoodScans(userId)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return helpers.SuccessResponse(c, "foodscans retrieved successfully", foodScans)
}

func (ctrl *foodScanController) GetFoodScanByID(c *fiber.Ctx) error {
	UserID := c.Locals("user_id").(string)
	id := c.Params("id")
	
	foodScan, err := ctrl.foodScanService.GetFoodScanByID(UserID, id)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return helpers.SuccessResponse(c, "foodscan retrieved successfully", foodScan)
}

// Scanfood godoc
// @Summary ScanFood
// @Description Access this endpoint to scanfoods (cukup kirim file dengan nama body "image")
// @Tags Foodscans
// @Produce  json
// @Param image formData file true "upload image"
// @Success 201 {object} response.SuccessResponse
// @Failure 400 {object} response.ErrorResponse
// @Failure 404 {object} response.ErrorResponse
// @Router /foodscans/scan [post]
// @Security BearerAuth
func (ctrl *foodScanController) ScanFood(c *fiber.Ctx) error {
	userId := c.Locals("user_id").(string)

	image, err := c.FormFile("image")
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	foodScan, err := ctrl.foodScanService.ScanFood(userId, request.ScanFoodRequest{Image: *image})
	return helpers.CreatedResponse(c, "food_scan created successfully", foodScan)
}

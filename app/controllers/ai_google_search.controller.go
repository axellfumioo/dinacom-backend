package controllers

import (
	"backend-dinakom/app/dto/request"
	"backend-dinakom/app/helpers"
	"backend-dinakom/app/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type AIGoogleSearchController interface {
	GetAllGoogleSearchs(c *fiber.Ctx) error
	GetGoogleSearchByID(c *fiber.Ctx) error
	CreateGoogleSearch(c *fiber.Ctx) error
	UpdateGoogleSearch(c *fiber.Ctx) error
	DeleteGoogleSearch(c *fiber.Ctx) error
}

type aiGoogleSearchController struct {
	AIGoogleSearchService services.AIGoogleSearchService
}

func NewAIGoogleSearchController(AIGoogleSearchService services.AIGoogleSearchService) AIGoogleSearchController {
	return &aiGoogleSearchController{AIGoogleSearchService: AIGoogleSearchService}
}

func (ctrl *aiGoogleSearchController) GetAllGoogleSearchs(c *fiber.Ctx) error {
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

	data, totalRows, err := ctrl.AIGoogleSearchService.GetAllGoogleSearchs(page, pageSize)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return helpers.PaginationResponse(c, "AI google searchs retrieved successfully", data, page, pageSize, totalRows)
}

func (ctrl *aiGoogleSearchController) GetGoogleSearchByID(c *fiber.Ctx) error {
	ID := c.Params("id")
	data, err := ctrl.AIGoogleSearchService.GetGoogleSearchByID(ID)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return helpers.SuccessResponse(c, "AI google search retrieved successfully", data)
}

func (ctrl *aiGoogleSearchController) CreateGoogleSearch(c *fiber.Ctx) error {
	var req request.CreateGoogleSearchRequest

	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	data, err := ctrl.AIGoogleSearchService.CreateGoogleSearch(req)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return helpers.CreatedResponse(c, "AI google search created successfully", data)
}

func (ctrl *aiGoogleSearchController) UpdateGoogleSearch(c *fiber.Ctx) error {
	ID := c.Params("id")
	var req request.UpdateGoogleSearchRequest

	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	data, err := ctrl.AIGoogleSearchService.UpdateGoogleSearch(ID, req)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return helpers.SuccessResponse(c, "AI google search updated successfully", data)
}

func (ctrl *aiGoogleSearchController) DeleteGoogleSearch(c *fiber.Ctx) error {
	ID := c.Params("id")

	data, err := ctrl.AIGoogleSearchService.DeleteGoogleSearch(ID)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return helpers.SuccessResponse(c, "AI google search deleted successfully", data)
}

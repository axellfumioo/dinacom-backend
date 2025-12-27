package controllers

import (
	"backend-dinakom/app/dto/request"
	"backend-dinakom/app/helpers"
	"backend-dinakom/app/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type AIWebExtractController interface {
	GetAllWebExtracts(c *fiber.Ctx) error
	GetWebExtractByID(c *fiber.Ctx) error
	CreateWebExtract(c *fiber.Ctx) error
	UpdateWebExtract(c *fiber.Ctx) error
	DeleteWebExtract(c *fiber.Ctx) error
}

type aiWebExtractController struct {
	AIWebExtractService services.AIWebExtractService
}

func NewAIWebExtractController(AIWebExtractService services.AIWebExtractService) AIWebExtractController {
	return &aiWebExtractController{AIWebExtractService: AIWebExtractService}
}

func (ctrl *aiWebExtractController) GetAllWebExtracts(c *fiber.Ctx) error {
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

	data, totalRows, err := ctrl.AIWebExtractService.GetAllWebExtracts(page, pageSize)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return helpers.PaginationResponse(c, "AI web extracts retrieved successfully", data, page, pageSize, totalRows)
}

func (ctrl *aiWebExtractController) GetWebExtractByID(c *fiber.Ctx) error {
	ID := c.Params("id")
	data, err := ctrl.AIWebExtractService.GetWebExtractByID(ID)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return helpers.SuccessResponse(c, "AI web extract retrieved successfully", data)
}

func (ctrl *aiWebExtractController) CreateWebExtract(c *fiber.Ctx) error {
	var req request.CreateWebExtractRequest

	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	data, err := ctrl.AIWebExtractService.CreateWebExtract(req)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return helpers.CreatedResponse(c, "AI web extract created successfully", data)
}

func (ctrl *aiWebExtractController) UpdateWebExtract(c *fiber.Ctx) error {
	ID := c.Params("id")
	var req request.UpdateWebExtractRequest

	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	data, err := ctrl.AIWebExtractService.UpdateWebExtract(ID, req)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return helpers.SuccessResponse(c, "AI web extract updated successfully", data)
}

func (ctrl *aiWebExtractController) DeleteWebExtract(c *fiber.Ctx) error {
	ID := c.Params("id")

	data, err := ctrl.AIWebExtractService.DeleteWebExtract(ID)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return helpers.SuccessResponse(c, "AI web extract deleted successfully", data)
}

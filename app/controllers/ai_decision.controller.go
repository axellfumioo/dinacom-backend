package controllers

import (
	"backend-dinakom/app/dto/request"
	"backend-dinakom/app/helpers"
	"backend-dinakom/app/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type AIDecisionController interface {
	GetAllDecisions(c *fiber.Ctx) error
	CreateDecision(c *fiber.Ctx) error
}

type aiDecisionController struct {
	AIDecisionService services.AIDecisionService
}

func NewAIDecisionController(AIDecisionService services.AIDecisionService) AIDecisionController {
	return &aiDecisionController{AIDecisionService: AIDecisionService}
}

func (ctrl *aiDecisionController) GetAllDecisions(c *fiber.Ctx) error {
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

	data, totalRows, err := ctrl.AIDecisionService.GetAllDecisions(page, pageSize)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return helpers.PaginationResponse(c, "decisions retrieved successfully", data, page, pageSize, totalRows)
}

func (ctrl *aiDecisionController) CreateDecision(c *fiber.Ctx) error {
	var req request.CreateDecisionRequest

	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	data, err := ctrl.AIDecisionService.CreateDecision(req)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return helpers.CreatedResponse(c, "decision created successfully", data)
}

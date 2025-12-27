package controllers

import (
	"backend-dinakom/app/dto/request"
	"backend-dinakom/app/helpers"
	"backend-dinakom/app/services"

	"github.com/gofiber/fiber/v2"
)

type AIDecisionController interface {
	CreateDecision(c *fiber.Ctx) error
}

type aiDecisionController struct {
	AIDecisionService services.AIDecisionService
}

func NewAIDecisionController(AIDecisionService services.AIDecisionService) AIDecisionController {
	return &aiDecisionController{AIDecisionService: AIDecisionService}
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

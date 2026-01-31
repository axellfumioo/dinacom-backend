package controllers

import (
	"backend-dinakom/app/helpers"
	"backend-dinakom/app/services"

	"github.com/gofiber/fiber/v2"
)

type AIInsightController interface {
	GetLatestInsight(c *fiber.Ctx) error
}

type aiInsightController struct {
	aiInsightService services.AIInsightService
}

func NewAIInsightController(aiInsightService services.AIInsightService) AIInsightController {
	return &aiInsightController{ aiInsightService: aiInsightService }
}

func (ctrl *aiInsightController) GetLatestInsight(c *fiber.Ctx) error {
	userId := c.Locals("user_id").(string)
	data, err := ctrl.aiInsightService.GetLatestInsight(userId)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return helpers.SuccessResponse(c, "latest insight retrieved successfully", data)
}

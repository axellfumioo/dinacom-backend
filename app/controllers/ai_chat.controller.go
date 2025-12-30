package controllers

import (
	"backend-dinakom/app/helpers"
	"backend-dinakom/app/services"

	"github.com/gofiber/fiber/v2"
)

type AIChatController interface {
	CreateNewChat(c *fiber.Ctx) error
}

type aiChatController struct {
	AIChatService services.AIChatService
}

func NewAIChatController(AIChatService services.AIChatService) AIChatController {
	return &aiChatController{AIChatService: AIChatService}
}

func (ctrl *aiChatController) CreateNewChat(c *fiber.Ctx) error {
	UserID := c.Locals("user_id").(string)
	data, err := ctrl.AIChatService.CreateNewChat(UserID)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return helpers.CreatedResponse(c, "chat created successfully", data)
}

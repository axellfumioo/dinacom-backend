package controllers

import (
	"backend-dinakom/app/helpers"
	"backend-dinakom/app/services"

	"github.com/gofiber/fiber/v2"
)

type AIChatController interface {
	GetUserAIChats(c *fiber.Ctx) error
	CreateNewChat(c *fiber.Ctx) error
	DeleteAIChat(c *fiber.Ctx) error
}

type aiChatController struct {
	AIChatService services.AIChatService
}

func NewAIChatController(AIChatService services.AIChatService) AIChatController {
	return &aiChatController{AIChatService: AIChatService}
}

func (ctrl *aiChatController) GetUserAIChats(c *fiber.Ctx) error {
	UserID := c.Locals("user_id").(string)
	data, err := ctrl.AIChatService.GetUserAIChats(UserID)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return helpers.SuccessResponse(c, "user AIchats retrieved successfully", data)
}

func (ctrl *aiChatController) CreateNewChat(c *fiber.Ctx) error {
	UserID := c.Locals("user_id").(string)
	data, err := ctrl.AIChatService.CreateNewChat(UserID)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return helpers.CreatedResponse(c, "chat created successfully", data)
}

func (ctrl *aiChatController) DeleteAIChat(c *fiber.Ctx) error {
	ID := c.Params("id")
	UserID := c.Locals("user_id").(string)

	data, err := ctrl.AIChatService.DeleteAIChat(ID, UserID)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return helpers.SuccessResponse(c, "chat deleted successfully", data)
}

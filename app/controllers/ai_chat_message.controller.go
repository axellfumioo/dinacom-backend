package controllers

import (
	"backend-dinakom/app/dto/request"
	"backend-dinakom/app/helpers"
	"backend-dinakom/app/services"

	"github.com/gofiber/fiber/v2"
)

type AIChatMessageController interface {
	GetChatMessagesByChatID(c *fiber.Ctx) error
	CreateNewMessage(c *fiber.Ctx) error
	CreateNewMessageWithImage(c *fiber.Ctx) error
	DeleteMessageByID(c *fiber.Ctx) error
	DeleteChatMessages(c *fiber.Ctx) error
}

type aiChatMessageController struct {
	aiChatMessageService services.AIChatMessageService
}

func NewAIChatMessageController(AIChatMessageService services.AIChatMessageService) AIChatMessageController {
	return &aiChatMessageController{aiChatMessageService: AIChatMessageService}
}

func (ctrl *aiChatMessageController) GetChatMessagesByChatID(c *fiber.Ctx) error {
	ChatID := c.Params("chatId")

	data, err := ctrl.aiChatMessageService.GetChatMessagesByChatID(ChatID)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return helpers.SuccessResponse(c, "messages retrieved successfully", data)
}

func (ctrl *aiChatMessageController) CreateNewMessage(c *fiber.Ctx) error {
	UserID := c.Locals("user_id").(string)
	ChatID := c.Params("chatId")

	var req request.CreateMessageRequest
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.ErrBadRequest.Code, err.Error())
	}

	data, err := ctrl.aiChatMessageService.CreateNewMessage(req, UserID, ChatID)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return helpers.CreatedResponse(c, "ai message created successfully", data)
}

func (ctrl *aiChatMessageController) CreateNewMessageWithImage(c *fiber.Ctx) error {
	UserID := c.Locals("user_id").(string)
	ChatID := c.Params("chatId")

	// Multipart handler
	image, err := c.FormFile("image")
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	content := c.FormValue("content")

	// Service
	data, err := ctrl.aiChatMessageService.CreateNewMessageWithImage(request.CreateMessageWithMediaRequest{Content: content, Image: image, UserID: UserID, ChatID: ChatID})
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return helpers.CreatedResponse(c, "ai message created successfully", data)
}

func (ctrl *aiChatMessageController) DeleteMessageByID(c *fiber.Ctx) error {
	UserID := c.Locals("user_id").(string)
	ID := c.Params("messageId")

	data, err := ctrl.aiChatMessageService.DeleteChatMessageByID(ID, UserID)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return helpers.SuccessResponse(c, "ai message deleted successfully", data)
}

func (ctrl *aiChatMessageController) DeleteChatMessages(c *fiber.Ctx) error {
	UserID := c.Locals("user_id").(string)
	ID := c.Params("chatId")

	data, err := ctrl.aiChatMessageService.DeleteChatMessages(ID, UserID)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return helpers.SuccessResponse(c, "ai messages deleted successfully", data)
}

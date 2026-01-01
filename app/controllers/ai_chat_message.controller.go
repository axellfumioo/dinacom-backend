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

// GetAIChatMessageByChatID godoc
// @Summary GetAIChatMessageByChatID
// @Description endpoint to get aiChats messages
// @Tags Aichats(message)
// @Produce  json
// @Param chatID path string true "chatID"
// @Success 201 {object} response.SuccessResponse
// @Failure 400 {object} response.ErrorResponse
// @Failure 404 {object} response.ErrorResponse
// @Router /aichats/{chatID}/message [get]
func (ctrl *aiChatMessageController) GetChatMessagesByChatID(c *fiber.Ctx) error {
	ChatID := c.Params("chatId")

	data, err := ctrl.aiChatMessageService.GetChatMessagesByChatID(ChatID)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return helpers.SuccessResponse(c, "messages retrieved successfully", data)
}

// CreateNewMessage godoc
// @Summary CreateNewMessage
// @Description endpoint to create new message (akses endopoint ini kalo mau kirim chat tanpa image/gambar)
// @Tags Aichats(message)
// @Produce  json
// @Param request body request.CreateMessageRequest true "Create new message"
// @Success 201 {object} response.SuccessResponse
// @Failure 400 {object} response.ErrorResponse
// @Failure 404 {object} response.ErrorResponse
// @Router /aichats/{aichatID}/message [post]
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

// CreateNewMessageWithImage godoc
// @Summary CreateNewMessageWithImage
// @Description endpoint to create new message with image (akses endopoint ini kalo mau kirim chat dengan image/gambar, Untuk body kirim data ke formdata key "image" (berisi file gambar), "content" (berisi pesan yg dikirim))
// @Tags Aichats(message)
// @Produce  json
// @Success 201 {object} response.SuccessResponse
// @Failure 400 {object} response.ErrorResponse
// @Failure 404 {object} response.ErrorResponse
// @Router /aichats/{aichatID}/message-img [post]
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

// DeleteMessageByID godoc
// @Summary DeleteMessageByID
// @Description endpoint to delete specific message with messageID
// @Tags Aichats(message)
// @Produce  json
// @Param messageID path string true "messageID"
// @Success 201 {object} response.SuccessResponse
// @Failure 400 {object} response.ErrorResponse
// @Failure 404 {object} response.ErrorResponse
// @Router /aichats/message/{messageId} [delete]
func (ctrl *aiChatMessageController) DeleteMessageByID(c *fiber.Ctx) error {
	UserID := c.Locals("user_id").(string)
	ID := c.Params("messageId")

	data, err := ctrl.aiChatMessageService.DeleteChatMessageByID(ID, UserID)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return helpers.SuccessResponse(c, "ai message deleted successfully", data)
}

// DeletechatMessages godoc
// @Summary DeleteChatMessages
// @Description endpoint to delete all messages in the chat (Hapus semua message di chat)
// @Tags Aichats(message)
// @Produce  json
// @Param chatID path string true "chatID"
// @Success 201 {object} response.SuccessResponse
// @Failure 400 {object} response.ErrorResponse
// @Failure 404 {object} response.ErrorResponse
// @Router /aichats/{chatId}/message [delete]
func (ctrl *aiChatMessageController) DeleteChatMessages(c *fiber.Ctx) error {
	UserID := c.Locals("user_id").(string)
	ID := c.Params("chatId")

	data, err := ctrl.aiChatMessageService.DeleteChatMessages(ID, UserID)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return helpers.SuccessResponse(c, "ai messages deleted successfully", data)
}

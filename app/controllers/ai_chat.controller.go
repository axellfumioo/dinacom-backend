package controllers

import (
	"backend-dinakom/app/helpers"
	"backend-dinakom/app/services"

	"github.com/gofiber/fiber/v2"
)

type AIChatController interface {
	GetAIChatByID(c *fiber.Ctx) error
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

func (ctrl *aiChatController) GetAIChatByID(c *fiber.Ctx) error {
	ID := c.Params("id")
	UserID := c.Locals("user_id").(string)
	data, err := ctrl.AIChatService.GetAIChatByID(ID, UserID)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return helpers.SuccessResponse(c, "AIchat retrieved successfully", data)
}


// GetUserAIChats godoc
// @Summary GetUserAIChats
// @Description endpoint to get user aiChats (room)
// @Tags Aichats
// @Produce  json
// @Success 201 {object} response.SuccessResponse
// @Failure 400 {object} response.ErrorResponse
// @Failure 404 {object} response.ErrorResponse
// @Router /aichats/user [get]
func (ctrl *aiChatController) GetUserAIChats(c *fiber.Ctx) error {
	UserID := c.Locals("user_id").(string)
	data, err := ctrl.AIChatService.GetUserAIChats(UserID)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return helpers.SuccessResponse(c, "user AIchats retrieved successfully", data)
}

// CreateNewChat godoc
// @Summary CreateNewChat
// @Description endpoint to create new chat (btw ini cara kerjanya mungkin ntr ada tombol buat tambah chat dengan a.i (kalau belum ada chat ai) trs nembak ke endpoint ini )
// @Tags Aichats
// @Produce  json
// @Success 201 {object} response.SuccessResponse
// @Failure 400 {object} response.ErrorResponse
// @Failure 404 {object} response.ErrorResponse
// @Router /aichats/ [post]
func (ctrl *aiChatController) CreateNewChat(c *fiber.Ctx) error {
	UserID := c.Locals("user_id").(string)
	data, err := ctrl.AIChatService.CreateNewChat(UserID)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return helpers.CreatedResponse(c, "chat created successfully", data)
}

// DeleteChat godoc
// @Summary DeleteChat
// @Description endpoint to delete specific chat with chatID
// @Tags Aichats
// @Produce  json
// @Param chatID path string true "chatID"
// @Success 201 {object} response.SuccessResponse
// @Failure 400 {object} response.ErrorResponse
// @Failure 404 {object} response.ErrorResponse
// @Router /aichats/{id} [delete]
func (ctrl *aiChatController) DeleteAIChat(c *fiber.Ctx) error {
	ID := c.Params("id")
	UserID := c.Locals("user_id").(string)

	data, err := ctrl.AIChatService.DeleteAIChat(ID, UserID)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return helpers.SuccessResponse(c, "chat deleted successfully", data)
}

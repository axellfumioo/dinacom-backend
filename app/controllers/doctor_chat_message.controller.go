package controllers

import (
	"backend-dinakom/app/dto/request"
	"backend-dinakom/app/helpers"
	"backend-dinakom/app/services"

	"github.com/gofiber/fiber/v2"
)

type DoctorChatMessageController interface {
	GetMessagesByRoomID(c *fiber.Ctx) error
	CreateNewMessage(c *fiber.Ctx) error
}

type doctorChatMessageController struct {
	doctorChatMessageService services.DoctorChatMessageService
}

func NewDoctorChatMessageController(doctorChatMessageService services.DoctorChatMessageService) DoctorChatMessageController {
	return &doctorChatMessageController{doctorChatMessageService: doctorChatMessageService}
}

func (ctrl *doctorChatMessageController) GetMessagesByRoomID(c *fiber.Ctx) error {
	roomID := c.Params("roomID")
	messages, err := ctrl.doctorChatMessageService.GetMessagesByRoomID(roomID)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return helpers.SuccessResponse(c, "messages retrieved successfully", messages)
}

func (ctrl *doctorChatMessageController) CreateNewMessage(c *fiber.Ctx) error {
	senderID := c.Locals("user_id").(string)

	var req request.CreateDoctorChatMessageRequest
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	data, err := ctrl.doctorChatMessageService.CreateNewMessage(senderID, req)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return helpers.CreatedResponse(c, "message created successfully", data)
}

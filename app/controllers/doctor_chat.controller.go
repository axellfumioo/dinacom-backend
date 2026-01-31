package controllers

import (
	"backend-dinakom/app/dto/request"
	"backend-dinakom/app/helpers"
	"backend-dinakom/app/services"

	"github.com/gofiber/fiber/v2"
)

type DoctorChatController interface {
	GetAllDoctors(c *fiber.Ctx) error
	GetAllUserDoctorChatRooms(c *fiber.Ctx) error
	GetOrCreateDoctor(c *fiber.Ctx) error
	CreateDoctorChatroom(c *fiber.Ctx) error
}

type doctorChatController struct {
	doctorChatService services.DoctorChatService
}

func NewDoctorChatController(doctorChatService services.DoctorChatService) DoctorChatController {
	return &doctorChatController{doctorChatService: doctorChatService}
}

func (ctrl *doctorChatController) GetAllDoctors(c *fiber.Ctx) error {
	doctors, err := ctrl.doctorChatService.GetAllDoctors()
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return helpers.SuccessResponse(c, "doctors retrieved successfully", doctors)
}

func (ctrl *doctorChatController) GetAllUserDoctorChatRooms(c *fiber.Ctx) error {
	var userID = c.Locals("user_id").(string)
	chatRooms, err := ctrl.doctorChatService.GetAllUserDoctorChats(userID)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return helpers.SuccessResponse(c, "doctor chat rooms retrieved successfully", chatRooms)
}

func (ctrl *doctorChatController) GetOrCreateDoctor(c *fiber.Ctx) error {
	var userID = c.Locals("user_id").(string)
	doctorID := c.Params("doctorID")
	chatRoom, err := ctrl.doctorChatService.GetOrCreateDoctorChatRoom(userID, doctorID)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return helpers.SuccessResponse(c, "doctor chat room retrieved or created successfully", chatRoom)
}

func (ctrl *doctorChatController) CreateDoctorChatroom(c *fiber.Ctx) error {
	var userID = c.Locals("user_id").(string)
	var req request.CreateDoctorChatroomRequest
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	chatRoom, err := ctrl.doctorChatService.CreateDoctorChatroom(userID, req.DoctorID)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return helpers.SuccessResponse(c, "doctor chat room created successfully", chatRoom)
}

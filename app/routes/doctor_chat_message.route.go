package router

import (
	"backend-dinakom/app/controllers"
	"backend-dinakom/app/middlewares"
	"backend-dinakom/app/services"
	"backend-dinakom/database"

	"github.com/gofiber/fiber/v2"
)

func DoctorChaMessageRoute(r fiber.Router) {
	doctorChatMessageService := services.NewDoctorChatMessageService(database.DB)
	doctorChatController := controllers.NewDoctorChatMessageController(doctorChatMessageService)
	doctorChats := r.Group("doctor-messages", middlewares.AuthMiddleware())
	doctorChats.Post("/", doctorChatController.CreateNewMessage)
	doctorChats.Get("/:roomId", doctorChatController.GetMessagesByRoomID)}
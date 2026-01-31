package router

import (
	"backend-dinakom/app/controllers"
	"backend-dinakom/app/middlewares"
	"backend-dinakom/app/services"
	"backend-dinakom/database"

	"github.com/gofiber/fiber/v2"
)

func DoctorChatRoute(r fiber.Router) {
	var doctorChatService = services.NewDoctorChatService(database.DB)
	var doctorChatController = controllers.NewDoctorChatController(doctorChatService)

	var doctorChats = r.Group("doctor-chats", middlewares.AuthMiddleware())
	doctorChats.Post("/", doctorChatController.CreateDoctorChatroom)
	doctorChats.Get("/", doctorChatController.GetAllDoctors)
	doctorChats.Get("/user", doctorChatController.GetAllUserDoctorChatRooms)
	doctorChats.Get("/:doctorID", doctorChatController.GetOrCreateDoctor)
}

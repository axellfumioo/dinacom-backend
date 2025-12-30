package router

import (
	"backend-dinakom/app/controllers"
	"backend-dinakom/app/middlewares"
	"backend-dinakom/app/services"
	"backend-dinakom/database"

	"github.com/gofiber/fiber/v2"
)

func AIChatRoute(r fiber.Router) {
	aiChatService := services.NewAIChatService(database.GetDb())
	aiChatController := controllers.NewAIChatController(aiChatService)

	aiChats := r.Group("/aichats")
	aiChats.Use(middlewares.AuthMiddleware())
	aiChats.Post("/", aiChatController.CreateNewChat)
}

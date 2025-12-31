package router

import (
	"backend-dinakom/app/controllers"
	"backend-dinakom/app/middlewares"
	"backend-dinakom/app/services"
	"backend-dinakom/configs"
	"backend-dinakom/database"

	"github.com/gofiber/fiber/v2"
)

func AIChatRoute(r fiber.Router) {
	aiChatService := services.NewAIChatService(database.GetDb())
	aiChatMessageService := services.NewAIChatMessageService(database.DB, configs.QueueClient, configs.MinioClient)
	aiChatController := controllers.NewAIChatController(aiChatService)
	aiChatMessageController := controllers.NewAIChatMessageController(aiChatMessageService)

	aiChats := r.Group("/aichats")
	aiChats.Use(middlewares.AuthMiddleware())
	aiChats.Get("/user", aiChatController.GetUserAIChats)
	aiChats.Post("/", aiChatController.CreateNewChat)
	aiChats.Delete("/:id", aiChatController.DeleteAIChat)

	aiChats.Get("/:chatID/message", aiChatMessageController.GetChatMessagesByChatID)
	aiChats.Post("/:chatId/message", aiChatMessageController.CreateNewMessage)
	aiChats.Post("/:chatId/message-img", aiChatMessageController.CreateNewMessageWithImage)
}

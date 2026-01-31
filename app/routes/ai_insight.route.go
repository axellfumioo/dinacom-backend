package router

import (
	"backend-dinakom/app/controllers"
	"backend-dinakom/app/middlewares"
	"backend-dinakom/app/services"
	"backend-dinakom/database"

	"github.com/gofiber/fiber/v2"
)

func AIInsightRoute(r fiber.Router) {
	aiInsightService := services.NewAIInsightService(database.DB)
	aiInsightController := controllers.NewAIInsightController(aiInsightService)
	aiInsight := r.Group("ai-insights", middlewares.AuthMiddleware())
	aiInsight.Get("/latest", aiInsightController.GetLatestInsight)
}

package router

import (
	"backend-dinakom/app/controllers"
	"backend-dinakom/app/services"
	"backend-dinakom/database"

	"github.com/gofiber/fiber/v2"
)

func AIRoute(r fiber.Router) {
	// Services
	aiDecisionService := services.NewAIDecisionService(database.GetDb())

	// Controllers
	aiDecisionController := controllers.NewAIDecisionController(aiDecisionService)

	ai := r.Group("/ai")
	decision := ai.Group("/decision")
	decision.Get("/", aiDecisionController.GetAllDecisions)
	decision.Get("/:id/get", aiDecisionController.GetDecisionByID)
	decision.Post("/", aiDecisionController.CreateDecision)
}

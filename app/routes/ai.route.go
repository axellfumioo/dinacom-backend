package router

import (
	"backend-dinakom/app/controllers"
	"backend-dinakom/app/services"
	"backend-dinakom/database"

	"github.com/gofiber/fiber/v2"
)

func AIRoute(r fiber.Router) {
	// Services
	aiDecisionService := services.NewAIDecisionService(database.DB)
	aiGoogleSearchService := services.NewAIGoogleSearchService(database.DB)

	// Controllers
	aiDecisionController := controllers.NewAIDecisionController(aiDecisionService)
	aiGoogleSearchController := controllers.NewAIGoogleSearchController(aiGoogleSearchService)

	ai := r.Group("/ai")
	// Decision
	decision := ai.Group("/decision")
	decision.Get("/", aiDecisionController.GetAllDecisions)
	decision.Get("/:id/get", aiDecisionController.GetDecisionByID)
	decision.Post("/", aiDecisionController.CreateDecision)
	decision.Patch("/:id/update", aiDecisionController.UpdateDecision)
	decision.Delete("/:id/delete", aiDecisionController.DeleteDecision)

	// Google Search
	googleSearch := ai.Group("/search")
	googleSearch.Get("/", aiGoogleSearchController.GetAllGoogleSearchs)
	googleSearch.Get("/:id/get", aiGoogleSearchController.GetGoogleSearchByID)
	googleSearch.Post("/", aiGoogleSearchController.CreateGoogleSearch)
	googleSearch.Patch("/:id/update", aiGoogleSearchController.UpdateGoogleSearch)
	googleSearch.Delete("/:id/delete", aiGoogleSearchController.DeleteGoogleSearch)
}

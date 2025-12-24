package router

import (
	"backend-dinakom/app/controllers"
	"backend-dinakom/app/middlewares"
	"backend-dinakom/app/services"
	"backend-dinakom/database"

	"github.com/gofiber/fiber/v2"
)

func FoodScanResultRoute(r fiber.Router) {
	fsResultService := services.NewFoodScanResultService(database.GetDb())
	fsResultController := controllers.NewFoodScanResultController(fsResultService)

	fsResults := r.Group("results")
	fsResults.Use(middlewares.AuthMiddleware())
	fsResults.Get("/", fsResultController.GetAllResults)
	fsResults.Get("/:id/get", fsResultController.GetResultByID)
}

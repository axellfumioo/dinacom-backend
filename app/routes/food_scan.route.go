package router

import (
	"backend-dinakom/app/controllers"
	"backend-dinakom/app/middlewares"
	"backend-dinakom/app/services"
	"backend-dinakom/configs"
	"backend-dinakom/database"

	"github.com/gofiber/fiber/v2"
)

func FoodScanRoute(r fiber.Router) {
	minioClient := configs.NewMinioClient()

	foodSCanService := services.NewFoodScanService(database.GetDb(), minioClient)
	foodScanController := controllers.NewFoodScanController(foodSCanService)

	foodScans := r.Group("foodscans")
	foodScans.Use(middlewares.AuthMiddleware())
	foodScans.Post("/scan", foodScanController.ScanFood)
}

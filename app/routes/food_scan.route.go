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
	minioClient := configs.MinioClient
	queueClient := configs.QueueClient

	foodScanService := services.NewFoodScanService(database.GetDb(), minioClient, queueClient)
	foodScanController := controllers.NewFoodScanController(foodScanService)

	foodScans := r.Group("foodscans")
	foodScans.Use(middlewares.AuthMiddleware())
	foodScans.Get("/", middlewares.RoleMiddleware("ADMIN"), foodScanController.GetAllFoodScans)
	foodScans.Post("/scan", foodScanController.ScanFood)
	foodScans.Get("/user", foodScanController.GetUserFoodScans)
	foodScans.Get("/:id/get", foodScanController.GetFoodScanByID)
}

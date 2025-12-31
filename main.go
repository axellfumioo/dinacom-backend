package main

import (
	"backend-dinakom/app/middlewares"
	router "backend-dinakom/app/routes"
	"backend-dinakom/configs"
	"backend-dinakom/database"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"

	_ "backend-dinakom/docs"
	docs "backend-dinakom/docs"
)

// @title NutriOne API
// @version 1.0
// @description Dokumentasi API untuk Nutrione

// @host localhost:8080
// @BasePath /api/v1

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	configs.LoadConfig()

	// Connect ke database
	database.ConnectDatabase()
	database.RunMigration()

	// Init Client
	configs.ConnectSocketIO()
	configs.InitQueueClient()
	configs.InitMinioClient()
	configs.InitRestyClient()

	app := fiber.New(fiber.Config{
		AppName:      configs.AppConfig.App.Name,
		ErrorHandler: middlewares.ErrorMiddleware,
	})
	router.SetupRouter(app)
	defer configs.QueueClient.Close()

	docs.SwaggerInfo.Title = "Nutrione API"
	docs.SwaggerInfo.Description = "Documentation For Nutrione API"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/api/v1"

	port := configs.AppConfig.App.Port
	log.Printf("Starting %s server on port %s", configs.AppConfig.App.Name, port)

	if err := app.Listen(":" + port); err != nil {
		log.Fatal(fmt.Sprintf("Failed to start server: %v", err))
	}
}

package main

import (
	"backend-dinakom/app/middlewares"
	router "backend-dinakom/app/routes"
	"backend-dinakom/app/socket"
	"backend-dinakom/configs"
	"backend-dinakom/database"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

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
	database.RunSeeder(database.DB)
	// Init Client
	configs.ConnectSocketIO()
	configs.InitQueueClient()
	configs.InitMinioClient()
	configs.InitRestyClient()

	// Connect socket
	socket.StartConnectionHandler()

	app := fiber.New(fiber.Config{
		AppName: configs.AppConfig.App.Name,

		ErrorHandler: middlewares.ErrorMiddleware,
	})

	app.Use(cors.New(cors.Config{
		AllowOrigins: "htpp://localhost:3000",
		AllowMethods: "GET,POST,PUT,PATCH,DELETE",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))
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
		log.Fatalf(fmt.Sprintf("Failed to start server: %v", err))
	}
}

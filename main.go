package main

import (
	"backend-dinakom/app/middlewares"
	"backend-dinakom/app/routes"
	"backend-dinakom/configs"
	"backend-dinakom/database"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	configs.LoadConfig()

	// Connect ke database
	database.ConnectDatabase()
	database.RunMigration()

	// Init Client
	configs.InitQueueClient()
	configs.InitMinioClient()
	configs.InitRestyClient()

	app := fiber.New(fiber.Config{
		AppName:      configs.AppConfig.App.Name,
		ErrorHandler: middlewares.ErrorMiddleware,
	})
	router.SetupRouter(app)
	defer configs.QueueClient.Close()

	port := configs.AppConfig.App.Port
	log.Printf("Starting %s server on port %s", configs.AppConfig.App.Name, port)

	if err := app.Listen(":" + port); err != nil {
		log.Fatal(fmt.Sprintf("Failed to start server: %v", err))
	}
}

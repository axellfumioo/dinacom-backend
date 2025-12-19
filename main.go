package main

import (
	router "backend-dinakom/app/routes"
	"backend-dinakom/configs"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	configs.LoadConfig()

	app := fiber.New()
	router.Route(app)

	port := configs.AppConfig.App.Port
	log.Printf("Starting %s server on port %s", configs.AppConfig.App.Name, port)

	if err := app.Listen(":" + port); err != nil {
		log.Fatal(fmt.Sprintf("Failed to start server: %v", err))
	}
}

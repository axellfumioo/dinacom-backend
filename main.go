package main

import (
	router "backend-dinakom/app/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	router.Route(app)

	app.Listen(":8080")
}

package router

import (
	"github.com/gofiber/fiber/v2"
)

func SetupRouter(app *fiber.App) {
	api := app.Group("/api/v1")
	api.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":  fiber.StatusOK,
			"message": "Server is running",
		})
	})

	UserRoute(api)
	AuthRoute(api)
	ProfileRoute(api)
	RoleRoute(api)
	FoodScanRoute(api)
}

package router

import (
	"backend-dinakom/app/workers"
	"backend-dinakom/database"

	fiberSwagger "github.com/swaggo/fiber-swagger"

	"github.com/gofiber/fiber/v2"
)

func SetupRouter(app *fiber.App) {
	api := app.Group("/api/v1")
	api.Get("/swagger/*", fiberSwagger.WrapHandler)
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
	FoodScanResultRoute(api)
	UserMealRoute(api)

	go workers.StartWorker(database.GetDb())
}

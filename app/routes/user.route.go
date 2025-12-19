package router

import "github.com/gofiber/fiber/v2"

func UserRoute(r fiber.Router) {
	userRoute := r.Group("/users")
	userRoute.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":  fiber.StatusOK,
			"message": "Hello User",
		})
	})
}

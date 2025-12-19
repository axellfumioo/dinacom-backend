package router

import (
	"backend-dinakom/app/helpers"

	"github.com/gofiber/fiber/v2"
)

func UserRoute(r fiber.Router) {
	userRoute := r.Group("/users")
	userRoute.Get("/", func(c *fiber.Ctx) error {
		return helpers.SuccessResponse(c, "hello users", "ini data")
	})
}

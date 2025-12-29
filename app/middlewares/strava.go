package middlewares

import (
	"github.com/gofiber/fiber/v2"
)

func StravaMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		stravaIntegrate, _ := c.Locals("strava_integrated").(bool)
		if stravaIntegrate == false {
			return fiber.NewError(fiber.StatusForbidden, "access denied, this user has not yet integrated with Strava account")
		}

		return c.Next()
	}
}

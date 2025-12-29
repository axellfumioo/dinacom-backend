package middlewares

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func StravaMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		stravaIntegrate, _ := strconv.ParseBool(c.Locals("strava_integrate").(string))
		if stravaIntegrate == false {
			return fiber.NewError(fiber.StatusForbidden, "access denied, this user has not yet integrated with Strava account")
		}

		return c.Next()
	}
}

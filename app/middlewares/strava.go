package middlewares

import (
	"backend-dinakom/app/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func StravaMiddleware(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var userID = c.Locals("user_id").(string)
		var user models.User
		if err := db.First(&user, "user_id =?", userID).Error; err != nil {
			return fiber.NewError(fiber.StatusForbidden, "access denied, user not found")
		}

		if user.StravaIntegrated == false {
			return fiber.NewError(fiber.StatusForbidden, "access denied, this user has not yet integrated with Strava account")
		}

		return c.Next()
	}
}

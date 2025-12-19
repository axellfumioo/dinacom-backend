package middlewares

import (
	"github.com/gofiber/fiber/v2"
)

func RoleMiddleware(roles ...string) fiber.Handler {
	allowed := make(map[string]bool)
	for _, r := range roles {
		allowed[r] = true
	}

	return func(c *fiber.Ctx) error {
		userRole := c.Locals("user_role").(string)

		if allowed[userRole] {
			return c.Next()
		}

		return fiber.NewError(fiber.StatusForbidden, "access denied, invalid role")
	}
}

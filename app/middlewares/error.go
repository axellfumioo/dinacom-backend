package middlewares

import (
	"backend-dinakom/app/dto/response"
	"log"

	"github.com/gofiber/fiber/v2"
)

func ErrorMiddleware(c *fiber.Ctx, err error) error {
	// Error default
	status := fiber.StatusInternalServerError
	message := "Internal Server Error"

	// Jika error instance dari (fiber.newError())
	if e, ok := err.(*fiber.Error); ok {
		status = e.Code
		message = e.Message
	}

	// LOG internal error (penting)
	log.Println(err)

	// Response
	return c.Status(status).JSON(response.BaseResponse{
		Success:    false,
		StatusCode: status,
		Message:    message,
	})
}

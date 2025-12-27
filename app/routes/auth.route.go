package router

import (
	"backend-dinakom/app/controllers"
	"backend-dinakom/app/services"
	"backend-dinakom/database"

	"github.com/gofiber/fiber/v2"
)

func AuthRoute(r fiber.Router) {
	authService := services.NewAuthService(*database.GetDb())
	authController := controllers.NewAuthController(authService)

	auth := r.Group("auth")
	auth.Post("register", authController.Register)
	auth.Post("login", authController.Login)
	auth.Get("/strava/redirect", authController.StravaRedirect)
	auth.Get("/strava/callback", authController.StravaCallback)
}

package router

import (
	"backend-dinakom/app/controllers"
	"backend-dinakom/app/middlewares"
	"backend-dinakom/app/services"
	"backend-dinakom/configs"
	"backend-dinakom/database"

	"github.com/gofiber/fiber/v2"
)

func StravaRoute(r fiber.Router) {
	stravaService := services.NewStravaService(database.DB, configs.RestyClient)
	stravaController := controllers.NewStravaController(stravaService)
	_ = stravaController
	strava := r.Group("strava")
	strava.Use(middlewares.AuthMiddleware())
}

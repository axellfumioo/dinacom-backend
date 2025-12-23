package router

import (
	"backend-dinakom/app/controllers"
	"backend-dinakom/app/middlewares"
	"backend-dinakom/app/services"
	"backend-dinakom/database"

	"github.com/gofiber/fiber/v2"
)

func UserMealRoute(r fiber.Router) {
	userMealService := services.NewUserMealService(database.GetDb())
	userMealController := controllers.NewUserMealController(userMealService)

	userMeals := r.Group("usermeals")
	userMeals.Get("/", middlewares.AuthMiddleware(), userMealController.GetAllUserMeals)
}

package router

import (
	"backend-dinakom/app/controllers"
	"backend-dinakom/app/middlewares"
	"backend-dinakom/app/services"
	"backend-dinakom/configs"
	"backend-dinakom/database"

	"github.com/gofiber/fiber/v2"
)

func UserMealRoute(r fiber.Router) {
	userMealService := services.NewUserMealService(database.GetDb(), configs.QueueClient)
	userMealController := controllers.NewUserMealController(userMealService)

	userMeals := r.Group("usermeals")
	userMeals.Use(middlewares.AuthMiddleware())
	userMeals.Post("/", userMealController.CreateUserMeal)
	userMeals.Get("/", userMealController.GetAllUserMeals)
	userMeals.Get("/today", userMealController.GetTodayUserMeals)
}

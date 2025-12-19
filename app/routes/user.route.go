package router

import (
	"backend-dinakom/app/controllers"
	"backend-dinakom/app/middlewares"
	"backend-dinakom/app/services"
	"backend-dinakom/database"

	"github.com/gofiber/fiber/v2"
)

func UserRoute(r fiber.Router) {
	userService := services.NewUserService(database.GetDb())
	userController := controllers.NewUserController(userService)

	users := r.Group("/users")
	users.Get("", middlewares.AuthMiddleware(), middlewares.RoleMiddleware("ADMIN"), userController.GetAllUsers)
	users.Get("/session", middlewares.AuthMiddleware(), userController.GetUserSession)
	users.Get("/:id/get", userController.GetUserByID)
	users.Post("", userController.CreateUser)
	users.Put("/:id", userController.UpdateUser)
	users.Delete("/:id", userController.DeleteUser)
	users.Post("/change-password", userController.ChangePassword)
	users.Get("/role", userController.GetUserByRoleName)
	users.Get("/total", userController.GetTotalUsers)
}

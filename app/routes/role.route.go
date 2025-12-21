package router

import (
	"backend-dinakom/app/controllers"
	"backend-dinakom/app/middlewares"
	"backend-dinakom/app/services"
	"backend-dinakom/database"

	"github.com/gofiber/fiber/v2"
)

func RoleRoute(r fiber.Router) {
	roleService := services.NewRoleService(database.GetDb())
	roleController := controllers.NewRoleController(roleService)

	roles := r.Group("roles")
	roles.Use(middlewares.AuthMiddleware(), middlewares.RoleMiddleware("USER"))
	{
		roles.Get("", roleController.GetAllRoles)
		roles.Post("create", roleController.CreateRole)
		roles.Patch("/:id/update", roleController.UpdateRole)
		roles.Delete("/:id", roleController.DeleteRole)
	}
}

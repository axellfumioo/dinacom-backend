package router

import (
	"backend-dinakom/app/controllers"
	"backend-dinakom/app/middlewares"
	"backend-dinakom/app/services"
	"backend-dinakom/configs"
	"backend-dinakom/database"

	"github.com/gofiber/fiber/v2"
)

func FamilyRoute(r fiber.Router) {
	familyService := services.NewFamilyService(database.DB, configs.MinioClient)
	familyController := controllers.NewFamilyController(familyService)

	families := r.Group("families")
	families.Use(middlewares.AuthMiddleware())
	families.Post("/", familyController.CreateNewFamily)
}

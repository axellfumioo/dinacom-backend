package router

import (
	"backend-dinakom/app/controllers"
	"backend-dinakom/app/middlewares"
	"backend-dinakom/app/services"
	"backend-dinakom/configs"
	"backend-dinakom/database"

	"github.com/gofiber/fiber/v2"
)

func ProfileRoute(r fiber.Router) {
	minioClient := configs.MinioClient
	profileService := services.NewProfileService(database.GetDb(), minioClient)
	profileController := controllers.NewProfileController(profileService)

	profiles := r.Group("profiles")
	profiles.Use(middlewares.AuthMiddleware())
	profiles.Get("/", profileController.GetUserProfile)
	profiles.Post("/avatar", profileController.UploadAvatar)
	profiles.Patch("/", profileController.UpdateProfile)
}

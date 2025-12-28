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
	profiles.Post("/avatar", middlewares.AuthMiddleware(), profileController.UploadAvatar)
	profiles.Patch("/", middlewares.AuthMiddleware(), profileController.UpdateProfile)
}

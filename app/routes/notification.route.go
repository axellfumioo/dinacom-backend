package router

import (
	"backend-dinakom/app/controllers"
	"backend-dinakom/app/services"
	"backend-dinakom/database"

	"github.com/gofiber/fiber/v2"
)

func NotificationRoute(r fiber.Router) {
	notificationService := services.NewNotificationService(database.DB);
	notificationController := controllers.NewNotificationController(notificationService);

	notifications :=  r.Group("notifications")
	notifications.Post("/daily", notificationController.DailyReminder)
}
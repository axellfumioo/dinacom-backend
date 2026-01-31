package controllers

import (
	"backend-dinakom/app/helpers"
	"backend-dinakom/app/services"

	"github.com/gofiber/fiber/v2"
)

type NotificationController interface {
	DailyReminder(c * fiber.Ctx) error 
}

type notificationController struct {
	notificationService services.NotificationService
}

func NewNotificationController( notificationService services.NotificationService ) NotificationController {
	return &notificationController{ notificationService: notificationService }
}

func (ctrl *notificationController) DailyReminder(c * fiber.Ctx) error {
	data , err := ctrl.notificationService.DailyReminder()
	if (err != nil) {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return helpers.SuccessResponse(c, "reminder send successfully", data)
}
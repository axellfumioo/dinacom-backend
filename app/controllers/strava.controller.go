package controllers

import (
	"backend-dinakom/app/helpers"
	"backend-dinakom/app/services"

	"github.com/gofiber/fiber/v2"
)

type StravaController interface {
	GetStravaProfile(c *fiber.Ctx) error
}

type stravaController struct {
	stravaService services.StravaService
}

func NewStravaController(stravaService services.StravaService) StravaController {
	return &stravaController{stravaService: stravaService}
}

func (ctrl *stravaController) GetStravaProfile(c *fiber.Ctx) error {
	UserID := c.Locals("user_id").(string)

	data, err := ctrl.stravaService.GetStravaProfile(UserID)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return helpers.SuccessResponse(c, "strava profile retrieved successfully", data)
}

package controllers

import (
	"backend-dinakom/app/helpers"
	"backend-dinakom/app/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type StravaController interface {
	GetStravaProfile(c *fiber.Ctx) error
	GetStravaActivities(c *fiber.Ctx) error
	GetStravaActivityByID(c *fiber.Ctx) error
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

func (ctrl *stravaController) GetStravaActivities(c *fiber.Ctx) error {
	UserID := c.Locals("user_id").(string)
	page, _ := strconv.Atoi(c.Query("page", "1"))
	pageSize, _ := strconv.Atoi(c.Query("page_size", "10"))

	data, err := ctrl.stravaService.GetStravaActivities(UserID, page, pageSize)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return helpers.SuccessResponse(c, "strava activities retrieved successfully", data)
}

func (ctrl *stravaController) GetStravaActivityByID(c *fiber.Ctx) error {
	UserID := c.Locals("user_id").(string)
	ID, _ := strconv.Atoi(c.Params("id"))

	data, err := ctrl.stravaService.GetStravaActivityByID(ID, UserID)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return helpers.SuccessResponse(c, "strava activity retrieved successfully", data)
}

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

// GetStravaProfile godoc
// @Summary GetStravaProfile
// @Description Access this endpoint to get user strava profile
// @Tags Strava
// @Produce  json
// @Success 200 {object} response.SuccessResponse
// @Failure 400 {object} response.ErrorResponse
// @Failure 404 {object} response.ErrorResponse
// @Router /strava/profile [get]
// @Security BearerAuth
func (ctrl *stravaController) GetStravaProfile(c *fiber.Ctx) error {
	UserID := c.Locals("user_id").(string)

	data, err := ctrl.stravaService.GetStravaProfile(UserID)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return helpers.SuccessResponse(c, "strava profile retrieved successfully", data)
}

// GetStravaActivities godoc
// @Summary GetStravaActivities
// @Description Access this endpoint to get user strava activities
// @Tags Strava
// @Produce  json
// @Param page query int false "Page number" default(1)
// @Param page_size query int false "Page size" default(10) minimum(1) maximum(100)
// @Success 200 {object} response.SuccessResponse
// @Failure 400 {object} response.ErrorResponse
// @Failure 404 {object} response.ErrorResponse
// @Router /strava/activities [get]
// @Security BearerAuth
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

// GetStravaActivityByID godoc
// @Summary GetStravaActivityByID
// @Description Access this endpoint to get user strava activity by ID
// @Tags Strava
// @Produce  json
// @Param id path string true "ID"
// @Success 200 {object} response.SuccessResponse
// @Failure 400 {object} response.ErrorResponse
// @Failure 404 {object} response.ErrorResponse
// @Router /strava/activities/{id} [get]
// @Security BearerAuth
func (ctrl *stravaController) GetStravaActivityByID(c *fiber.Ctx) error {
	UserID := c.Locals("user_id").(string)
	ID, _ := strconv.Atoi(c.Params("id"))

	data, err := ctrl.stravaService.GetStravaActivityByID(ID, UserID)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return helpers.SuccessResponse(c, "strava activity retrieved successfully", data)
}

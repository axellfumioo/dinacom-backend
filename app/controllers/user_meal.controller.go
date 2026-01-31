package controllers

import (
	"backend-dinakom/app/dto/request"
	"backend-dinakom/app/helpers"
	"backend-dinakom/app/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type UserMealController interface {
	GetAllUserMeals(c *fiber.Ctx) error
	GetTodayUserMeals(c *fiber.Ctx) error
	GetLatestUserMeal(c *fiber.Ctx) error
	CreateUserMeal(c *fiber.Ctx) error
}

type userMealController struct {
	userMealService services.UserMealService
}

func NewUserMealController(userMealService services.UserMealService) UserMealController {
	return &userMealController{userMealService: userMealService}
}

// GetAllUserMeals godoc
// @Summary GetAllUserMeals
// @Description Endpoint to get all user meals data
// @Tags Usermeals
// @Produce  json
// @Param page query int false "Page number" default(1)
// @Param page_size query int false "Page size" default(10) minimum(1) maximum(100)
// @Success 200 {object} response.PaginationResponse
// @Failure 400 {object} response.ErrorResponse
// @Failure 404 {object} response.ErrorResponse
// @Router /usermeals [get]
// @Security BearerAuth
func (ctrl *userMealController) GetAllUserMeals(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(string)
	page, _ := strconv.Atoi(c.Query("page", "1"))
	pageSize, _ := strconv.Atoi(c.Query("page_size", "10"))

	userMeals, totalRows, err := ctrl.userMealService.GetAllUserMeals(userID, page, pageSize)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return helpers.PaginationResponse(c, "usermeals data retrieved successfully", userMeals, page, pageSize, totalRows)
}

// GetLatestUserMeal godoc
// @Summary GetLatestUserMeal
// @Description Endpoint to get latest user meal data
// @Tags Usermeals
// @Produce  json
// @Success 200 {object} response.SuccessResponse
// @Failure 400 {object} response.ErrorResponse
// @Failure 404 {object} response.ErrorResponse
// @Router /usermeals/latest [get]
// @Security BearerAuth
func (ctrl *userMealController) GetLatestUserMeal(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(string)
	userMeal, err := ctrl.userMealService.GetLatestUserMeal(userID)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return helpers.SuccessResponse(c, "latest user meal retrieved successfully", userMeal)
}

// GetTodayUserMeals godoc
// @Summary GetTodayUserMeals
// @Description Endpoint to get all today user meals data
// @Tags Usermeals
// @Produce  json
// @Success 200 {object} response.SuccessResponse
// @Failure 400 {object} response.ErrorResponse
// @Failure 404 {object} response.ErrorResponse
// @Router /usermeals/today [get]
// @Security BearerAuth
func (ctrl *userMealController) GetTodayUserMeals(c *fiber.Ctx) error {
	UserID := c.Locals("user_id").(string)
	data, err := ctrl.userMealService.GetTodayUserMeals(UserID)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return helpers.SuccessResponse(c, "today user meals retrieved successfully", data)
}

// CreateUserMeal godoc
// @Summary Add new Usermeal
// @Description Add new Usermeal
// @Tags Usermeals
// @Accept json
// @Produce json
// @Param request body request.CreateUserMealRequest true "Usermeal create data"
// @Success 200 {object} response.SuccessResponse
// @Failure 400 {object} response.ErrorResponse
// @Router /usermeals [post]
// @Security BearerAuth
func (ctrl *userMealController) CreateUserMeal(c *fiber.Ctx) error {
	UserID := c.Locals("user_id").(string)
	var req request.CreateUserMealRequest
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	data, err := ctrl.userMealService.CreateUserMeal(UserID, req)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return helpers.CreatedResponse(c, "user meal created successfully", data)
}

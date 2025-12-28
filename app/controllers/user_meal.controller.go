package controllers

import (
	"backend-dinakom/app/helpers"
	"backend-dinakom/app/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type UserMealController interface {
	GetAllUserMeals(c *fiber.Ctx) error
	GetTodayUserMeals(c *fiber.Ctx) error
}

type userMealController struct {
	userMealService services.UserMealService
}

func NewUserMealController(userMealService services.UserMealService) UserMealController {
	return &userMealController{userMealService: userMealService}
}

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

func (ctrl *userMealController) GetTodayUserMeals(c *fiber.Ctx) error {
	UserID := c.Locals("user_id").(string)
	data, err := ctrl.userMealService.GetTodayUserMeals(UserID)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return helpers.SuccessResponse(c, "today user meals retrieved successfully", data)
}

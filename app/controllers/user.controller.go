package controllers

import (
	"backend-dinakom/app/dto/request"
	"backend-dinakom/app/helpers"
	"backend-dinakom/app/services"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type UserController interface {
	GetAllUsers(c *fiber.Ctx) error
	GetUserSession(c *fiber.Ctx) error
	GetUserByID(c *fiber.Ctx) error
	CreateUser(c *fiber.Ctx) error
	UpdateUser(c *fiber.Ctx) error
	DeleteUser(c *fiber.Ctx) error
	ChangePassword(c *fiber.Ctx) error
	GetUserByRoleName(c *fiber.Ctx) error
	GetTotalUsers(c *fiber.Ctx) error
}

type userController struct {
	userService services.UserService
}

func NewUserController(userService services.UserService) UserController {
	return &userController{userService: userService}
}

func (ctrl *userController) GetAllUsers(c *fiber.Ctx) error {
	// Get pagination parameters
	page, _ := strconv.Atoi(c.Query("page", "1"))
	pageSize, _ := strconv.Atoi(c.Query("page_size", "10"))

	// Validate pagination
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}

	users, totalRows, err := ctrl.userService.GetAllUsers(page, pageSize)
	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, "Failed to fetch users")
	}

	return helpers.PaginationResponse(c, "Users retrieved successfully", users, page, pageSize, totalRows)
}

func (ctrl *userController) GetUserSession(c *fiber.Ctx) error {

	userId := c.Locals("user_id").(string)
	if userId == "" {
		return helpers.ErrorResponse(c, http.StatusUnauthorized, "user_id not found", nil)
	}

	user, err := ctrl.userService.GetUserSession(userId)
	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, "Failed to fetch session")
	}

	return helpers.SuccessResponse(c, "Session retrieved successfully", user)
}

func (ctrl *userController) GetUserByID(c *fiber.Ctx) error {
	id := c.Params("id")

	user, err := ctrl.userService.GetUserByID(id)
	if err != nil {
		return fiber.NewError(http.StatusNotFound, "User not found")
	}

	return helpers.SuccessResponse(c, "User retrieved successfully", user)
}

func (ctrl *userController) CreateUser(c *fiber.Ctx) error {
	var req request.CreateUserRequest

	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Validation error")
	}

	user, err := ctrl.userService.CreateUser(req)
	if err != nil {
		return fiber.NewError(http.StatusBadRequest, "Failed to create user", err.Error())
	}

	return helpers.CreatedResponse(c, "User created successfully", user)
}

func (ctrl *userController) UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")

	var req request.UpdateUserRequest
	if err := c.BodyParser(&req); err != nil {
		helpers.ValidationErrorResponse(c, err.Error())
	}

	user, err := ctrl.userService.UpdateUser(id, req)
	if err != nil {
		return fiber.NewError(http.StatusBadRequest, "Failed to update user")
	}

	return helpers.SuccessResponse(c, "User updated successfully", user)
}

func (ctrl *userController) DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := ctrl.userService.DeleteUser(id); err != nil {
		return fiber.NewError(http.StatusBadRequest, "Failed to delete user")
	}

	return helpers.SuccessResponse(c, "User deleted successfully", nil)
}

func (ctrl *userController) ChangePassword(c *fiber.Ctx) error {
	// Get user ID from JWT token (set by auth middleware)
	userID := c.Get("user_id")
	if userID == "" {
		return fiber.NewError(http.StatusUnauthorized, "Unauthorized")
	}

	var req request.ChangePasswordRequest
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Validation error")

	}

	if err := ctrl.userService.ChangePassword(userID, req); err != nil {
		return fiber.NewError(http.StatusBadRequest, "Failed to change password")
	}

	return helpers.SuccessResponse(c, "Password changed successfully", nil)
}

func (ctrl *userController) GetUserByRoleName(c *fiber.Ctx) error {
	roleName := c.Query("role_name")
	users, err := ctrl.userService.GetUserByRoleName(roleName)
	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, "Failed to fetch users", err.Error())
	}

	return helpers.SuccessResponse(c, "Users retrieved successfully", users)
}

func (ctrl *userController) GetTotalUsers(c *fiber.Ctx) error {
	total, err := ctrl.userService.GetTotalUsers()
	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, "Failed to fetch total users")
	}

	return helpers.SuccessResponse(c, "Total users retrieved successfully", total)
}

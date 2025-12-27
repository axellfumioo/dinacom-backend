package controllers

import (
	"backend-dinakom/app/dto/request"
	"backend-dinakom/app/helpers"
	"backend-dinakom/app/services"
	"backend-dinakom/configs"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type AuthController interface {
	Register(c *fiber.Ctx) error
	Login(c *fiber.Ctx) error
	StravaRedirect(c *fiber.Ctx) error
	StravaCallback(c *fiber.Ctx) error
}

type authController struct {
	authService services.AuthService
}

func NewAuthController(authService services.AuthService) AuthController {
	return &authController{authService: authService}
}

// Register
func (ctrl *authController) Register(c *fiber.Ctx) error {
	var req request.RegisterRequest

	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	data, err := ctrl.authService.Register(req)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return helpers.CreatedResponse(c, "user created successfully", data)
}

// Login
func (ctrl *authController) Login(c *fiber.Ctx) error {
	var req request.LoginRequest

	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.ErrBadRequest.Code, err.Error())
	}

	access_token, err := ctrl.authService.Login(req)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return helpers.SuccessResponse(c, "login successfully", access_token)
}

func (ctrl *authController) StravaRedirect(c *fiber.Ctx) error {
	clientID := configs.AppConfig.Strava.CLIENT_ID
	redirect := "http://localhost:8080/api/v1/auth/strava/callback"
	url := fmt.Sprintf(
		"https://www.strava.com/oauth/authorize?client_id=%s&response_type=code&redirect_uri=%s&approval_prompt=auto&scope=read,activity:read_all",
		clientID,
		redirect,
	)

	return c.Redirect(url)
}

func (ctrl *authController) StravaCallback(c *fiber.Ctx) error {
	clientId := configs.AppConfig.Strava.CLIENT_ID
	clientSecret := configs.AppConfig.Strava.CLIENT_KEY
	code := c.Query("code")
	if code == "" {
		return c.Status(400).JSON(fiber.Map{
			"error": "missing code from strava",
		})
	}
	resp, err := ctrl.authService.StravaCallbackHandle(clientId, clientSecret, code)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	// Redirect ke frontend

	return c.JSON(fiber.Map{
		"access_token": resp,
	})
}

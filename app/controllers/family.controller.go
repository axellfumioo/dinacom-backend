package controllers

import (
	"backend-dinakom/app/dto/request"
	"backend-dinakom/app/helpers"
	"backend-dinakom/app/services"

	"github.com/gofiber/fiber/v2"
)

type FamilyController interface {
	CreateNewFamily(c *fiber.Ctx) error
}

type familyController struct {
	familyService services.FamilyService
}

func NewFamilyController(familyService services.FamilyService) FamilyController {
	return &familyController{familyService: familyService}
}

func (ctrl *familyController) CreateNewFamily(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(string)
	// Body
	name := c.FormValue("name")
	desc := c.FormValue("description")

	// File handler
	image, err := c.FormFile("image")
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	data, err := ctrl.familyService.CreateNewFamily(userID, request.CreateFamilyRequest{Name: name, Description: desc, FamilyAvatar: image})
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return helpers.CreatedResponse(c, "family created successfully", data)
}

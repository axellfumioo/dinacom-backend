package controllers

import (
	"backend-dinakom/app/dto/request"
	"backend-dinakom/app/helpers"
	"backend-dinakom/app/services"

	"github.com/gofiber/fiber/v2"
)

type FamilyController interface {
	CreateNewFamily(c *fiber.Ctx) error
	UpdateFamilyAvatar(c *fiber.Ctx) error
	UpdateFamily(c *fiber.Ctx) error
	DeleteFamily(c *fiber.Ctx) error
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
		return fiber.NewError(fiber.StatusBadRequest, "validation error:"+err.Error())
	}
	data, err := ctrl.familyService.CreateNewFamily(userID, request.CreateFamilyRequest{Name: name, Description: desc, FamilyAvatar: image})
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return helpers.CreatedResponse(c, "family created successfully", data)
}

func (ctrl *familyController) UpdateFamily(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(string)
	familyID := c.Params("id")

	var req request.UpdateFamilyRequest
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "validation error:"+err.Error())
	}

	data, err := ctrl.familyService.UpdateFamily(familyID, userID, req)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return helpers.SuccessResponse(c, "family updated successfully", data)
}

func (ctrl *familyController) UpdateFamilyAvatar(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(string)
	familyID := c.Params("id")

	avatar, err := c.FormFile("avatar")
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "validation error:"+err.Error())
	}

	data, err := ctrl.familyService.UpdateFamilyAvatar(familyID, userID, avatar)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return helpers.SuccessResponse(c, "family avatar updated successfully", data)
}

func (ctrl *familyController) DeleteFamily(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(string)
	familyID := c.Params("id")

	data, err := ctrl.familyService.DeleteFamily(familyID, userID)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return helpers.SuccessResponse(c, "family avatar deleted successfully", data)
}

package controllers

import (
	"backend-dinakom/app/dto/request"
	"backend-dinakom/app/helpers"
	"backend-dinakom/app/services"

	"github.com/gofiber/fiber/v2"
)

type FamilyController interface {
	GetFamilyByID(c *fiber.Ctx) error
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

// GetFamilyByID godoc
// @Summary GetFamilyByID
// @Description endpoint to get family by familyID
// @Tags Families
// @Produce  json
// @Param ID path string true "id"
// @Success 201 {object} response.SuccessResponse
// @Failure 400 {object} response.ErrorResponse
// @Failure 404 {object} response.ErrorResponse
// @Router /families/{id} [get]
// @Security BearerAuth
func (ctrl *familyController) GetFamilyByID(c *fiber.Ctx) error {
	ID := c.Params("id")
	UserID := c.Locals("user_id").(string)

	data, err := ctrl.familyService.GetFamilyByID(ID, UserID)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return helpers.SuccessResponse(c, "family retrieved successfully", data)
}

// CreatenewFamily godoc
// @Summary CreatenewFamily
// @Description endpoint to create new family
// @Tags Families
// @Produce  json
// @Param name formData string true "enter family name"
// @Param description formData string true "family description"
// @Param image formData file true "upload family avatar"
// @Success 201 {object} response.SuccessResponse
// @Failure 400 {object} response.ErrorResponse
// @Failure 404 {object} response.ErrorResponse
// @Router /families [post]
// @Security BearerAuth
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

// UpdateFamily godoc
// @Summary UpdateFamily
// @Description endpoint to update family
// @Tags Families
// @Produce  json
// @Param request body request.UpdateFamilyRequest true "Update family body"
// @Success 201 {object} response.SuccessResponse
// @Failure 400 {object} response.ErrorResponse
// @Failure 404 {object} response.ErrorResponse
// @Router /families/{id} [patch]
// @Security BearerAuth
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

// UpdateFamilyAvatar godoc
// @Summary UpdateFamilyAvatar
// @Description endpoint to update family avatar
// @Tags Families
// @Produce  json
// @Param avatar formData file true "upload family avatar"
// @Success 201 {object} response.SuccessResponse
// @Failure 400 {object} response.ErrorResponse
// @Failure 404 {object} response.ErrorResponse
// @Router /families/{id}/avatar [patch]
// @Security BearerAuth
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

// Delete Family godoc
// @Summary DeleteFamily
// @Description endpoint to delete family
// @Tags Families
// @Produce  json
// @Param ID path string true "id"
// @Success 201 {object} response.SuccessResponse
// @Failure 400 {object} response.ErrorResponse
// @Failure 404 {object} response.ErrorResponse
// @Router /families/{id}/delete [delete]
// @Security BearerAuth
func (ctrl *familyController) DeleteFamily(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(string)
	familyID := c.Params("id")

	data, err := ctrl.familyService.DeleteFamily(familyID, userID)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return helpers.SuccessResponse(c, "family avatar deleted successfully", data)
}

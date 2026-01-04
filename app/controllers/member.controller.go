package controllers

import (
	"backend-dinakom/app/dto/request"
	"backend-dinakom/app/helpers"
	"backend-dinakom/app/services"

	"github.com/gofiber/fiber/v2"
)

type MemberController interface {
	GetFamilyMembers(c *fiber.Ctx) error
	AddFamilyMembers(c *fiber.Ctx) error
}

type memberController struct {
	memberService services.MemberService
}

func NewMemberService(memberService services.MemberService) MemberController {
	return &memberController{memberService: memberService}
}

// GetFamilyMembers godoc
// @Summary GetFamilyMembers
// @Description endpoint to get family members
// @Tags Members
// @Produce  json
// @Param familyID path string true "familyID"
// @Success 201 {object} response.SuccessResponse
// @Failure 400 {object} response.ErrorResponse
// @Failure 404 {object} response.ErrorResponse
// @Router /members/family/{familyID} [get]
// @Security BearerAuth
func (ctrl *memberController) GetFamilyMembers(c *fiber.Ctx) error {
	familyID := c.Params("familyID")
	data, err := ctrl.memberService.GetFamilyMembers(familyID)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return helpers.SuccessResponse(c, "family members retrieved successfully", data)
}

// AddFamilyMembers godoc
// @Summary AddFamilyMembers
// @Description endpoint to add family members
// @Tags Members
// @Produce  json
// @Param request body request.AddFamilyMemberRequest true "add family member body"
// @Success 201 {object} response.SuccessResponse
// @Failure 400 {object} response.ErrorResponse
// @Failure 404 {object} response.ErrorResponse
// @Router /members/ [post]
// @Security BearerAuth
func (ctrl *memberController) AddFamilyMembers(c *fiber.Ctx) error {
	var req request.AddFamilyMemberRequest
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "validation error:"+err.Error())
	}

	data, err := ctrl.memberService.AddFamilyMembers(req)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return helpers.CreatedResponse(c, "family members added successfully", data)
}

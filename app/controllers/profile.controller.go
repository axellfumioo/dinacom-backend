package controllers

import (
	"backend-dinakom/app/dto/request"
	"backend-dinakom/app/helpers"
	"backend-dinakom/app/services"

	"github.com/gofiber/fiber/v2"
)

type ProfileController interface {
	GetUserProfile(c *fiber.Ctx) error
	UpdateProfile(c *fiber.Ctx) error
	UploadAvatar(c *fiber.Ctx) error
}

type profileController struct {
	profileService services.ProfileService
}

func NewProfileController(profileService services.ProfileService) ProfileController {
	return &profileController{profileService: profileService}
}

func (ctrl *profileController) GetUserProfile(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(string)
	profile, err := ctrl.profileService.GetUserProfile(userID)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return helpers.SuccessResponse(c, "profile retrieved successfully", profile)
}

// Update Profile godoc
// @Summary Update Profile
// @Description Access this endpoint to update User profile
// @Tags Profiles
// @Accept json
// @Produce  json
// @Param request body request.UpdateProfileRequest true "Update Profile Body"
// @Success 201 {object} response.SuccessResponse
// @Failure 400 {object} response.ErrorResponse
// @Failure 404 {object} response.ErrorResponse
// @Router /profiles [patch]
// @Security BearerAuth
func (ctrl *profileController) UpdateProfile(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(string)

	var req request.UpdateProfileRequest
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	profile, err := ctrl.profileService.UpdateProfile(userID, req)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return helpers.SuccessResponse(c, "profile updated successfully", profile)
}

// Upload Profile Avatar godoc
// @Summary Upload Profile Avatar
// @Description Access this endpoint to upload user avatar (kirim image di body dengan nama field "avatar")
// @Param avatar formData file true "upload profile avatar"
// @Tags Profiles
// @Accept json
// @Produce  json
// @Success 201 {object} response.SuccessResponse
// @Failure 400 {object} response.ErrorResponse
// @Failure 404 {object} response.ErrorResponse
// @Router /profiles/avatar [post]
// @Security BearerAuth
func (ctrl *profileController) UploadAvatar(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(string)

	file, err := c.FormFile("avatar")
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	avatar, err := ctrl.profileService.UploadAvatar(request.UploadAvatarRequest{
		UserID: userID,
		Avatar: file,
	})
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return helpers.SuccessResponse(c, "avatar uploaded successfully", avatar)
}

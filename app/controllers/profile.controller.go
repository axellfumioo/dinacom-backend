package controllers

import (
	"backend-dinakom/app/dto/request"
	"backend-dinakom/app/helpers"
	"backend-dinakom/app/services"

	"github.com/gofiber/fiber/v2"
)

type ProfileController interface {
	UploadAvatar(c *fiber.Ctx) error
}

type profileController struct {
	profileService services.ProfileService
}

func NewProfileController(profileService services.ProfileService) ProfileController {
	return &profileController{profileService: profileService}
}

func (ctrl *profileController) UploadAvatar(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(string)

	file, err := c.FormFile("avatar")
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	avatar, err := ctrl.profileService.UpdateAvatar(request.UploadAvatarRequest{
		UserID: userID,
		Avatar: file,
	})

	return helpers.CreatedResponse(c, "avatar uploaded successfully", avatar)
}

package controllers

import (
	"backend-dinakom/app/dto/request"
	"backend-dinakom/app/helpers"
	"backend-dinakom/app/services"

	"github.com/gofiber/fiber/v2"
)

type QuestionnaireController interface {
	GetUserQuestionnaires(c *fiber.Ctx) error
	UpdateQuestionnaires(c *fiber.Ctx) error
}

type questionnaireController struct {
	questionnaireService services.QuestionnaireService
}

func NewQuestionnaireController(questionnaireService services.QuestionnaireService) QuestionnaireController {
	return &questionnaireController{questionnaireService: questionnaireService}
}

// GetUserQuestionnaires godoc
// @Summary GetUserQuestionnaires
// @Description endpoint to get user questionnaires
// @Tags Quest
// @Produce  json
// @Success 201 {object} response.SuccessResponse
// @Failure 400 {object} response.ErrorResponse
// @Failure 404 {object} response.ErrorResponse
// @Router /quest [get]
// @Security BearerAuth
func (ctrl *questionnaireController) GetUserQuestionnaires(c *fiber.Ctx) error {
	UserID := c.Locals("user_id").(string)
	data, err := ctrl.questionnaireService.GetUserQuestionnaires(UserID)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return helpers.SuccessResponse(c, "questionnaires retrieved successfully", data)
}

func (ctrl *questionnaireController) UpdateQuestionnaires(c *fiber.Ctx) error {
	UserID := c.Locals("user_id").(string)
	var req request.UpdateQuestionnairesRequest
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	err := ctrl.questionnaireService.UpdateQuestionnaires(UserID, &req)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return helpers.SuccessResponse(c, "questionnaires updated successfully", nil)
}

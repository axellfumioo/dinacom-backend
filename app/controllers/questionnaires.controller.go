package controllers

import (
	"backend-dinakom/app/helpers"
	"backend-dinakom/app/services"

	"github.com/gofiber/fiber/v2"
)

type QuestionnaireController interface {
	GetUserQuestionnaires(c *fiber.Ctx) error
}

type questionnaireController struct {
	questionnaireService services.QuestionnaireService
}

func newQuestionnaireController(questionnaireService services.QuestionnaireService) QuestionnaireController {
	return &questionnaireController{questionnaireService: questionnaireService}
}

func (ctrl *questionnaireController) GetUserQuestionnaires(c *fiber.Ctx) error {
	UserID := c.Locals("user_id").(string)
	data, err := ctrl.questionnaireService.GetUserQuestionnaires(UserID)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return helpers.SuccessResponse(c, "questionnaires retrieved successfully", data)
}

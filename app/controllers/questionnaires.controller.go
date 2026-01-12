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
// @Param userId query string true "Isi ini pake global state user_id" default("")
// @Success 201 {object} response.SuccessResponse
// @Failure 400 {object} response.ErrorResponse
// @Failure 404 {object} response.ErrorResponse
// @Router /quest [get]
func (ctrl *questionnaireController) GetUserQuestionnaires(c *fiber.Ctx) error {
	UserID := c.Query("userId", "")
	data, err := ctrl.questionnaireService.GetUserQuestionnaires(UserID)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return helpers.SuccessResponse(c, "questionnaires retrieved successfully", data)
}

// UpdateQuestionnaires godoc
// @Summary UpdateQuestionnaires
// @Description endpoint to update questionnaires (Ini endpoint untuk jawaban)
// @Tags Quest
// @Produce  json
// @Param userId query string true "Isi ini pake global state userId" default("")
// @Param request body request.UpdateQuestionnairesRequest true "Update Question Body"
// @Success 201 {object} response.SuccessResponse
// @Failure 400 {object} response.ErrorResponse
// @Failure 404 {object} response.ErrorResponse
// @Router /quest/answer [patch]
func (ctrl *questionnaireController) UpdateQuestionnaires(c *fiber.Ctx) error {
	UserID := c.Query("userId", "")
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

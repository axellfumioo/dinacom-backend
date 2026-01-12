package router

import (
	"backend-dinakom/app/controllers"
	"backend-dinakom/app/middlewares"
	"backend-dinakom/app/services"
	"backend-dinakom/configs"
	"backend-dinakom/database"

	"github.com/gofiber/fiber/v2"
)

func QuestionRoute(r fiber.Router) {
	questions := r.Group("quest")

	questionService := services.NewQuestionnaireService(database.DB, configs.QueueClient)
	questionController := controllers.NewQuestionnaireController(questionService)

	questions.Use(middlewares.AuthMiddleware())
	questions.Get("/", questionController.GetUserQuestionnaires)
	questions.Patch("/answer", questionController.UpdateQuestionnaires)
}

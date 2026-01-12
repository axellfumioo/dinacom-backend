package services

import (
	"backend-dinakom/app/models"
	"errors"

	"github.com/hibiken/asynq"
	"gorm.io/gorm"
)

type QuestionnaireService interface {
	GetUserQuestionnaires(UserID string) ([]models.Questionnaire, error)
}

type questionnaireService struct {
	db          *gorm.DB
	asynqClient *asynq.Client
}

func newQuestionnaireService(db *gorm.DB, asynqClient *asynq.Client) QuestionnaireService {
	return &questionnaireService{db: db, asynqClient: asynqClient}
}

func (s *questionnaireService) GetUserQuestionnaires(UserID string) ([]models.Questionnaire, error) {
	var existingUser models.User
	if err := s.db.First(&existingUser, "id = ?", UserID).Error; err != nil {
		return nil, errors.New("failed to get user: " + err.Error())
	}

	var questionnaires []models.Questionnaire
	if err := s.db.Where("user_id = ?", UserID).Find(&questionnaires).Error; err != nil {
		return nil, errors.New("failed to get questionnaires: " + err.Error())
	}

	if len(questionnaires) == 0 {
		return nil, errors.New("failed to get questionnaires: data not founds")
	}

	return questionnaires, nil
}

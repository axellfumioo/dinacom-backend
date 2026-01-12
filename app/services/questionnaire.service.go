package services

import (
	"backend-dinakom/app/dto/request"
	"backend-dinakom/app/models"
	"errors"
	"strconv"
	"strings"

	"github.com/hibiken/asynq"
	"gorm.io/gorm"
)

type QuestionnaireService interface {
	GetUserQuestionnaires(UserID string) ([]models.Questionnaire, error)
	UpdateQuestionnaires(UserID string, req *request.UpdateQuestionnairesRequest) error
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

func (s *questionnaireService) UpdateQuestionnaires(UserID string, req *request.UpdateQuestionnairesRequest) error {
	var isSmooking bool
	var sportDuration string
	var sleepDuration int

	tx := s.db.Begin()
	for _, answer := range req.Answers {
		if answer.Number == 2 {
			isSmooking = strings.ToLower(answer.Answer) == "ya"
		}

		if answer.Number == 3 {
			val, err := strconv.Atoi(answer.Answer)
			if err != nil {
				tx.Rollback()
				return errors.New("invalid sleep duration")
			}
			sleepDuration = val
		}

		if answer.Number == 4 {
			sportDuration = answer.Answer
		}

		if err := tx.Model(&models.Questionnaire{}).Where("id = ? AND user_id = ?", answer.QuestionID, UserID).Update("Answer", answer.Answer).Error; err != nil {
			tx.Rollback()
			return errors.New("failed to update questionnaires: " + err.Error())
		}
	}

	_ = isSmooking
	_ = sleepDuration
	_ = sportDuration

	return tx.Commit().Error
}

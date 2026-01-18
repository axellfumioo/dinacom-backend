package helpers

import (
	"backend-dinakom/app/models"
	"errors"
	"strconv"
	"strings"

	"gorm.io/gorm"
)

// Return type
type QuestionnaireAnswer struct {
	IsSmoking     *bool
	SleepDuration *int
	SportDuration *string
}

func GetQuestionnaireAnswer(db *gorm.DB, userID string) (*QuestionnaireAnswer, error) {
	var isSmooking *bool
	var sportDuration *string
	var sleepDuration *int

	// Number of questions
	const (
		QSmoking = 2
		QSleep   = 3
		QSport   = 4
	)

	var questions []models.Questionnaire
	if err := db.Where("user_id = ? AND number IN (?, ?, ?)", userID, QSmoking, QSleep, QSport).Order("number ASC").Find(&questions).Error; err != nil {
		return nil, errors.New("failed to get questionnaire: " + err.Error())
	}

	for _, question := range questions {
		if question.Number == QSmoking {
			val := strings.ToLower(*question.Answer) == "ya"
			isSmooking = &val
		}

		if question.Number == QSleep {
			val, err := strconv.Atoi(*question.Answer)
			if err != nil {
				return nil, errors.New("invalid sleep duration")
			}
			sleepDuration = &val
		}

		if question.Number == QSport {
			sportDuration = question.Answer
		}
	}

	result := QuestionnaireAnswer{
		IsSmoking:     isSmooking,
		SleepDuration: sleepDuration,
		SportDuration: sportDuration,
	}
	return &result, nil
}

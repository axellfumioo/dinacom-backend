package services

import (
	"backend-dinakom/app/dto/payload"
	"backend-dinakom/app/dto/request"
	"backend-dinakom/app/dto/response"
	"backend-dinakom/app/helpers"
	"backend-dinakom/app/models"
	"encoding/json"
	"errors"
	"time"

	"github.com/hibiken/asynq"
	"gorm.io/gorm"
)

type UserMealService interface {
	GetAllUserMeals(userID string, page int, pageSize int) ([]response.UserMealResponse, int64, error)
	GetTodayUserMeals(userID string) ([]response.UserMealResponse, error)
	CreateUserMeal(userID string, req request.CreateUserMealRequest) (*response.UserMealResponse, error)
}

type userMealService struct {
	db          *gorm.DB
	asynqClient *asynq.Client
}

func NewUserMealService(db *gorm.DB, asynqClient *asynq.Client) UserMealService {
	return &userMealService{db: db, asynqClient: asynqClient}
}

func (s *userMealService) GetAllUserMeals(userID string, page int, pageSize int) ([]response.UserMealResponse, int64, error) {
	// Check if user exist
	var existingUsers models.User
	if err := s.db.First(&existingUsers, "id = ?", userID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, 0, errors.New("user is not registered")
		}

		return nil, 0, err
	}

	var totalRows int64
	if err := s.db.Model(&models.UserMeal{}).Count(&totalRows).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * pageSize

	var userMeals []models.UserMeal
	if err := s.db.Preload("User").Offset(offset).Limit(pageSize).Order("created_at DESC").Find(&userMeals).Error; err != nil {
		return nil, 0, errors.New("failed to get foodScan history")
	}

	userMealsResponse := helpers.ToUserMealsResponse(userMeals)
	return userMealsResponse, totalRows, nil
}

func (s *userMealService) GetTodayUserMeals(UserID string) ([]response.UserMealResponse, error) {
	today := time.Now()
	startOfDay := time.Date(today.Year(), today.Month(), today.Day(), 0, 0, 0, 0, today.Location())
	endOfDay := startOfDay.Add(24 * time.Hour)

	var userMeals []models.UserMeal
	if err := s.db.Where("user_id = ? AND created_at >= ? AND created_at < ?", UserID, startOfDay, endOfDay).Find(&userMeals).Error; err != nil {
		return nil, err
	}

	if len(userMeals) == 0 {
		return nil, errors.New("no meals found today for this user")
	}

	userMealsResponse := helpers.ToUserMealsResponse(userMeals)
	return userMealsResponse, nil
}

func (s *userMealService) CreateUserMeal(userID string, req request.CreateUserMealRequest) (*response.UserMealResponse, error) {
	var existingUser models.User
	if err := s.db.Preload("Profile").First(&existingUser, "user_id = ?", userID).Error; err != nil {
		return nil, errors.New("failed to find user: " + err.Error())
	}

	userMeal := models.UserMeal{
		FoodName: req.FoodName,
		Calories: req.Calories,
		Protein:  req.Protein,
		Fat:      req.Fat,
		Carbs:    req.Carbs,
		Portion:  req.Portion,
		Time:     req.Time,
		UserID:   userID,
	}
	if err := s.db.Create(&userMeal).Error; err != nil {
		return nil, errors.New("failed to create usermeal: " + err.Error())
	}

	// Ini untuk AI insight
	answer, err := helpers.GetQuestionnaireAnswer(s.db, userID)
	if err != nil {
		return nil, err
	}
	userAge := helpers.CalculateAge(existingUser.Profile.DateOfBirth)
	payload := payload.AIInsightPayload{
		User: payload.UserAIInsight{
			ID:            userID,
			Name:          existingUser.FullName,
			Age:           userAge,
			Gender:        existingUser.Profile.Gender,
			Smoking:       *answer.IsSmoking,
			SleepDuration: *answer.SleepDuration,
			SportDuration: *answer.SportDuration,
		},
		DailyNutritionSummary: payload.DailyNutritionSummary{
			CaloriesKcal: userMeal.Calories,
			Nutrition: payload.DailyNutrition{
				CarbsG:   userMeal.Carbs,
				ProteinG: userMeal.Protein,
				FatG:     userMeal.Fat,
			},
			Vitamins: userMeal.Vitamins,
		},
	}

	b, _ := json.Marshal(payload)
	task := asynq.NewTask("ai-insight:process", b)
	s.asynqClient.Enqueue(task, asynq.MaxRetry(1))

	userMealResponse := helpers.ToUserMealResponse(&userMeal)
	return &userMealResponse, nil
}

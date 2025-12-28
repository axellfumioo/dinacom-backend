package services

import (
	"backend-dinakom/app/dto/response"
	"backend-dinakom/app/helpers"
	"backend-dinakom/app/models"
	"errors"
	"time"

	"gorm.io/gorm"
)

type UserMealService interface {
	GetAllUserMeals(userID string, page int, pageSize int) ([]response.UserMealResponse, int64, error)
	GetTodayUserMeals(userID string) ([]response.UserMealResponse, error)
}

type userMealService struct {
	db *gorm.DB
}

func NewUserMealService(db *gorm.DB) UserMealService {
	return &userMealService{db: db}
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

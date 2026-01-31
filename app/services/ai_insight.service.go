package services

import (
	"backend-dinakom/app/models"

	"gorm.io/gorm"
)

type AIInsightService interface {
	GetLatestInsight(userID string) (*models.AIInsight, error)
}

type aiInsightService struct {
	db *gorm.DB
}

func NewAIInsightService(db *gorm.DB) AIInsightService {
	return &aiInsightService{db: db}
}

func (s *aiInsightService) GetLatestInsight(userID string) (*models.AIInsight, error) {
	var existingUser models.User
	if err := s.db.First(&existingUser, "id = ?", userID).Error; err != nil {
		return nil, err
	}

	var latestInsight models.AIInsight
	if err := s.db.Where("user_id = ?", userID).Order("created_at DESC").First(&latestInsight).Error; err != nil {
		return nil, err
	}

	return &latestInsight, nil
}

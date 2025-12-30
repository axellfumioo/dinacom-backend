package services

import (
	"backend-dinakom/app/models"
	"errors"

	"gorm.io/gorm"
)

type AIChatService interface {
	CreateNewChat(UserID string) (*models.AiChat, error)
}

type aIChatService struct {
	db *gorm.DB
}

func NewAIChatService(db *gorm.DB) AIChatService {
	return &aIChatService{db: db}
}

func (s *aIChatService) CreateNewChat(UserID string) (*models.AiChat, error) {
	var exist models.AiChat
	if err := s.db.First(&exist, "user_id = ?", UserID).Error; err == nil {
		return nil, errors.New("chat is already exist")
	}

	AiChat := &models.AiChat{
		UserID: UserID,
	}

	if err := s.db.Create(&AiChat).Error; err != nil {
		return nil, err
	}

	return AiChat, nil
}

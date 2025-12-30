package services

import (
	"backend-dinakom/app/dto/request"
	"backend-dinakom/app/models"
	"errors"

	"github.com/hibiken/asynq"
	"gorm.io/gorm"
)

type AIChatMessageService interface {
	CreateNewMessage(req request.CreateMessageRequest, UserID string, ChatID string) (*models.AIChatMessage, error)
}

type aiChatMessageService struct {
	db           *gorm.DB
	asyncqClient *asynq.Client
}

func NewAIChatMessageService(db *gorm.DB, asynqClient *asynq.Client) AIChatMessageService {
	return &aiChatMessageService{db: db, asyncqClient: asynqClient}
}

func (s *aiChatMessageService) CreateNewMessage(req request.CreateMessageRequest, UserID string, ChatID string) (*models.AIChatMessage, error) {
	var existingChat models.AiChat
	if err := s.db.First(&existingChat, "id = ?", ChatID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("this chat doesn't exist")
		}

		return nil, err
	}

	ChatMessage := &models.AIChatMessage{
		Content: req.Content,
		ChatID:  ChatID,
		UserID:  &UserID,
	}

	if err := s.db.Create(&ChatMessage).Error; err != nil {
		return nil, err
	}

	return ChatMessage, nil
}

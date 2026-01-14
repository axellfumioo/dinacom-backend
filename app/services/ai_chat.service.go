package services

import (
	"backend-dinakom/app/dto/response"
	"backend-dinakom/app/helpers"
	"backend-dinakom/app/models"
	"errors"

	"gorm.io/gorm"
)

type AIChatService interface {
	GetAIChatByID(ID string, userID string) (*response.AiChatResponse, error)
	GetUserAIChat(UserID string) (*response.AiChatResponse, error)
	CreateNewChat(UserID string) (*response.AiChatResponse, error)
	DeleteAIChat(ID string, UserID string) (*response.AiChatResponse, error)
}

type aIChatService struct {
	db *gorm.DB
}

func NewAIChatService(db *gorm.DB) AIChatService {
	return &aIChatService{db: db}
}

func (s *aIChatService) GetAIChatByID(ID string, UserID string) (*response.AiChatResponse, error) {
	var aiChat models.AiChat
	if err := s.db.Preload("Messages").Where("id = ? AND user_id = ?", ID, UserID).First(&aiChat).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("failed to get aichat: aichat not found")
		}
		return nil, errors.New("failed to get aiChat: " + err.Error())
	}

	aiChatResponse := helpers.ToAIChatResponse(&aiChat)
	return &aiChatResponse, nil
}

func (s *aIChatService) GetUserAIChat(UserID string) (*response.AiChatResponse, error) {
	var aiChat models.AiChat
	if err := s.db.Preload("Messages").Where("user_id = ?", UserID).First(&aiChat).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("ai chats not found")
		}
		return nil, err
	}

	aiChatResponse := helpers.ToAIChatResponse(&aiChat)
	return &aiChatResponse, nil
}

func (s *aIChatService) CreateNewChat(UserID string) (*response.AiChatResponse, error) {
	AiChat := &models.AiChat{
		UserID: UserID,
	}

	if err := s.db.Create(&AiChat).Error; err != nil {
		return nil, err
	}

	aiChatResponse := helpers.ToAIChatResponse(AiChat)
	return &aiChatResponse, nil
}

func (s *aIChatService) DeleteAIChat(ID string, UserID string) (*response.AiChatResponse, error) {
	var exist models.AiChat
	if err := s.db.Where("id = ? AND user_id = ?", ID, UserID).First(&exist).Error; err != nil {
		return nil, err
	}

	if err := s.db.Delete(&exist).Error; err != nil {
		return nil, err
	}

	aiChatResponse := helpers.ToAIChatResponse(&exist)
	return &aiChatResponse, nil
}

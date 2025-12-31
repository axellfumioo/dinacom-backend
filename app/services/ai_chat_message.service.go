package services

import (
	"backend-dinakom/app/dto/payload"
	"backend-dinakom/app/dto/request"
	"backend-dinakom/app/dto/response"
	"backend-dinakom/app/helpers"
	"backend-dinakom/app/models"

	"backend-dinakom/configs"
	"encoding/json"
	"errors"
	"fmt"
	"path/filepath"

	"github.com/hibiken/asynq"
	"github.com/minio/minio-go/v7"
	"gorm.io/gorm"
)

type AIChatMessageService interface {
	GetChatMessagesByChatID(ChatID string) ([]response.AIChatMessageResponse, error)
	CreateNewMessage(req request.CreateMessageRequest, UserID string, ChatID string) (*response.AIChatMessageResponse, error)
	CreateNewMessageWithImage(req request.CreateMessageWithMediaRequest) (*response.AIChatMessageResponse, error)
	DeleteChatMessageByID(ID string, UserID string) (*response.AIChatMessageResponse, error)
	DeleteChatMessages(ChatID string, UserID string) ([]response.AIChatMessageResponse, error)
}

type aiChatMessageService struct {
	db           *gorm.DB
	asyncqClient *asynq.Client
	minioClient  *minio.Client
}

func NewAIChatMessageService(db *gorm.DB, asynqClient *asynq.Client, minioClient *minio.Client) AIChatMessageService {
	return &aiChatMessageService{db: db, asyncqClient: asynqClient, minioClient: minioClient}
}

func (s *aiChatMessageService) GetChatMessagesByChatID(ChatID string) ([]response.AIChatMessageResponse, error) {
	var aiChat models.AiChat
	if err := s.db.First(&aiChat, "id = ?", ChatID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("chat not found")
		}
		return nil, err
	}

	var messages []models.AIChatMessage
	if err := s.db.Preload("User").Preload("Chat").Order("created_at ASC").Where("chat_id", ChatID).Find(&messages).Error; err != nil {
		return nil, err
	}

	messagesResponse := helpers.ToAIChatMessagesResponse(messages)
	return messagesResponse, nil
}

func (s *aiChatMessageService) CreateNewMessage(req request.CreateMessageRequest, UserID string, ChatID string) (*response.AIChatMessageResponse, error) {
	var existingChat models.AiChat
	if err := s.db.First(&existingChat, "id = ?", ChatID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("this chat doesn't exist")
		}

		return nil, err
	}

	var lastMessages []models.AIChatMessage
	s.db.Where("chat_id = ?", ChatID).Order("created_at DESC").Limit(5).Find(&lastMessages)

	ChatMessage := &models.AIChatMessage{
		Content: req.Content,
		ChatID:  ChatID,
		UserID:  &UserID,
	}

	if err := s.db.Create(&ChatMessage).Error; err != nil {
		return nil, err
	}

	aiChatMessageResponse := helpers.ToAIChatMessageResponse(ChatMessage)
	payload := payload.CreateAIMessagePayload{
		ChatID:      ChatID,
		UserID:      UserID,
		Message:     req.Content,
		ChatHistory: lastMessages,
	}

	b, _ := json.Marshal(payload)
	task := asynq.NewTask("aichat:process", b)
	s.asyncqClient.Enqueue(task)

	return &aiChatMessageResponse, nil
}

func (s *aiChatMessageService) CreateNewMessageWithImage(req request.CreateMessageWithMediaRequest) (*response.AIChatMessageResponse, error) {
	var existingChat models.AiChat
	if err := s.db.First(&existingChat, "id = ?", req.ChatID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("this chat doesn't exist")
		}

		return nil, err
	}

	// File handler
	ext := filepath.Ext(req.Image.Filename)
	object := fmt.Sprintf("messages/%s%s", req.UserID, ext)

	baseUrl := configs.AppConfig.Minio.BaseUrl
	bucket := configs.AppConfig.Minio.Bucket
	url, err := helpers.UploadFile(s.minioClient, req.Image, baseUrl, bucket, object)
	if err != nil {
		return nil, err
	}

	// Query last message untuk a.i
	var lastMessages []models.AIChatMessage
	s.db.Where("chat_id = ?", req.ChatID).Order("created_at DESC").Limit(5).Find(&lastMessages)

	// Create message handler
	ChatMessage := &models.AIChatMessage{
		Content:  req.Content,
		ChatID:   req.ChatID,
		UserID:   &req.UserID,
		ImageURL: &url,
	}

	if err := s.db.Create(&ChatMessage).Error; err != nil {
		return nil, err
	}

	aiChatMessageResponse := helpers.ToAIChatMessageResponse(ChatMessage)
	payload := payload.CreateAIMessagePayload{
		ChatID:      req.ChatID,
		UserID:      req.UserID,
		Message:     req.Content,
		ChatHistory: lastMessages,
	}

	b, _ := json.Marshal(payload)
	task := asynq.NewTask("aichat:process", b)
	s.asyncqClient.Enqueue(task)

	return &aiChatMessageResponse, nil
}

func (s *aiChatMessageService) DeleteChatMessageByID(ID string, UserID string) (*response.AIChatMessageResponse, error) {
	var exist models.AIChatMessage
	if err := s.db.Where("id = ? AND user_id = ?", ID, UserID).First(&exist).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("message not found")
		}
		return nil, err
	}

	if err := s.db.Delete(&exist).Error; err != nil {
		return nil, err
	}

	messageResponse := helpers.ToAIChatMessageResponse(&exist)
	return &messageResponse, nil
}

func (s *aiChatMessageService) DeleteChatMessages(ChatID string, UserID string) ([]response.AIChatMessageResponse, error) {
	var aiChat models.AiChat
	if err := s.db.Where("id = ? AND user_id = ?", ChatID, UserID).First(&aiChat).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("chat not found")
		}
		return nil, err
	}

	var messages []models.AIChatMessage
	if err := s.db.Where("chat_id = ?", ChatID).Find(&messages).Error; err != nil {
		return nil, err
	}

	if len(messages) == 0 {
		return nil, errors.New("messages don't exist in this chat")
	}

	if err := s.db.Delete(&messages).Error; err != nil {
		return nil, err
	}

	messagesResponse := helpers.ToAIChatMessagesResponse(messages)
	return messagesResponse, nil
}

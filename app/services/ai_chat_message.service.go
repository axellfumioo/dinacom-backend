package services

import (
	"backend-dinakom/app/dto/request"
	"backend-dinakom/app/helpers"
	"backend-dinakom/app/models"
	"backend-dinakom/configs"
	"errors"
	"fmt"
	"path/filepath"

	"github.com/hibiken/asynq"
	"github.com/minio/minio-go/v7"
	"gorm.io/gorm"
)

type AIChatMessageService interface {
	CreateNewMessage(req request.CreateMessageRequest, UserID string, ChatID string) (*models.AIChatMessage, error)
	CreateNewMessageWithImage(req request.CreateMessageWithMediaRequest) (*models.AIChatMessage, error)
}

type aiChatMessageService struct {
	db           *gorm.DB
	asyncqClient *asynq.Client
	minioClient  *minio.Client
}

func NewAIChatMessageService(db *gorm.DB, asynqClient *asynq.Client, minioClient *minio.Client) AIChatMessageService {
	return &aiChatMessageService{db: db, asyncqClient: asynqClient, minioClient: minioClient}
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

func (s *aiChatMessageService) CreateNewMessageWithImage(req request.CreateMessageWithMediaRequest) (*models.AIChatMessage, error) {
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

	return ChatMessage, nil
}

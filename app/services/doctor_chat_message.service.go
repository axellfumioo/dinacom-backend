package services

import (
	"backend-dinakom/app/dto/request"
	"backend-dinakom/app/models"
	"errors"

	"gorm.io/gorm"
)

type DoctorChatMessageService interface {
	GetMessagesByRoomID(RoomID string) ([]models.DoctorChatMessage, error)
	CreateNewMessage(senderID string, req request.CreateDoctorChatMessageRequest) (*models.DoctorChatMessage, error)
}

type doctorChatMessageService struct {
	db *gorm.DB
}

func NewDoctorChatMessageService(db *gorm.DB) DoctorChatMessageService {
	return &doctorChatMessageService{db: db}
}

func (s *doctorChatMessageService) GetMessagesByRoomID(RoomID string) ([]models.DoctorChatMessage, error) {
	var existingRoom models.DoctorChatRoom
	if err := s.db.Where("id = ?", RoomID).First(&existingRoom).Error; err != nil {
		return nil, err
	}

	var messages []models.DoctorChatMessage
	if err := s.db.Order("createdAt ASC").Where("doctor_chat_id = ?", RoomID).Find(&messages).Error; err != nil {
		return nil, err
	}

	return messages, nil
}

func (s *doctorChatMessageService) CreateNewMessage(senderID string, req request.CreateDoctorChatMessageRequest) (*models.DoctorChatMessage, error) {
	var existingRoom models.DoctorChatRoom
	if err := s.db.Where("id = ?", req.RoomID).First(&existingRoom).Error; err != nil {
		return nil, errors.New("failed to get chat room: " + err.Error())
	}

	var newMessage = models.DoctorChatMessage{
		DoctorChatID: existingRoom.ID,
		SenderID:     senderID,
		Message:      req.Message,
	}

	if err := s.db.Create(&newMessage).Error; err != nil {
		return nil, errors.New("failed to create chat message: " + err.Error())
	}

	return &newMessage, nil
}

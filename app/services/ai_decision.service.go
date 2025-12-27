package services

import (
	"backend-dinakom/app/dto/request"
	"backend-dinakom/app/models"
	"errors"

	"gorm.io/gorm"
)

type AIDecisionService interface {
	CreateDecision(req request.CreateDecisionRequest) (any, error)
}

type aIDecisionService struct {
	db *gorm.DB
}

func NewAIDecisionService(db *gorm.DB) AIDecisionService {
	return &aIDecisionService{db: db}
}

func (s *aIDecisionService) CreateDecision(req request.CreateDecisionRequest) (any, error) {
	var msg models.AIChatMessage
	if err := s.db.First(&msg, "id = ?", req.MessageId).Error; err != nil {
		return nil, err
	}

	var existingDecision models.AIDecision
	if err := s.db.First(&existingDecision, "chat_message_id = ?", req.MessageId).Error; err == nil {
		return nil, errors.New("AI decision in this message is already exist")
	}

	decision := &models.AIDecision{
		ChatMessageID: req.MessageId,
		Queries:       req.Queries,
		NeedSearch:    req.NeedSearch,
		RiskLevel:     req.RiskLevel,
		RequestType:   req.RequestType,
	}

	if err := s.db.Create(&decision).Error; err != nil {
		return nil, errors.New("failed to create ai decision")
	}

	return &decision, nil
}

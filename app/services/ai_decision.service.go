package services

import (
	"backend-dinakom/app/dto/request"
	"backend-dinakom/app/dto/response"
	"backend-dinakom/app/helpers"
	"backend-dinakom/app/models"
	"errors"

	"gorm.io/gorm"
)

type AIDecisionService interface {
	GetAllDecisions(page int, pageSize int) ([]response.AIDecisionResponse, int64, error)
	GetDecisionByID(ID string) (*response.AIDecisionResponse, error)
	CreateDecision(req request.CreateDecisionRequest) (any, error)
}

type aIDecisionService struct {
	db *gorm.DB
}

func NewAIDecisionService(db *gorm.DB) AIDecisionService {
	return &aIDecisionService{db: db}
}

func (s *aIDecisionService) GetAllDecisions(page int, pageSize int) ([]response.AIDecisionResponse, int64, error) {
	var totalRows int64
	if err := s.db.Model(&models.AIDecision{}).Count(&totalRows).Error; err != nil {
		return nil, 0, err
	}

	if totalRows == 0 {
		return nil, 0, errors.New("decisions not founds")
	}

	var decisions []models.AIDecision
	offset := (page - 1) * pageSize

	if err := s.db.Offset(offset).Limit(pageSize).Order("created_at DESC").Find(&decisions).Error; err != nil {
		return nil, 0, err
	}

	decisionsResponse := helpers.ToAIDecisionsResponse(decisions)
	return decisionsResponse, totalRows, nil
}

func (s *aIDecisionService) GetDecisionByID(ID string) (*response.AIDecisionResponse, error) {
	var existingDecision models.AIDecision
	if err := s.db.First(&existingDecision, "id = ?", ID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("decisions not found")
		}
		return nil, err
	}

	decisionResponse := helpers.ToAIDecisionResponse(&existingDecision)
	return &decisionResponse, nil
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

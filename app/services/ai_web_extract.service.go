package services

import (
	"backend-dinakom/app/dto/response"
	"backend-dinakom/app/helpers"
	"backend-dinakom/app/models"
	"errors"

	"gorm.io/gorm"
)

type AIWebExtractService interface {
	GetAllWebExtracts(page int, pageSize int) ([]response.AIWebExtractResponse, int64, error)
	GetWebExtractByID(ID string) (*response.AIWebExtractResponse, error)
	// CreateWebExtract(req request.CreateWebExtractRequest) (*response.AIWebExtractResponse, error)
	// UpdateWebExtract(ID string, req request.UpdateWebExtractRequest) (*response.AIWebExtractResponse, error)
	// DeleteWebExtract(ID string) (*response.AIWebExtractResponse, error)
}

type aIWebExtractService struct {
	db *gorm.DB
}

func NewAIWebExtractService(db *gorm.DB) AIWebExtractService {
	return &aIWebExtractService{db: db}
}

func (s *aIWebExtractService) GetAllWebExtracts(page int, pageSize int) ([]response.AIWebExtractResponse, int64, error) {
	var totalRows int64
	if err := s.db.Model(&models.AIDecision{}).Count(&totalRows).Error; err != nil {
		return nil, 0, err
	}

	if totalRows == 0 {
		return nil, 0, errors.New("decisions not founds")
	}

	var aiWebExtract []models.AIWebExtract
	offset := (page - 1) * pageSize

	if err := s.db.Offset(offset).Limit(pageSize).Order("created_at DESC").Find(&aiWebExtract).Error; err != nil {
		return nil, 0, err
	}

	aiWebExtractResponse := helpers.ToAIWebExtractsResponse(aiWebExtract)
	return aiWebExtractResponse, totalRows, nil
}

func (s *aIWebExtractService) GetWebExtractByID(ID string) (*response.AIWebExtractResponse, error) {
	var existing models.AIWebExtract
	if err := s.db.First(&existing, "id = ?", ID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("AI google search not found")
		}
		return nil, err
	}

	webExtractResponse := helpers.ToAIWebExtractResponse(&existing)
	return &webExtractResponse, nil
}
package services

import (
	"backend-dinakom/app/dto/response"
	"backend-dinakom/app/helpers"
	"backend-dinakom/app/models"
	"errors"

	"gorm.io/gorm"
)

type AIGoogleSearchService interface {
	GetAllGoogleSearchs(page int, pageSize int) ([]response.AIGoogleSearchResponse, int64, error)
	GetGoogleSearchByID(ID string) (*response.AIGoogleSearchResponse, error)
	// CreateGoogleSearch(req request.CreateGoogleSearchRequest) (any, error)
	// UpdateGoogleSearch(ID string, req request.UpdateGoogleSearchRequest) (*response.AIGoogleSearchResponse, error)
	// DeleteGoogleSearch(ID string) (*response.AIGoogleSearchResponse, error)
}

type aIGoogleSearchService struct {
	db *gorm.DB
}

func NewAIGoogleSearchService(db *gorm.DB) AIGoogleSearchService {
	return &aIGoogleSearchService{db: db}
}

func (s *aIGoogleSearchService) GetAllGoogleSearchs(page int, pageSize int) ([]response.AIGoogleSearchResponse, int64, error) {
	var totalRows int64
	if err := s.db.Model(&models.AIDecision{}).Count(&totalRows).Error; err != nil {
		return nil, 0, err
	}

	if totalRows == 0 {
		return nil, 0, errors.New("decisions not founds")
	}

	var aigoogleSearch []models.AIGoogleSearch
	offset := (page - 1) * pageSize

	if err := s.db.Offset(offset).Limit(pageSize).Order("created_at DESC").Find(&aigoogleSearch).Error; err != nil {
		return nil, 0, err
	}

	aiGoogleSearchResponse := helpers.ToAIGoogleSearchResponse(aigoogleSearch)
	return aiGoogleSearchResponse, totalRows, nil
	
	return nil, 0, nil
}

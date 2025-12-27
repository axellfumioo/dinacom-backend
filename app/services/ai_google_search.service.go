package services

import (
	"backend-dinakom/app/dto/request"
	"backend-dinakom/app/dto/response"
	"backend-dinakom/app/helpers"
	"backend-dinakom/app/models"
	"errors"

	"gorm.io/gorm"
)

type AIGoogleSearchService interface {
	GetAllGoogleSearchs(page int, pageSize int) ([]response.AIGoogleSearchResponse, int64, error)
	GetGoogleSearchByID(ID string) (*response.AIGoogleSearchResponse, error)
	CreateGoogleSearch(req request.CreateGoogleSearchRequest) (*response.AIGoogleSearchResponse, error)
	UpdateGoogleSearch(ID string, req request.UpdateGoogleSearchRequest) (*response.AIGoogleSearchResponse, error)
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

	aiGoogleSearchResponse := helpers.ToAIGoogleSearchsResponse(aigoogleSearch)
	return aiGoogleSearchResponse, totalRows, nil
}

func (s *aIGoogleSearchService) GetGoogleSearchByID(ID string) (*response.AIGoogleSearchResponse, error) {
	var existing models.AIGoogleSearch
	if err := s.db.First(&existing, "id = ?", ID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("AI google search not found")
		}
		return nil, err
	}

	googleSearchResponse := helpers.ToAIGoogleSearchResponse(&existing)
	return &googleSearchResponse, nil
}

func (s *aIGoogleSearchService) CreateGoogleSearch(req request.CreateGoogleSearchRequest) (*response.AIGoogleSearchResponse, error) {
	googleSearch := &models.AIGoogleSearch{
		URL:        req.URL,
		Content:    req.Content,
		DecisionID: req.DecisionID,
	}

	if err := s.db.Create(&googleSearch).Error; err != nil {
		return nil, errors.New("failed to create ai decision")
	}

	googleSearchResponse := helpers.ToAIGoogleSearchResponse(googleSearch)
	return &googleSearchResponse, nil
}

func (s *aIGoogleSearchService) UpdateGoogleSearch(ID string, req request.UpdateGoogleSearchRequest) (*response.AIGoogleSearchResponse, error) {
	var existing models.AIGoogleSearch
	if err := s.db.First(&existing, "id = ?", ID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("AI google search not found")
		}
		return nil, err
	}

	updateData := map[string]interface{}{}

	if req.URL != nil && req.URL != &existing.URL {
		updateData["url"] = req.URL
	}

	if req.Content != nil && req.Content != &existing.Content {
		updateData["content"] = req.Content
	}

	if err := s.db.Model(&existing).Where("id = ?", ID).Updates(updateData).Error; err != nil {
		return nil, err
	}

	googleSearchResponse := helpers.ToAIGoogleSearchResponse(&existing)
	return &googleSearchResponse, nil
}

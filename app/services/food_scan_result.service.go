package services

import (
	"backend-dinakom/app/dto/response"
	"backend-dinakom/app/helpers"
	"backend-dinakom/app/models"
	"errors"

	"gorm.io/gorm"
)

type FoodScanResultService interface {
	GetResultByID(id string) (*response.FoodScanResultResponse, error)
}

type foodScanResultService struct {
	db *gorm.DB
}

func NewFoodScanResultService(db *gorm.DB) FoodScanResultService {
	return &foodScanResultService{db: db}
}

func (s *foodScanResultService) GetResultByID(id string) (*response.FoodScanResultResponse, error) {
	var existingResult models.FoodScanResult
	if err := s.db.Preload("FoodScan").First(&existingResult, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("result not found")
		}
		return nil, err
	}

	resultResponse := helpers.ToFoodScanResultResponse(&existingResult)
	return &resultResponse, nil
}

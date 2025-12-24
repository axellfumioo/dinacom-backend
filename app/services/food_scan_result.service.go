package services

import (
	"backend-dinakom/app/dto/response"
	"backend-dinakom/app/helpers"
	"backend-dinakom/app/models"
	"errors"

	"gorm.io/gorm"
)

type FoodScanResultService interface {
	GetAllResults(page int, pageSize int) ([]response.FoodScanResultResponse, int64, error)
	GetAllUserResults(userID string, page int, pageSize int) ([]response.FoodScanResultResponse, int64, error)
	GetResultByID(id string) (*response.FoodScanResultResponse, error)
}

type foodScanResultService struct {
	db *gorm.DB
}

func NewFoodScanResultService(db *gorm.DB) FoodScanResultService {
	return &foodScanResultService{db: db}
}

func (s *foodScanResultService) GetAllResults(page int, pageSize int) ([]response.FoodScanResultResponse, int64, error) {
	var totalRows int64
	if err := s.db.Model(&models.FoodScanResult{}).Count(&totalRows).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * pageSize

	var fsResults []models.FoodScanResult
	if err := s.db.Preload("FoodScan").Offset(offset).Limit(pageSize).Order("created_at DESC").Find(&fsResults).Error; err != nil {
		return nil, 0, err
	}

	if len(fsResults) == 0 {
		return nil, 0, errors.New("foodscan results not found")
	}

	fsResultResponse := helpers.ToFoodScanResultsResponse(fsResults)
	return fsResultResponse, totalRows, nil
}

func (s *foodScanResultService) GetAllUserResults(userID string, page int, pageSize int) ([]response.FoodScanResultResponse, int64, error) {
	var totalRows int64
	if err := s.db.Model(&models.FoodScanResult{}).Count(&totalRows).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * pageSize

	var fsResults []models.FoodScanResult
	if err := s.db.Preload("FoodScan").Joins("JOIN food_scans s ON s.id = food_scan_result.food_scan_id").Where("s.user_id = ?", userID).Offset(offset).Limit(pageSize).Order("food_scan_result.created_at DESC").Find(&fsResults).Error; err != nil {
		return nil, 0, err
	}

	if len(fsResults) == 0 {
		return nil, 0, errors.New("foodscan results not found")
	}

	fsResultResponse := helpers.ToFoodScanResultsResponse(fsResults)
	return fsResultResponse, totalRows, nil
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

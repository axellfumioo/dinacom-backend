package services

import (
	"backend-dinakom/app/dto/request"
	"backend-dinakom/app/dto/response"
	"backend-dinakom/app/helpers"
	"backend-dinakom/app/models"
	"backend-dinakom/configs"
	"errors"
	"fmt"
	"path/filepath"

	"github.com/minio/minio-go/v7"
	"gorm.io/gorm"
)

type FoodScanService interface {
	ScanFood(userID string, req request.ScanFoodRequest) (*response.FoodScanResponse, error)
}

type foodScanService struct {
	db     *gorm.DB
	client *minio.Client
}

func NewFoodScanService(db *gorm.DB, client *minio.Client) FoodScanService {
	return &foodScanService{db: db, client: client}
}

func (s *foodScanService) ScanFood(userID string, req request.ScanFoodRequest) (*response.FoodScanResponse, error) {
	ext := filepath.Ext(req.Image.Filename)
	object := fmt.Sprintf("foodscans/%s%s", userID, ext)

	baseUrl := configs.AppConfig.Minio.BaseUrl
	bucket := configs.AppConfig.Minio.Bucket

	url, err := helpers.UploadFile(s.client, &req.Image, baseUrl, bucket, object)
	if err != nil {
		return nil, err
	}

	foodScan := &models.FoodScan{
		ImageURL: url,
		UserID:   userID,
	}
	if err := s.db.Create(&foodScan).Error; err != nil {
		return nil, errors.New("create food scan error")
	}

	foodScanResponse := helpers.ToFoodScanResponse(foodScan)
	return &foodScanResponse, nil
}

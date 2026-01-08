package services

import (
	"backend-dinakom/app/dto/payload"
	"backend-dinakom/app/dto/request"
	"backend-dinakom/app/dto/response"
	"backend-dinakom/app/helpers"
	"backend-dinakom/app/models"

	"backend-dinakom/configs"
	"encoding/json"
	"errors"
	"fmt"
	"path/filepath"

	"github.com/google/uuid"
	"github.com/hibiken/asynq"
	"github.com/minio/minio-go/v7"
	"gorm.io/gorm"
)

type FoodScanService interface {
	GetAllFoodScans(page int, pageSize int) (*[]response.FoodScanResponse, int64, error)
	GetUserFoodScans(id string) (*[]response.FoodScanResponse, error)
	ScanFood(userID string, req request.ScanFoodRequest) (*response.FoodScanResponse, error)
}

type foodScanService struct {
	db          *gorm.DB
	minioClient *minio.Client
	queueClient *asynq.Client
}

func NewFoodScanService(db *gorm.DB, client *minio.Client, queueClient *asynq.Client) FoodScanService {
	return &foodScanService{db: db, minioClient: client, queueClient: queueClient}
}

func (s *foodScanService) GetAllFoodScans(page int, pageSize int) (*[]response.FoodScanResponse, int64, error) {
	var foodScans []models.FoodScan
	var totalRows int64

	if err := s.db.Model(&models.FoodScan{}).Count(&totalRows).Error; err != nil {
		return nil, 0, errors.New("foodscan history data not founds")
	}

	offset := (page - 1) * pageSize

	if err := s.db.Preload("User").Offset(offset).Limit(pageSize).Order("created_at DESC").Find(&foodScans).Error; err != nil {
		return nil, 0, errors.New("failed to get foodScan history")
	}

	foodScansResponse := helpers.ToFoodScansResponse(foodScans)
	return &foodScansResponse, totalRows, nil
}

func (s *foodScanService) GetUserFoodScans(id string) (*[]response.FoodScanResponse, error) {
	var existingUser models.User
	if err := s.db.First(&existingUser, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("User is not registered")
		}
		return nil, err
	}

	var foodScans []models.FoodScan
	if err := s.db.Preload("User").Find(&foodScans, "user_id = ?", id).Error; err != nil {
		return nil, err
	}

	foodScansSesponse := helpers.ToFoodScansResponse(foodScans)
	return &foodScansSesponse, nil
}

func (s *foodScanService) ScanFood(userID string, req request.ScanFoodRequest) (*response.FoodScanResponse, error) {
	ext := filepath.Ext(req.Image.Filename)
	object := fmt.Sprintf(
		"foodscans/%s-%s-%s",
		userID,
		uuid.NewString(),
		ext,
	)

	baseUrl := configs.AppConfig.Minio.BaseUrl
	bucket := configs.AppConfig.Minio.Bucket

	url, err := helpers.UploadFile(s.minioClient, &req.Image, baseUrl, bucket, object)
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

	payload := payload.FoodScanPayload{
		FoodScanID: foodScan.ID,
		UserID:     userID,
		ImageURL:   url,
	}

	b, _ := json.Marshal(payload)
	task := asynq.NewTask("foodscan:process", b)
	s.queueClient.Enqueue(task)

	foodScanResponse := helpers.ToFoodScanResponse(foodScan)
	return &foodScanResponse, nil
}

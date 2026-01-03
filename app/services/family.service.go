package services

import (
	"backend-dinakom/app/dto/request"
	"backend-dinakom/app/helpers"
	"backend-dinakom/app/models"
	"backend-dinakom/configs"
	"errors"
	"fmt"
	"path/filepath"

	"github.com/minio/minio-go/v7"
	"gorm.io/gorm"
)

type FamilyService interface {
	CreateNewFamily(UserID string, req request.CreateFamilyRequest) (*models.Family, error)
}

type familyService struct {
	db          *gorm.DB
	minioClient *minio.Client
}

func NewFamilyService(db *gorm.DB, minioClient *minio.Client) FamilyService {
	return &familyService{db: db, minioClient: minioClient}
}

func (s *familyService) CreateNewFamily(UserID string, req request.CreateFamilyRequest) (*models.Family, error) {
	var existingUser models.User
	if err := s.db.Preload("Profile").First(&existingUser, "id = ?", UserID).Error; err != nil {
		return nil, errors.New("user query error:" + err.Error())
	}

	countAge := helpers.CalculateAge(existingUser.Profile.DateOfBirth)
	if countAge < 17 {
		return nil, errors.New("your age is not enough to create a family")
	}

	// File Upload
	ext := filepath.Ext(req.FamilyAvatar.Filename)
	object := fmt.Sprintf("family/%s%s", req.Name, ext)
	baseUrl := configs.AppConfig.Minio.BaseUrl
	bucket := configs.AppConfig.Minio.Bucket

	// Minio Uploader
	avatarUrl, err := helpers.UploadFile(s.minioClient, req.FamilyAvatar, baseUrl, bucket, object)
	if err != nil {
		return nil, err
	}

	// Create family handle
	family := &models.Family{
		Name:      req.Name,
		Desc:      &req.Description,
		AvatarUrl: avatarUrl,
		Member:    []models.FamilyMember{{UserID: UserID, Role: "PARENT"}},
	}
	if err := s.db.Create(&family).Error; err != nil {
		return nil, errors.New("failed to create family:" + err.Error())
	}

	return family, nil
}

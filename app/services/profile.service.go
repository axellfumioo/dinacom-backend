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

type ProfileService interface {
	UpdateAvatar(req request.UploadAvatarRequest) (*response.ProfileResponse, error)
}

type profileService struct {
	db     *gorm.DB
	client *minio.Client
}

func NewProfileService(db *gorm.DB, client *minio.Client) ProfileService {
	return &profileService{db: db, client: client}
}

func (s *profileService) UpdateAvatar(req request.UploadAvatarRequest) (*response.ProfileResponse, error) {
	var existingProfile *models.UserProfile
	if err := s.db.First(&existingProfile, "user_id", req.UserID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("user not found")
		} else {
			return nil, err
		}
	}

	// File handler
	ext := filepath.Ext(req.Avatar.Filename)
	object := fmt.Sprintf("avatars/%s%s", req.UserID, ext)

	baseUrl := configs.AppConfig.Minio.BaseUrl
	bucket := configs.AppConfig.Minio.Bucket

	url, err := helpers.UploadFile(s.client, req.Avatar, baseUrl, bucket, object)
	if err != nil {
		return nil, err
	}

	updates := map[string]interface{}{}

	if url != "" {
		updates["avatar"] = url
	}

	if err := s.db.Model(&existingProfile).Updates(updates).Error; err != nil {
		return nil, err
	}

	profileResponse := helpers.ToProfileResponse(existingProfile)
	return &profileResponse, nil
}

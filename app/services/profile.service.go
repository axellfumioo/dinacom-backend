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

	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
	"gorm.io/gorm"
)

type ProfileService interface {
	GetUserProfile(userId string) (*response.ProfileResponse, error)
	UpdateProfile(userId string, req request.UpdateProfileRequest) (*response.ProfileResponse, error)
	UploadAvatar(req request.UploadAvatarRequest) (*response.ProfileResponse, error)
}

type profileService struct {
	db     *gorm.DB
	client *minio.Client
}

func NewProfileService(db *gorm.DB, client *minio.Client) ProfileService {
	return &profileService{db: db, client: client}
}

func (s *profileService) GetUserProfile(userId string) (*response.ProfileResponse, error) {
	var existingUser models.User
	if err := s.db.First(&existingUser, "id = ?", userId).Error; err != nil {
		return  nil, errors.New("failed to get user: " + err.Error())
	}
	
	var existingProfile models.UserProfile
	if err := s.db.Preload("User").First(&existingProfile, "user_id = ?", userId).Error; err != nil {
		return  nil, errors.New("failed to get profile: " + err.Error())
	}

	profileResponse := helpers.ToProfileResponse(&existingProfile)
	return &profileResponse, nil
}

func (s *profileService) UpdateProfile(userId string, req request.UpdateProfileRequest) (*response.ProfileResponse, error) {
	var existingProfile *models.UserProfile
	if err := s.db.First(&existingProfile, "user_id", userId).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("profile not found")
		} else {
			return nil, err
		}
	}

	updates := map[string]interface{}{}

	if req.DateOfBirth != nil {
		updates["date_of_birth"] = &req.DateOfBirth
	}

	if req.WeightKG != nil {
		updates["weight_kg"] = &req.WeightKG
	}

	if req.HeightCM != nil {
		updates["height_cm"] = &req.HeightCM
	}

	if req.ActivityLevel != nil {
		updates["activity_level"] = &req.ActivityLevel
	}

	if err := s.db.Model(&existingProfile).Updates(updates).Error; err != nil {
		return nil, err
	}

	profileResponse := helpers.ToProfileResponse(existingProfile)

	return &profileResponse, nil
}

func (s *profileService) UploadAvatar(req request.UploadAvatarRequest) (*response.ProfileResponse, error) {
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
	object := fmt.Sprintf("avatars/%s%s%s", req.UserID, uuid.NewString(), ext)

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

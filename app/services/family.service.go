package services

import (
	"backend-dinakom/app/dto/request"
	"backend-dinakom/app/helpers"
	"backend-dinakom/app/models"
	"backend-dinakom/configs"
	"errors"
	"fmt"
	"mime/multipart"
	"path/filepath"

	"github.com/minio/minio-go/v7"
	"gorm.io/gorm"
)

type FamilyService interface {
	CreateNewFamily(UserID string, req request.CreateFamilyRequest) (*models.Family, error)
	UpdateFamilyAvatar(familyID string, userID string, avatar *multipart.FileHeader) (*models.Family, error)
	UpdateFamily(familyID string, userID string, req request.UpdateFamilyRequest) (*models.Family, error)
	DeleteFamily(familyID string, userID string) (*models.Family, error)
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
	if err := s.db.Preload("Profile").Preload("member_of").First(&existingUser, "id = ?", UserID).Error; err != nil {
		return nil, errors.New("failed to get user:" + err.Error())
	}

	if existingUser.MembersOf != nil {
		return nil, errors.New("failed to create family: this member already has a family")
	}

	countAge := helpers.CalculateAge(existingUser.Profile.DateOfBirth)
	if countAge < 17 {
		return nil, errors.New("failed to create family: user age is not enough to create a family")
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

func (s *familyService) UpdateFamilyAvatar(familyID string, userID string, avatar *multipart.FileHeader) (*models.Family, error) {
	// Check is exist
	var existingMember models.FamilyMember
	if err := s.db.First(&existingMember, "user_id = ? AND family_id = ?", userID, familyID).Error; err != nil {
		return nil, errors.New("failed to get member data:" + err.Error())
	}

	var existingFamily models.Family
	if err := s.db.First(&existingFamily, "id = ?", familyID).Error; err != nil {
		return nil, errors.New("failed to get family:" + err.Error())
	}

	// File Upload
	ext := filepath.Ext(avatar.Filename)
	object := fmt.Sprintf("family/%s%s", existingFamily.Name, ext)
	baseUrl := configs.AppConfig.Minio.BaseUrl
	bucket := configs.AppConfig.Minio.Bucket

	// Minio Uploader
	avatarUrl, err := helpers.UploadFile(s.minioClient, avatar, baseUrl, bucket, object)
	if err != nil {
		return nil, err
	}

	// Update avatar
	if avatarUrl != existingFamily.AvatarUrl {
		if err := s.db.Model(&existingFamily).Update("AvatarUrl", avatarUrl).Error; err != nil {
			return nil, errors.New("failed to update family avatar:" + err.Error())
		}
	}

	return &existingFamily, nil
}

func (s *familyService) UpdateFamily(familyID string, userID string, req request.UpdateFamilyRequest) (*models.Family, error) {
	// Check is exist
	var existingMember models.FamilyMember
	if err := s.db.First(&existingMember, "user_id = ? AND family_id = ?", userID, familyID).Error; err != nil {
		return nil, errors.New("failed to get member data:" + err.Error())
	}

	var existingFamily models.Family
	if err := s.db.First(&existingFamily, "id = ?", familyID).Error; err != nil {
		return nil, errors.New("failed to get family:" + err.Error())
	}

	updateData := map[string]interface{}{}
	if req.Name != nil && req.Name != &existingFamily.Name {
		updateData["name"] = req.Name
	}

	if req.Description != nil && req.Description != existingFamily.Desc {
		updateData["desc"] = req.Name
	}

	if err := s.db.Model(&existingFamily).Updates(updateData).Error; err != nil {
		return nil, errors.New("failed to update family:" + err.Error())
	}

	return &existingFamily, nil
}

func (s *familyService) DeleteFamily(familyID string, userID string) (*models.Family, error) {
	// Check if exist
	var existingMember models.FamilyMember
	if err := s.db.First(&existingMember, "user_id = ? AND family_id = ?", userID, familyID).Error; err != nil {
		return nil, errors.New("failed to get member data:" + err.Error())
	}

	var existingFamily models.Family
	if err := s.db.First(&existingFamily, "id = ?", familyID).Error; err != nil {
		return nil, errors.New("failed to get family:" + err.Error())
	}

	// Delete family
	if err := s.db.Delete(&existingFamily).Error; err != nil {
		return nil, errors.New("failed to delete family:" + err.Error())
	}

	return &existingFamily, nil
}

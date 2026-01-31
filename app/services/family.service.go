package services

import (
	"backend-dinakom/app/dto/request"
	"backend-dinakom/app/dto/response"
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
	GetUserFamily(UserID string) (*response.FamilyResponse, error)
	GetFamilyByID(ID string, UserID string) (*response.FamilyResponse, error)
	CreateNewFamily(UserID string, req request.CreateFamilyRequest) (*response.FamilyResponse, error)
	UpdateFamilyAvatar(familyID string, userID string, avatar *multipart.FileHeader) (*response.FamilyResponse, error)
	UpdateFamily(familyID string, userID string, req request.UpdateFamilyRequest) (*response.FamilyResponse, error)
	DeleteFamily(familyID string, userID string) (*response.FamilyResponse, error)
}

type familyService struct {
	db          *gorm.DB
	minioClient *minio.Client
}

func NewFamilyService(db *gorm.DB, minioClient *minio.Client) FamilyService {
	return &familyService{db: db, minioClient: minioClient}
}

func (s *familyService) GetUserFamily(UserID string) (*response.FamilyResponse, error) {
	var family models.Family
	if err := s.db.
		Joins("JOIN family_members fm ON fm.family_id = families.id").
		Where("fm.user_id = ?", UserID).
		Preload("Member.User").
		First(&family).Error; err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not in any family")
		}
		return nil, err
	}

	familyResponse := helpers.ToFamilyResponse(&family)
	return &familyResponse, nil
}

func (s *familyService) GetFamilyByID(ID string, UserID string) (*response.FamilyResponse, error) {
	var existingMember models.FamilyMember
	if err := s.db.First(&existingMember, "family_id = ? AND user_id = ?", ID, UserID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("access denied:you are not member in this family")
		}
		return nil, errors.New("failed to find member:" + err.Error())
	}

	var family models.Family
	if err := s.db.First(&family, "id = ?", ID).Error; err != nil {
		return nil, errors.New("failed to find family:" + err.Error())
	}

	familyResponse := helpers.ToFamilyResponse(&family)
	return &familyResponse, nil
}

func (s *familyService) CreateNewFamily(UserID string, req request.CreateFamilyRequest) (*response.FamilyResponse, error) {
	var existingUser models.User
	if err := s.db.Preload("Profile").Preload("MemberOf").First(&existingUser, "id = ?", UserID).Error; err != nil {
		return nil, errors.New("failed to get user:" + err.Error())
	}

	if existingUser.MemberOf != nil {
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

	familyResponse := helpers.ToFamilyResponse(family)
	return &familyResponse, nil
}

func (s *familyService) UpdateFamilyAvatar(familyID string, userID string, avatar *multipart.FileHeader) (*response.FamilyResponse, error) {
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

	familyResponse := helpers.ToFamilyResponse(&existingFamily)
	return &familyResponse, nil
}

func (s *familyService) UpdateFamily(familyID string, userID string, req request.UpdateFamilyRequest) (*response.FamilyResponse, error) {
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

	familyResponse := helpers.ToFamilyResponse(&existingFamily)
	return &familyResponse, nil
}

func (s *familyService) DeleteFamily(familyID string, userID string) (*response.FamilyResponse, error) {
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

	familyResponse := helpers.ToFamilyResponse(&existingFamily)
	return &familyResponse, nil
}

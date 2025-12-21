package services

import (
	"backend-dinakom/app/dto/request"
	"backend-dinakom/app/helpers"
	"backend-dinakom/configs"
	"fmt"
	"path/filepath"

	"github.com/minio/minio-go/v7"
	"gorm.io/gorm"
)

type ProfileService interface {
	UpdateAvatar(req request.UploadAvatarRequest) (string, error)
}

type profileService struct {
	db     *gorm.DB
	client *minio.Client
}

func NewProfileService(db *gorm.DB, client *minio.Client) ProfileService {
	return &profileService{db: db, client: client}
}

func (s *profileService) UpdateAvatar(req request.UploadAvatarRequest) (string, error) {
	ext := filepath.Ext(req.Avatar.Filename)
	object := fmt.Sprintf("avatars/%s%s", req.UserID, ext)

	baseUrl := configs.AppConfig.Minio.BaseUrl
	bucket := configs.AppConfig.Minio.Bucket

	url, err := helpers.UploadFile(s.client, req.Avatar, baseUrl, bucket, object)
	if err != nil {
		return "", nil
	}

	return url, nil
}

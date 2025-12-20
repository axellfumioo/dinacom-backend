package services

import (
	"backend-dinakom/app/dto/request"
	"backend-dinakom/configs"
	"context"
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

	src, err := req.Avatar.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	if _, err := s.client.PutObject(
		context.Background(),
		bucket,
		object,
		src,
		req.Avatar.Size,
		minio.PutObjectOptions{
			ContentType: req.Avatar.Header.Get("Content-Type"),
		},
	); err != nil {
		return "", err
	}

	url := fmt.Sprintf("%s/%s/%s", baseUrl, bucket, object)

	return url, nil
}

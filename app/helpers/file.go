package helpers

import (
	"context"
	"fmt"
	"mime/multipart"

	"github.com/minio/minio-go/v7"
)

func UploadFile(client *minio.Client, file *multipart.FileHeader, baseUrl, bucket string, object string) (string, error) {

	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	if _, err := client.PutObject(
		context.Background(),
		bucket,
		object,
		src,
		file.Size,
		minio.PutObjectOptions{
			ContentType: file.Header.Get("Content-Type"),
		},
	); err != nil {
		return "", err
	}

	url := fmt.Sprintf("%s/%s/%s", baseUrl, bucket, object)

	return url, nil
}

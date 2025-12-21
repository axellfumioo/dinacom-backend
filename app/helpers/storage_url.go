package helpers

import (
	"backend-dinakom/configs"
	"fmt"
)

func GenerateFileURL(bucket, object string) string {
	baseURL := configs.AppConfig.Minio.BaseUrl

	return fmt.Sprintf(
		"%s/%s/%s",
		baseURL,
		bucket,
		object,
	)
}
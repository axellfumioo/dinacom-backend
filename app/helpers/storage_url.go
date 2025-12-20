package helpers

import (
	"backend-dinakom/configs"
	"fmt"
)

func GenerateFileURL(bucket, object string) string {
	baseURL := configs.GetEnv("STORAGE_PUBLIC_URL", "")

	return fmt.Sprintf(
		"%s/%s/%s",
		baseURL,
		bucket,
		object,
	)
}
package configs

import (
	"log"
	
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func NewMinioClient() *minio.Client {
	client, err := minio.New(
		AppConfig.Minio.Endpoint,
		&minio.Options{
			Creds: credentials.NewStaticV4(
				AppConfig.Minio.AccessKey,
				AppConfig.Minio.SecretKey,
				"",
			),
			Secure: AppConfig.Minio.UseSSL,
		},
	)
	if err != nil {
		log.Fatal("failed init minio:", err)
	}

	return client
}

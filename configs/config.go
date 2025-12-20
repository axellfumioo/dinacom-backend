package configs

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	App      AppConfigType
	Database DatabaseConfig
	JWT      JWTConfig
	Minio    MinioConfig
}

type AppConfigType struct {
	Name string
	Env  string
	Port string
}

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	SSLMode  string
	Timezone string
}

type JWTConfig struct {
	Secret      string
	ExpireHours int
}

type MinioConfig struct {
	BaseUrl   string
	Bucket    string
	UseSSL    bool
	AccessKey string
	SecretKey string
	Endpoint  string
}

var AppConfig *Config

func LoadConfig() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found, using environment variables")
	}

	expireHours, _ := strconv.Atoi(GetEnv("JWT_EXPIRE_HOURS", "24"))
	minioSSL, _ := strconv.ParseBool(GetEnv("MINIO_USE_SSL", "false"))

	AppConfig = &Config{
		App: AppConfigType{
			Name: GetEnv("APP_NAME", "Backend-Dinacom"),
			Env:  GetEnv("APP_ENV", "development"),
			Port: GetEnv("APP_PORT", "8080"),
		},
		Database: DatabaseConfig{
			Host:     GetEnv("DB_HOST", "localhost"),
			Port:     GetEnv("DB_PORT", "5432"),
			User:     GetEnv("DB_USER", "postgres"),
			Password: GetEnv("DB_PASSWORD", "postgres"),
			Name:     GetEnv("DB_NAME", "test_crud_api"),
			SSLMode:  GetEnv("DB_SSL_MODE", "disable"),
			Timezone: GetEnv("DB_TIMEZONE", "Asia/Jakarta"),
		},
		JWT: JWTConfig{
			Secret:      GetEnv("JWT_SECRET", "suifhs8fgsigsigfseuih"),
			ExpireHours: expireHours,
		},
		Minio: MinioConfig{
			BaseUrl:   GetEnv("MINIO_BASE_URL", "http://localhost:9000"),
			Endpoint:  GetEnv("MINIO_ENDPOINT", "localhost:9000"),
			UseSSL:    minioSSL,
			AccessKey: GetEnv("MINIO_ACCESS_KEY", "minioadmin"),
			SecretKey: GetEnv("MINIO_SECRET_KEY", "minioadmin"),
			Bucket:    GetEnv("MINIO_BUCKET", ""),
		},
	}

	log.Println("Configuration loaded successfully")
}

func GetEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

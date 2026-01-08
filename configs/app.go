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
	Redis    RedisConfig
	Strava   StravaConfig
}

type AppConfigType struct {
	Name              string
	Env               string
	Port              string
	Frontend_URL      string
	AI_BACKEND_URL    string
	AI_BACKEND_BEARER string
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

type RedisConfig struct {
	ADDRESS    string
	CONCURENCY int
}

type StravaConfig struct {
	CLIENT_KEY string
	CLIENT_ID  string
}

var AppConfig *Config

func LoadConfig() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found, using environment variables")
	}

	expireHours, _ := strconv.Atoi(GetEnv("JWT_EXPIRE_HOURS", "24"))
	minioSSL, _ := strconv.ParseBool(GetEnv("MINIO_USE_SSL", "false"))
	concurency, _ := strconv.Atoi(GetEnv("REDIS_CONCURENCY", "5"))

	AppConfig = &Config{
		App: AppConfigType{
			Name:              GetEnv("APP_NAME", "Backend-Dinacom"),
			Env:               GetEnv("APP_ENV", "development"),
			Port:              GetEnv("APP_PORT", "8080"),
			Frontend_URL:      GetEnv("FRONTEND_BASE_URL", "http://localhost:3000"),
			AI_BACKEND_URL:    GetEnv("AI_BACKEND_BASE_URL", "http://localhost:8000/api/v1"),
			AI_BACKEND_BEARER: GetEnv("AI_BACKEND_BEARER", ""),
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
		Redis: RedisConfig{
			ADDRESS:    GetEnv("REDIS_ADDRESS", "localhost:6379"),
			CONCURENCY: concurency,
		},
		Strava: StravaConfig{
			CLIENT_KEY: GetEnv("STRAVA_CLIENT_KEY", ""),
			CLIENT_ID:  GetEnv("STRAVA_CLIENT_ID", ""),
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

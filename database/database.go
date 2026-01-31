package database

import (
	"backend-dinakom/configs"
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectDatabase() {
	cfg := configs.AppConfig.Database

	loc, locErr := time.LoadLocation(cfg.Timezone)
	if locErr != nil {
		log.Printf("Invalid/unknown time zone %q (%v); falling back to UTC", cfg.Timezone, locErr)
		cfg.Timezone = "UTC"
		loc = time.UTC
	}

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s TimeZone=%s",
		cfg.Host,
		cfg.Port,
		cfg.User,
		cfg.Password,
		cfg.Name,
		cfg.SSLMode,
		cfg.Timezone,
	)

	// Configure GORM logger
	var gormLogger logger.Interface
	if configs.AppConfig.App.Env == "development" {
		gormLogger = logger.Default.LogMode(logger.Info)
	} else {
		gormLogger = logger.Default.LogMode(logger.Error)
	}

	var db *gorm.DB
	var err error
	const maxAttempts = 30
	for attempt := 1; attempt <= maxAttempts; attempt++ {
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
			Logger: gormLogger,
			NowFunc: func() time.Time {
				return time.Now().In(loc)
			},
		})
		if err == nil {
			break
		}

		log.Printf("Failed to connect to database (attempt %d/%d): %v", attempt, maxAttempts, err)
		time.Sleep(2 * time.Second)
	}
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Set connection pool settings
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Failed to get database instance:", err)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	DB = db
	log.Println("Database connected successfully")
}

func GetDb() *gorm.DB {
	return DB
}

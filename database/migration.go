package database

import (
	"backend-dinakom/app/models"
	"log"
)

func RunMigration() {
	err := DB.AutoMigrate(
		&models.Role{},
		&models.User{},
		&models.UserProfile{},
		&models.FoodScan{},
		&models.FoodScanResult{},
		&models.UserMeal{},
		&models.AiChat{},
		&models.AIChatMessage{},
		&models.AIDecision{},
		&models.AIGoogleSearch{},
		&models.AIWebExtract{},
	)

	if err != nil {
		log.Fatal("Failed to run migration:", err)
	}

	log.Println("Database migration completed successfully")
}

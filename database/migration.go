package database

import (
	"backend-dinakom/app/models"
	"log"
)

func RunMigration() {
	err := DB.AutoMigrate(
		&models.User{},
	)

	if err != nil {
		log.Fatal("Failed to run migration:", err)
	}

	log.Println("Database migration completed successfully")
}

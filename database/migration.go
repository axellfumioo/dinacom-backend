package database

import "log"

func RunMigration() {
	err := DB.AutoMigrate()

	if err != nil {
		log.Fatal("Failed to run migration:", err)
	}

	log.Println("Database migration completed successfully")
}
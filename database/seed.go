package database

import (
	"backend-dinakom/seeds"

	"gorm.io/gorm"
)

func RunSeeder(db *gorm.DB) {
	seeds.SeedRole(db)
}
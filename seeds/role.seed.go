package seeds

import (
	"backend-dinakom/app/models"
	"log"

	"gorm.io/gorm"
)

func SeedRole(db *gorm.DB) {
	roles := []models.Role{
		{
			RoleName: "ADMIN",
		},
		{
			RoleName: "USER",
		},
	}

	for _, role := range roles {
		if err := db.FirstOrCreate(&role, models.Role{RoleName: role.RoleName}).Error; err != nil {
			log.Printf("Failed to seed investor type %s: %v", role.RoleName, err)
		}
	}

	log.Printf("roles seeded successfully")
}

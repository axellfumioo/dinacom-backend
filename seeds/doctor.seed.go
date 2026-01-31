package seeds

import (
	"backend-dinakom/app/models"
	"log"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func SeedDoctor(db *gorm.DB) {
	var doctorRole models.Role
	if err := db.Where("role_name = ?", "DOCTOR").First(&doctorRole).Error; err != nil {
		log.Fatalf("role DOCTOR not found: %v", err)
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword(
		[]byte("doctor123"),
		bcrypt.DefaultCost,
	)
	password := string(hashedPassword)

	doctors := []models.User{
		{
			Email:    "doctor1@dinakom.com",
			Password: &password,
			FullName: "Dr. Andi Wijaya",
			PhoneNumber: func() *string {
				p := "628123456789"
				return &p
			}(),
			RoleID: &doctorRole.ID,
		},
		{
			Email:    "doctor2@dinakom.com",
			Password: &password,
			FullName: "Dr. Ahmad Maharani",
			PhoneNumber: func() *string {
				p := "628987654321"
				return &p
			}(),
			RoleID: &doctorRole.ID,
		},
	}

	if err := db.Clauses(clause.OnConflict{DoNothing: true}).Create(&doctors).Error; err != nil {
		log.Fatalf("failed seed doctors: %v", err)
	}

	for _, d := range doctors {
		var user models.User
		if err := db.Where("email = ?", d.Email).First(&user).Error; err != nil {
			continue
		}

		profile := models.UserProfile{
			UserID: user.ID,
			Gender: "MALE",
		}

		db.FirstOrCreate(&profile, models.UserProfile{
			UserID: user.ID,
		})
	}
}

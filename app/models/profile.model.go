package models

import "time"

type UserProfile struct {
	UserID        string    `gorm:"type:uuid;primaryKey" json:"user_id"`
	DateOfBirth   time.Time `json:"date_of_birth" db:"date_of_birth"`
	Gender        string    `json:"gender" db:"gender"`
	HeightCM      float64   `json:"height_cm" db:"height_cm"`
	WeightKG      float64   `json:"weight_kg" db:"weight_kg"`
	ActivityLevel string    `json:"activity_level" db:"activity_level"`

	User *User `gorm:"foreignKey:UserID;references:UserID;constraint:OnDelete:CASCADE;"`

	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

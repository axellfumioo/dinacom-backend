package models

import "time"

type UserProfile struct {
	ID            string    `gorm:"type:uuid;default:gen_random_uuid();primaryKey;"`
	UserID        string    `gorm:"uniqueIndex;foreignKey:UserID;references:UserID;not null;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Avatar        *string   `gorm:"type:varchar(500);" json:"avatar"`
	DateOfBirth   time.Time `json:"date_of_birth"`
	Gender        string    `json:"gender"`
	HeightCM      *float64  `json:"height_cm"`
	WeightKG      *float64  `json:"weight_kg"`
	ActivityLevel *string   `json:"activity_level"`

	User *User

	CreatedAt time.Time `json:"created_at"`
}

func (UserProfile) TableName() string {
	return "user_profiles"
}

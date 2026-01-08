package models

import (
	"time"
)

type UserMeal struct {
	ID       string `gorm:"type:uuid;default:gen_random_uuid();primaryKey;" json:"id"`
	FoodName string `gorm:"type:text"`
	Calories float64
	Protein  float64
	Fat      float64
	Carbs    float64

	UserID string `gorm:"type:uuid;not null;" json:"user_id"`
	User   *User  `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user"`

	CreatedAt time.Time `json:"created_at"`
}

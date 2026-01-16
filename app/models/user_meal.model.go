package models

import (
	"backend-dinakom/app/constants"
	"time"
)

type UserMeal struct {
	ID       string `gorm:"type:uuid;default:gen_random_uuid();primaryKey;" json:"id"`
	FoodName string `gorm:"type:text"`
	Calories float64
	Protein  float64
	Fat      float64
	Carbs    float64

	Portion int                    `json:"portion"`
	Time    constants.UserMealTime `gorm:"type:varchar(25);not null"` // BREAKFAST, LUNCH, DINNER, SNACK

	UserID string `gorm:"type:uuid;not null;" json:"user_id"`
	User   *User  `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user"`

	CreatedAt time.Time `json:"created_at"`
}

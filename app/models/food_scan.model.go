package models

import (
	"backend-dinakom/app/constants"
	"time"
)

type FoodScan struct {
	ID       string `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	ImageURL string `gorm:"type:varchar(500);not null;" json:"image_url"`

	// Status PENDING, FAILED, SUCCESS
	Status constants.FoodScanStatus `gorm:"type:varchar(20);not null;default:'PENDING'"`

	// Owner
	UserID string `gorm:"type:uuid;not null;"`
	User   *User  `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (FoodScan) TableName() string {
	return "food_scans"
}

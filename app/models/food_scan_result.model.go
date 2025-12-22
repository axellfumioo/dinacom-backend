package models

import "time"

type FoodScanResult struct {
	ID string `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`

	FoodNames   []string `gorm:"type:json"`
	Calories    float64
	Protein     float64
	Fat         float64
	Carbs       float64
	Ingredients []string `gorm:"type:json"` // optional

	FoodScanID string `gorm:"type:uuid;not null;uniqueIndex;" json:"food_scan_id"`
	FoodScan *FoodScan `gorm:"foreignKey:FoodScanID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"food_scan"`

	CreatedAt time.Time
}

func (FoodScanResult) TableName() string {
	return "food_scan_result"
}

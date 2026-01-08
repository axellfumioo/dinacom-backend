package response

import (
	"time"
)

type FoodScanResultResponse struct {
	ID string `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`

	FoodNames    string `json:"food_names"`
	Foodtype     string   `json:"food_type"`
	CaloriesKcal float64  `json:"calories_kcal"`
	ProteinG     float64  `json:"protein_g"`
	FatG         float64  `json:"fat_g"`
	CarbsG       float64  `json:"carbs_g"`
	Vitamins     []string `json:"vitamins"`

	FoodScanID string            `gorm:"type:uuid;not null;uniqueIndex;" json:"food_scan_id"`
	FoodScan   *FoodScanResponse `gorm:"foreignKey:FoodScanID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"food_scan"`

	Confidence float64 `json:"confidence"`
	CreatedAt  time.Time
}

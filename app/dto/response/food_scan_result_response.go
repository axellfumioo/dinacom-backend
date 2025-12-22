package response

import (
	"time"
)

type FoodScanResultResponse struct {
	ID string `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`

	FoodNames   []string `json:"food_names"`
	Calories    float64  `json:"calories"`
	Protein     float64  `json:"protein"`
	Fat         float64  `json:"fat"`
	Carbs       float64  `json:"carbohydrate"`
	Ingredients []string `json:"ingrendients"`

	FoodScanID string            `gorm:"type:uuid;not null;uniqueIndex;" json:"food_scan_id"`
	FoodScan   *FoodScanResponse `gorm:"foreignKey:FoodScanID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"food_scan"`

	CreatedAt time.Time
}

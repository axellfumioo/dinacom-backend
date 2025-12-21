package response

import (
	"backend-dinakom/app/constants"
	"time"
)

type FoodScanResponse struct {
	ID       string `json:"id"`
	ImageURL string `json:"image_url"`

	// Status PENDING, FAILED, SUCCESS
	Status constants.FoodScanStatus `json:"status"`

	// Owner
	UserID string        `json:"user_id"`
	User   *UserResponse `json:"user"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

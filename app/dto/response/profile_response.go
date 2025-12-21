package response

import "time"

type ProfileResponse struct {
	ID            string
	UserID        string    `json:"user_id"`
	Avatar        *string   `json:"avatar"`
	DateOfBirth   time.Time `json:"date_of_birth"`
	Gender        string    `json:"gender"`
	HeightCM      *float64  `json:"height_cm"`
	WeightKG      *float64  `json:"weight_kg"`
	ActivityLevel *string   `json:"activity_level"`

	User *UserResponse `json:"user"`

	CreatedAt time.Time `json:"created_at"`
}

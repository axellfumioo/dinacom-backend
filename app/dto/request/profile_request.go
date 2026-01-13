package request

import (
	"mime/multipart"
	"time"
)

type UpdateProfileRequest struct {
	DateOfBirth   *time.Time `binding:"omitempty" json:"date_of_birth"`
	HeightCM      *float64   `binding:"omitempty" json:"height_cm"`
	WeightKG      *float64   `binding:"omitempty" json:"weight_kg"`
	ActivityLevel *string    `binding:"omitempty" json:"activity_level"`
}

type UploadAvatarRequest struct {
	UserID string                `json:"user_id"`
	Avatar *multipart.FileHeader `json:"avatar"`
}

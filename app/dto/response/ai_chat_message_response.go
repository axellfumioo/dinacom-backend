package response

import (
	"backend-dinakom/app/constants"
	"time"
)

type AIChatMessageResponse struct {
	ID         string                  `json:"id"`
	ImageURL   *string                 `json:"image_url"`
	Content    string                  `json:"content"`
	Confidence *float64                `json:"confidence"`
	SenderRole constants.AIMessageRole `json:"sender_role"`

	// Relations
	UserID *string `json:"user_id"`
	User *UserResponse `json:"user"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

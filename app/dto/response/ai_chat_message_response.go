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
	UserID *string         `json:"user_id"`
	ChatID string          `json:"chat_id"`
	User   *UserResponse   `json:"user"`
	Chat   *AiChatResponse `json:"chat"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

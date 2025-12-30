package response

import "time"

type AiChatResponse struct {
	ID     string
	UserID string `json:"user_id"`

	// Relation
	User     *UserResponse           `json:"user"`
	Messages []AIChatMessageResponse `json:"messages"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

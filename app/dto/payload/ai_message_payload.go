package payload

import "backend-dinakom/app/models"

type CreateAIMessagePayload struct {
	ChatID      string
	UserID      string
	ChatHistory []models.AIChatMessage `json:"chat_history"`
	Message     string                 `json:"member"`
}

package models

import (
	"backend-dinakom/app/constants"
	"time"
)

type AIChatMessage struct {
	ID string `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`

	ImageURL   *string                 `gorm:"type:varchar(500);"`
	Content    string                  `gorm:"type:text;"`
	Confidence *float64                `gorm:"type:float"`
	SenderRole constants.AIMessageRole `gorm:"type:varchar(20);default:'USER'" json:"sender_role"` // USER, ASSISTANT

	// Relations
	ChatID string  `gorm:"type:uuid;not null" json:"chat_id"`
	UserID *string `gorm:"type:uuid" json:"user_id"`

	Chat *AiChat `gorm:"references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	User *User   `gorm:"references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (AIChatMessage) TableName() string {
	return "ai_chat_messages"
}

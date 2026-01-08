package models

import "time"

type AIChatMessageSource struct {
	ID    string `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Url   string `gorm:"type:text" json:"url"`
	Title string `gorm:"type:text" json:"title"`
	Query string `gorm:"type:text" json:"query"`

	MessageID string        `gorm:"type:uuid;not null"`
	Message   AIChatMessage `gorm:"foreignKey:MessageID;references:ID;constraint:OnUpdate:CASCADE;OnDelete:CASCADE"`

	CreatedAt time.Time `json:"created_at"`
}

func (AIChatMessageSource) TableName() string {
	return "ai_chat_message_sources"
}

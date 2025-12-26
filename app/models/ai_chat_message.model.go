package models

import "time"

type AIChatMessage struct {
	ID string `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`

	ImageURL   *string `gorm:"type:varchar(500);"`
	Content    string  `gorm:"type:varchar(500);"`
	Confidence int

	// Relations
	ChatID string  `gorm:"type:uuid;not null" json:"chat_id"`
	UserID *string `gorm:"type:uuid" json:"user_id"`

	Chat *AiChat `gorm:"foreignKey:ChatID;references:ID;constraint:OnUpdate:CASCADE;OnDelete:CASCADE"`
	User *User   `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE;OnDelete:CASCADE"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (AIChatMessage) TableName() string {
	return "ai_chat_messages"
}

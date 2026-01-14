package models

import "time"

type AiChat struct {
	ID        string `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	UserID    string `gorm:"type:uuid;not null;uniqueIndex" json:"user_id"`
	CreatedAt time.Time
	UpdatedAt time.Time

	User     *User           `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE;OnDelete:CASCADE"`
	Messages []AIChatMessage `gorm:"foreignKey:ChatID" json:"messages"`
}

func (AiChat) TableName() string {
	return "ai_chats"
}

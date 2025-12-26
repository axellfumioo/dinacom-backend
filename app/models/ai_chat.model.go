package models

import "time"

type AiChat struct {
	ID        string `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	UserID    string `gorm:"type:uuid;not null;uniqueIndex" json:"user_id"`
	UserMood  string ``
	Summary   string `gorm:"type:text"`
	CreatedAt time.Time
	UpdatedAt time.Time

	User *User `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE;OnDelete:CASCADE"`
}

func (AiChat) TableName() string {
	return "ai_chats"
}

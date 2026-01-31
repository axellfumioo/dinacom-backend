package models

import "time"

type DoctorChatMessage struct {
	ID           string `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	DoctorChatID string `gorm:"type:uuid;"`
	SenderID     string `gorm:"type:uuid;"`
	Message      string `gorm:"type:text;"`

	CreatedAt time.Time
}

func (DoctorChatMessage) TableName() string {
	return "doctor_chat_messages"
}
package models

import "time"

type DoctorChatMessage struct {
	ID           string `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	DoctorChatID string `gorm:"type:uuid;" json:"doctor_chat_id"`
	SenderID     string `gorm:"type:uuid;" json:"sender_id"`
	Message      string `gorm:"type:text;" json:"message"`

	DoctorChat DoctorChatRoom `gorm:"foreignKey:DoctorChatID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt  time.Time
}

func (DoctorChatMessage) TableName() string {
	return "doctor_chat_messages"
}

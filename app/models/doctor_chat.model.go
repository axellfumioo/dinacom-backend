package models

import "time"

type DoctorChatRoom struct {
	ID       string `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	DoctorID string `gorm:"type:uuid;not null;index:idx_doctor_user,unique"`
	UserID   string `gorm:"type:uuid;not null;index:idx_doctor_user,unique"`

	Doctor *User `gorm:"foreignKey:DoctorID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"doctor"`
	User   *User `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"user"`

	CreatedAt time.Time
}

func (DoctorChatRoom) TableName() string {
	return "doctor_chat_rooms"
}

package models

import "time"

type Questionnaire struct {
	ID       string  `gorm:"type:uuid;default:gen_random_uuid();primaryKey;"`
	Number   int     `gorm:""`
	Question string  `gorm:"type:text;not null"`
	Answer   *string `gorm:"type:text"`

	UserID string `gorm:"type:uuid;not null"`
	User   *User  `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE, OnDelete:CASCADE"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (Questionnaire) TableName() string {
	return "questionnaires"
}

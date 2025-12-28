package models

import "time"

type StravaToken struct {
	ID string `gorm:"type:uuid;default:gen_random_uuid();primaryKey;" json:"id"`

	RefreshToken string `gorm:"type:text;not null" json:"refresh_token"`
	AccessToken  string `gorm:"type:text;not null" json:"access_token"`
	ExpiresAt    int64  `json:"expires_at"`

	UserID string `gorm:"type:uuid;uniqueIndex"`
	User   *User  `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	CreatedAt time.Time `json:"created_at"`
}

func (StravaToken) TableName() string {
	return "strava_tokens"
}

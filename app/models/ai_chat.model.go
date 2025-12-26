package models

import "time"

type AiChat struct {
	ID         string `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	user_id    string `gorm:"type:uuid;not null"`
	user_mood  string
	summary    string `gorm:"type:text"`
	created_at time.Time 
	updated_at time.Time
}

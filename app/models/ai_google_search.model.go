package models

import "time"

type AIGoogleSearch struct {
	ID      string `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	URL     string `gorm:"type:varchar(500);" json:"url"`
	Content string `gorm:"type:text"`

	DecisionID string      `gorm:"type:uuid;"`
	Decision   *AIDecision `gorm:"foreignKey:DecisionID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"decision"`

	CreatedAt time.Time
}

func (AIGoogleSearch) TableName() string {
	return "ai_google_search"
}

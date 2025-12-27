package models

import "time"

type AIWebExtract struct {
	ID      string `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Domain  string `gorm:"type:text"`
	Content string `gorm:"type:text"`

	DecisionID string      `gorm:"type:uuid;"`
	Decision   *AIDecision `gorm:"foreignKey:DecisionID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"decision"`

	CreatedAt time.Time
}

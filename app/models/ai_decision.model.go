package models

import (
	"backend-dinakom/app/constants"
	"time"
)

type AIDecision struct {
	ID string `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`

	NeedSearch  bool                  `gorm:"default:false" json:"need_search"`
	RequestType string                `gorm:"type:varchar(100);not null" json:"request_type"`
	RiskLevel   constants.AIRiskLevel `gorm:"type:varchar(20);"`

	// ForeignKey
	ChatMessageID string `gorm:"type:uuid;uniqueIndex"`
	// Relations

	AIChatMessage *AIChatMessage `gorm:"foreignKey:ChatMessageID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	CreatedAt time.Time
}


func (AIDecision) TableName() string {
	return "ai_decisions"
}

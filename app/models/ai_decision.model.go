package models

import (
	"backend-dinakom/app/constants"
	"time"

	"github.com/lib/pq"
)

type AIDecision struct {
	ID string `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`

	Queries     pq.StringArray        `gorm:"type:text[]"`
	NeedSearch  bool                  `gorm:"default:false" json:"need_search"`
	RequestType string                `gorm:"type:varchar(100);not null" json:"request_type"`
	RiskLevel   constants.AIRiskLevel `gorm:"type:varchar(20);"`

	// ForeignKey
	ChatMessageID string           `gorm:"type:uuid;uniqueIndex"`
	AIChatMessage *AIChatMessage   `gorm:"foreignKey:ChatMessageID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Searchs       []AIGoogleSearch `gorm:"-" json:"google_searchs"`
	WebExtracts   []AIWebExtract   `gorm:"-" json:"web_extracts"`

	CreatedAt time.Time
}

func (AIDecision) TableName() string {
	return "ai_decisions"
}

package response

import (
	"backend-dinakom/app/constants"
	"time"
)

type AIDecisionResponse struct {
	ID string `json:"id"`

	Queries     []string              `json:"queries"`
	NeedSearch  bool                  `json:"need_search"`
	RequestType string                `json:"request_type"`
	RiskLevel   constants.AIRiskLevel `json:"risk_level"`

	ChatMessageID string `json:"chat_message_id"`
	// AIChatMessage any   `gorm:"foreignKey:ChatMessageID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Searchs []AIGoogleSearchResponse `json:"google_searchs"`
	// WebExtracts   []any   `gorm:"-" json:"web_extracts"`

	CreatedAt time.Time
}

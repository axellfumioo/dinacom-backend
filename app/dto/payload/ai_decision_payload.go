package payload

import "backend-dinakom/app/constants"

type CreateAIDecisionPayload struct {
	Queries     []string              `json:"queries"`
	NeedSearch  bool                  `json:"need_search"`
	RequestType string                `json:"request_type"`
	RiskLevel   constants.AIRiskLevel `json:"risk_level"` // LOW, MEDIUM, HIGH
}

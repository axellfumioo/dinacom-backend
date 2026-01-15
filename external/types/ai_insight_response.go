package types

type AIInsightResponse struct {
	HealthScore int `json:"health_score"`
	PersonalAIInsight string `json:"personal_ai_insight"`
	AIImportantNotice string `json:"ai_important_notice"`
	ConfidenceLevel int `json:"confidence_level"` 
}

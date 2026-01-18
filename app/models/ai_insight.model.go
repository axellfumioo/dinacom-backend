package models

type AIInsight struct {
	ID string `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`

	HealthScore       int    `json:"health_score"`
	PersonalAIInsight string `gorm:"type:text" json:"personal_ai_insight"`
	AIImportantNotice string `gorm:"type:text" json:"ai_important_notice"`
	Confidence        int    `json:"confidence"`
}

func (AIInsight) TableName() string {
	return "ai_insights"
}

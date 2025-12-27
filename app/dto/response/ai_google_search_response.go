package response

import "time"

type AIGoogleSearchResponse struct {
	ID      string
	URL     string `json:"url"`
	Content string `json:"content"`

	DecisionID string             `json:"decision_id"`
	Decision   AIDecisionResponse `json:"decision"`

	CreatedAt time.Time
}

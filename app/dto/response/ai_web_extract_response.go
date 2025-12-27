package response

import "time"

type AIWebExtractResponse struct {
	ID      string
	Domain  string `json:"domain"`
	Content string `json:"content"`

	DecisionID string              `json:"decision_id"`
	Decision   *AIDecisionResponse `json:"decision"`

	CreatedAt time.Time
}

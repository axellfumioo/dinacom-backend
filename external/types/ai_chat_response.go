package types

type AIChatResponse struct {
	Message    string  `json:"message"`
	Confidence float64 `json:"confidence"`
}

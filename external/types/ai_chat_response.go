package types

type AIChatAPIResponse struct {
	Response string `json:"response"`
}

type AIChatResponse struct {
	Message    string  `json:"message"`
	Confidence float64 `json:"confidence"`
}

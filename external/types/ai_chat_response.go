package types

type AIChatAPIResponse struct {
	Answer     string             `json:"answer"`
	Sources    []AISourceResponse `json:"sources"`
	Confidence float64            `json:"confidence"`
}

type AIChatResponse struct {
	Message    string
	Confidence float64
	Sources    []AISourceResponse
}

type AISourceResponse struct {
	Url   string `json:"url"`
	Title string `json:"title"`
	Query string `json:"query"`
}

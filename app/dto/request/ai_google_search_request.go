package request

type CreateGoogleSearchRequest struct {
	URL        string `binding:"required"`
	Content    string `binding:"required"`
	DecisionID string `binding:"required,uuid"`
}

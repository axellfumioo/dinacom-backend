package request

type CreateGoogleSearchRequest struct {
	URL        string `binding:"required"`
	Content    string `binding:"required"`
	DecisionID string `binding:"required,uuid"`
}

type UpdateGoogleSearchRequest struct {
	URL     *string `binding:"optional"`
	Content *string `binding:"optional"`
}

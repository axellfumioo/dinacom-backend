package request

type CreateWebExtractRequest struct {
	Domain     string `binding:"required"`
	Content    string `binding:"required"`
	DecisionID string `binding:"required,uuid"`
}

type UpdateWebExtractRequest struct {
	Domain  *string `binding:"optional"`
	Content *string `binding:"optional"`
}

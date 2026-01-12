package request

type UpdateQuestionnairesRequest struct {
	Answers []QuestionnaresAnswerRequest
}

type QuestionnaresAnswerRequest struct {
	QuestionID     string `json:"question_id"`
	Number int    `json:"number"`
	Answer string `json:"answer"`
}

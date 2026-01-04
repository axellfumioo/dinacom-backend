package request

type AddFamilyMemberRequest struct {
	Members []Member `binding:"required"`
}

type Member struct {
	UserID string
	Role   string
}

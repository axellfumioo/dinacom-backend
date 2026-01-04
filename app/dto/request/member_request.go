package request

type AddFamilyMemberRequest struct {
	Members  []Member `binding:"required"`
	FamilyID string   `binding:"required,min:3"`
}

type Member struct {
	UserID string
	Role   string
}

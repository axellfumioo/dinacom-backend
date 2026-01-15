package response

import (
	"backend-dinakom/app/constants"
	"time"
)

type FamilyMemberResponse struct {
	ID       string               `json:"id"`
	Role     constants.MemberRole `json:"role"`
	FamilyID string               `json:"family_id"`
	UserID   string               `json:"user_id"`

	User   *UserResponse   `json:"user"`
	Family *FamilyResponse `json:"family"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

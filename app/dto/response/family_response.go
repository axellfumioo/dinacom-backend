package response

import "time"

type FamilyResponse struct {
	ID        string  `json:"id"`
	Name      string  `json:"name"`
	AvatarUrl string  `json:"avatar_url"`
	Desc      *string `json:"description"`

	Member []FamilyMemberResponse `json:"members"`

	CreatedAt time.Time `json:"created_at"`
	UpdateAt  time.Time `json:"updated_at"`
}

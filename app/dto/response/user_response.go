package response

import "time"

type UserResponse struct {
	UserID      string  `json:"user_id"`
	Email       string  `json:"email"`
	FullName    string  `json:"full_name"`
	PhoneNumber *string `json:"phone_number"`

	RoleID string        `json:"role_id"`
	Role   *RoleResponse `json:"role"`

	Profile   *ProfileResponse `json:"profile"`
	CreatedAt time.Time        `json:"created_at"`
	UpdatedAt time.Time        `json:"updated_at"`
}

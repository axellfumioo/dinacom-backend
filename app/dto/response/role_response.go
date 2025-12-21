package response

import "time"

type RoleResponse struct {
	RoleID   string `json:"role_id"`
	RoleName string `json:"role_name"`

	Users     *[]UserResponse `json:"user"`
	UpdatedAt time.Time       `json:"updated_at"`
	CreatedAt time.Time       `json:"created_at"`
}

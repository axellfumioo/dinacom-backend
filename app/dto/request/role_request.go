package request

type CreateRoleRequest struct {
	Name string `json:"name" binding:"required,min=3,max=100"`
}

type UpdateRoleRequest struct {
	Name string `json:"name" binding:"required,min=3,max=100"`
}
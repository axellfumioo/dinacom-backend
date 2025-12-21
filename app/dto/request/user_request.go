package request

type CreateUserRequest struct {
	Name        string `json:"name" binding:"required,min=3,max=100"`
	Email       string `json:"email" binding:"required,email"`
	Password    string `json:"password" binding:"required,min=6"`
	PhoneNumber string `json:"phone_number" binding:"omitempty,min=10,max=20"`
}

type UpdateUserRequest struct {
	Name        string `json:"name" binding:"omitempty,min=3,max=100"`
	Email       string `json:"email" binding:"omitempty,email"`
	PhoneNumber string `json:"phone_number" binding:"omitempty,min=10,max=20"`
}

type ChangePasswordRequest struct {
	OldPassword string `json:"old_password" binding:"required"`
	NewPassword string `json:"new_password" binding:"required,min=6"`
}

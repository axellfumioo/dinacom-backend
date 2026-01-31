package request

import "time"

type RegisterRequest struct {
	Name        string    `json:"name" validate:"required,min=3"`
	Email       string    `json:"email" validate:"required"`
	Password    string    `json:"password" validate:"required"`
	DateOfBirth time.Time `json:"date_of_birth" validate:"required"`
	Gender      string    `json:"gender" validate:"required,min=1"`
	PhoneNumber string    `json:"phone_number"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

package helpers

import (
	"backend-dinakom/app/dto/response"
	"backend-dinakom/app/models"
)

func ToUserResponse(user *models.User) response.UserResponse {
	return response.UserResponse{
		UserID:      user.UserID,
		FullName:    user.FullName,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
		Role:        user.Role,
		CreatedAt:   user.CreatedAt,
		UpdatedAt:   user.UpdatedAt,
	}
}

func ToUsersResponse(users []models.User) []response.UserResponse {
	var usersResponse []response.UserResponse
	for _, user := range users {
		usersResponse = append(usersResponse, ToUserResponse(&user))
	}
	return usersResponse
}

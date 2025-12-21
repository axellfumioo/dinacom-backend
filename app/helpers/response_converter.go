package helpers

import (
	"backend-dinakom/app/dto/response"
	"backend-dinakom/app/models"
)

func ToUserResponse(user *models.User) response.UserResponse {
	var profileResponse *response.ProfileResponse
	if user.Profile != nil {
		r := ToProfileResponse(user.Profile)
		profileResponse = &r

	}
	return response.UserResponse{
		UserID:      user.UserID,
		FullName:    user.FullName,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
		Role:        nil,
		Profile:     profileResponse,
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

func ToProfileResponse(profile *models.UserProfile) response.ProfileResponse {
	var userResponse *response.UserResponse
	if profile.User != nil {
		r := ToUserResponse(profile.User)
		userResponse = &r
	}

	return response.ProfileResponse{
		ID:            profile.ID,
		UserID:        profile.UserID,
		Avatar:        profile.Avatar,
		Gender:        profile.Gender,
		DateOfBirth:   profile.DateOfBirth,
		HeightCM:      profile.HeightCM,
		WeightKG:      profile.WeightKG,
		ActivityLevel: profile.ActivityLevel,
		User:          userResponse,
		CreatedAt:     profile.CreatedAt,
	}
}

func ToProfilesResponse(profiles []models.UserProfile) []response.ProfileResponse {
	var profilesResponse []response.ProfileResponse
	for _, profile := range profiles {
		profilesResponse = append(profilesResponse, ToProfileResponse(&profile))
	}
	return profilesResponse
}

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
		UserID:      user.ID,
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

func ToRoleResponse(role *models.Role) response.RoleResponse {
	var usersResponse []response.UserResponse
	if role.Users != nil {
		r := ToUsersResponse(role.Users)
		usersResponse = r
	}
	return response.RoleResponse{
		RoleID:    role.ID,
		RoleName:  role.RoleName,
		Users:     &usersResponse,
		UpdatedAt: role.UpdatedAt,
		CreatedAt: role.CreatedAt,
	}
}

func ToRolesResponse(roles []models.Role) []response.RoleResponse {
	var rolesResponse []response.RoleResponse
	for _, role := range roles {
		rolesResponse = append(rolesResponse, ToRoleResponse(&role))
	}
	return rolesResponse
}

func ToFoodScanResponse(fs *models.FoodScan) response.FoodScanResponse {
	var userResponse response.UserResponse
	if fs.User != nil {
		r := ToUserResponse(fs.User)
		userResponse = r
	}
	return response.FoodScanResponse{
		ID:        fs.ID,
		ImageURL:  fs.ImageURL,
		Status:    fs.Status,
		UserID:    fs.UserID,
		User:      &userResponse,
		CreatedAt: fs.CreatedAt,
		UpdatedAt: fs.UpdatedAt,
	}
}

func ToFoodScansResponse(foodScans []models.FoodScan) []response.FoodScanResponse {
	var foodScansResponse []response.FoodScanResponse
	for _, fs := range foodScans {
		foodScansResponse = append(foodScansResponse, ToFoodScanResponse(&fs))
	}
	return foodScansResponse
}

func ToFoodScanResultResponse(fsResult *models.FoodScanResult) response.FoodScanResultResponse {
	var foodScanResponse response.FoodScanResponse
	if fsResult.FoodScan != nil {
		r := ToFoodScanResponse(fsResult.FoodScan)
		foodScanResponse = r
	}
	return response.FoodScanResultResponse{
		ID:         fsResult.ID,
		FoodNames:  fsResult.FoodNames,
		Calories:   fsResult.Calories,
		Protein:    fsResult.Protein,
		Fat:        fsResult.Fat,
		Carbs:      fsResult.Carbs,
		FoodScanID: fsResult.FoodScanID,
		FoodScan:   &foodScanResponse,
		CreatedAt:  fsResult.CreatedAt,
	}
}

func ToFoodScanResultsResponse(foodScans []models.FoodScanResult) []response.FoodScanResultResponse {
	var foodScanResultResponse []response.FoodScanResultResponse
	for _, fsresult := range foodScans {
		foodScanResultResponse = append(foodScanResultResponse, ToFoodScanResultResponse(&fsresult))
	}
	return foodScanResultResponse
}

func ToUserMealResponse(um *models.UserMeal) response.UserMealResponse {
	var userResponse response.UserResponse
	if um.User != nil {
		r := ToUserResponse(um.User)
		userResponse = r
	}
	return response.UserMealResponse{
		ID:        um.ID,
		FoodNames: um.FoodNames,
		Calories:  um.Calories,
		Protein:   um.Protein,
		Fat:       um.Fat,
		Carbs:     um.Carbs,
		User:      &userResponse,
		UserID:    um.UserID,
		CreatedAt: um.CreatedAt,
	}
}

func ToUserMealsResponse(userMeals []models.UserMeal) []response.UserMealResponse {
	var userMealsResponse []response.UserMealResponse
	for _, um := range userMeals {
		userMealsResponse = append(userMealsResponse, ToUserMealResponse(&um))
	}
	return userMealsResponse
}

func ToAIDecisionResponse(decision *models.AIDecision) response.AIDecisionResponse {
	return response.AIDecisionResponse{
		ID:          decision.ID,
		Queries:     decision.Queries,
		NeedSearch:  decision.NeedSearch,
		RequestType: decision.RequestType,
		RiskLevel:   decision.RiskLevel,
		CreatedAt:   decision.CreatedAt,
	}
}

func ToAIDecisionsResponse(decisions []models.AIDecision) []response.AIDecisionResponse {
	var decisionsResponse []response.AIDecisionResponse
	for _, decision := range decisions {
		decisionsResponse = append(decisionsResponse, ToAIDecisionResponse(&decision))
	}
	return decisionsResponse
}

func ToAIGoogleSearchResponse(googleSearch *models.AIGoogleSearch) response.AIGoogleSearchResponse {
	var decision response.AIDecisionResponse
	if googleSearch.Decision != nil {
		r := ToAIDecisionResponse(googleSearch.Decision)
		decision = r
	}
	return response.AIGoogleSearchResponse{
		ID:         googleSearch.ID,
		URL:        googleSearch.URL,
		Content:    googleSearch.Content,
		DecisionID: googleSearch.DecisionID,
		Decision:   &decision,
		CreatedAt:  googleSearch.CreatedAt,
	}
}

func ToAIGoogleSearchsResponse(googleSearchs []models.AIGoogleSearch) []response.AIGoogleSearchResponse {
	var googleSearchsResponse []response.AIGoogleSearchResponse
	for _, googleSearch := range googleSearchs {
		googleSearchsResponse = append(googleSearchsResponse, ToAIGoogleSearchResponse(&googleSearch))
	}
	return googleSearchsResponse
}

func ToAIWebExtractResponse(we *models.AIWebExtract) response.AIWebExtractResponse {
	var decision response.AIDecisionResponse
	if we.Decision != nil {
		r := ToAIDecisionResponse(we.Decision)
		decision = r
	}
	return response.AIWebExtractResponse{
		ID:         we.ID,
		Domain:     we.Domain,
		Content:    we.Content,
		DecisionID: we.DecisionID,
		Decision:   &decision,
		CreatedAt:  we.CreatedAt,
	}
}

func ToAIWebExtractsResponse(webExtracts []models.AIWebExtract) []response.AIWebExtractResponse {
	var webExtractsResponse []response.AIWebExtractResponse
	for _, we := range webExtracts {
		webExtractsResponse = append(webExtractsResponse, ToAIWebExtractResponse(&we))
	}
	return webExtractsResponse
}

func ToAIChatResponse(aiChat *models.AiChat) response.AiChatResponse {
	var user response.UserResponse
	if aiChat.User != nil {
		r := ToUserResponse(aiChat.User)
		user = r
	}
	var messages []response.AIChatMessageResponse
	if aiChat.Messages != nil {
		r := ToAIChatMessagesResponse(aiChat.Messages)
		messages = r
	}

	return response.AiChatResponse{
		ID:        aiChat.ID,
		Messages:  messages,
		UserID:    aiChat.UserID,
		User:      &user,
		CreatedAt: aiChat.CreatedAt,
		UpdatedAt: aiChat.UpdatedAt,
	}
}

func ToAIChatsResponse(aiChats []models.AiChat) []response.AiChatResponse {
	var aiChatsResponse []response.AiChatResponse
	for _, ac := range aiChats {
		aiChatsResponse = append(aiChatsResponse, ToAIChatResponse(&ac))
	}
	return aiChatsResponse
}

func ToAIChatMessageResponse(msg *models.AIChatMessage) response.AIChatMessageResponse {
	var user response.UserResponse
	if msg.User != nil {
		r := ToUserResponse(msg.User)
		user = r
	}
	return response.AIChatMessageResponse{
		ID:         msg.ID,
		ImageURL:   msg.ImageURL,
		Content:    msg.Content,
		Confidence: msg.Confidence,
		SenderRole: msg.SenderRole,
		UserID:     msg.UserID,
		User:       &user,
		CreatedAt:  msg.CreatedAt,
		UpdatedAt:  msg.UpdatedAt,
	}
}

func ToAIChatMessagesResponse(msgs []models.AIChatMessage) []response.AIChatMessageResponse {
	var messageResponse []response.AIChatMessageResponse
	for _, msg := range msgs {
		messageResponse = append(messageResponse, ToAIChatMessageResponse(&msg))
	}
	return messageResponse
}

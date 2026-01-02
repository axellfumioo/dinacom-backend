package extservices

import (
	"backend-dinakom/app/models"
	"backend-dinakom/configs"
	"backend-dinakom/external/types"
	"encoding/json"
	"errors"
	"fmt"
)

// AI Fetching (ntar dipisah)
func FetchFoodScanAI(imageURL string) (string, error) {
	restyClient := configs.RestyClient
	apiBaseUrl := configs.AppConfig.App.AI_BACKEND_URL
	bodyRequest := map[string]interface{}{"image_url": imageURL}

	var result *types.AIFoodScanResponse
	url := fmt.Sprintf("%s/ai/foodscan", apiBaseUrl)
	_, err := restyClient.R().
		SetBody(&bodyRequest).
		SetAuthToken("DinacomAIService#2025").
		SetResult(&result).
		Post(url)
	if err != nil || result == nil {
		return "", errors.New("failed to fetch ai foodscan")
	}

	jsonBytes, err := json.Marshal(result.Response)
	if err != nil {
		return "", err
	}

	return string(jsonBytes), nil
}

// AI chat answer
func FetchAIChat(message string, chatHistory []models.AIChatMessage) (string, error) {
	var restyClient = configs.RestyClient
	var chatHs []map[string]interface{}
	for _, hs := range chatHistory {
		chatHs = append(chatHs, map[string]interface{}{"role": hs.SenderRole, "content": hs.Content})
	}

	bodyRequest := map[string]interface{}{
		"content":      message,
		"chat_history": chatHs,
	}

	var result map[string]interface{}
	url := fmt.Sprintf("%s/", "")
	_, err := restyClient.R().
		SetBody(&bodyRequest).
		SetAuthToken("DinacomAIService#2025").
		SetResult(&result).
		Post(url)

	// Return AI {response :""}
	response := types.AIChatResponse{
		Message:    "Halo juga fiky",
		Confidence: 9.0,
	}

	jsonBytes, err := json.Marshal(response)
	if err != nil {
		return "", err
	}
	return string(jsonBytes), nil
}

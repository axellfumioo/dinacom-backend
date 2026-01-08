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

	var result types.AIFoodScanResponse
	url := fmt.Sprintf("%s/ai/foodscan", apiBaseUrl)
	resp, err := restyClient.R().
		SetBody(map[string]string{
			"image_url": imageURL,
		}).
		SetAuthToken(configs.AppConfig.App.AI_BACKEND_BEARER).
		SetResult(&result).
		Post(url)

	if err != nil {
		return "", err
	}

	if resp.IsError() {
		return "", errors.New("ai foodscan returned error status")
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
	apiBaseUrl := configs.AppConfig.App.AI_BACKEND_URL

	for _, hs := range chatHistory {
		chatHs = append(chatHs, map[string]interface{}{"role": hs.SenderRole, "content": hs.Content})
	}

	bodyRequest := map[string]interface{}{
		"message":      message,
		"chat_history": chatHs,
	}

	var result *types.AIChatAPIResponse
	url := fmt.Sprintf("%s/ai/chat", apiBaseUrl)
	_, err := restyClient.R().
		SetBody(&bodyRequest).
		SetAuthToken(configs.AppConfig.App.AI_BACKEND_BEARER).
		SetResult(&result).
		Post(url)
	if err != nil || result == nil {
		return "", errors.New("failed to fetch ai chat")
	}

	response := types.AIChatResponse{
		Message:    result.Response,
		Confidence: 3,
	}

	jsonBytes, err := json.Marshal(response)
	if err != nil {
		return "", err
	}
	return string(jsonBytes), nil
}

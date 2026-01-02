package extservices

import (
	"backend-dinakom/app/models"
	"backend-dinakom/configs"
	"backend-dinakom/external/types"
	"encoding/json"
	"fmt"
	"time"
)

// AI Fetching (ntar dipisah)
func FetchFoodScanAI(imageURL string) (string, error) {
	// Simulasi fetch ke AI
	_ = imageURL
	time.Sleep(2 * time.Second)
	var result = types.AINutritionResponse{
		FoodName: []string{"Tempe", "Tahu", "Ayam"},
		FoodType: "traditional",
		Nutrition: types.Nutritions{
			ProteinG: 8,
			CarbsG:   6,
			FatG:     10,
		},
		CaloriesKcal:    730,
		Vitamins:        make([]string, 0),
		HealthScores:    8,
		HealthNote:      "sdsds",
		ConfidenceScore: 9.1,
	}

	jsonBytes, err := json.Marshal(result)
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

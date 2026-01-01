package extservices

import (
	"backend-dinakom/app/models"
	"backend-dinakom/external/types"
	"encoding/json"
	"time"
)

// AI Fetching (ntar dipisah)
func FetchFoodScanAI(imageURL string) (string, error) {
	// Simulasi fetch ke AI
	_ = imageURL
	time.Sleep(2 * time.Second)
	var result = types.AINutritionResponse{
		FoodName:      []string{"Tempe", "Tahu", "Ayam"},
		Calories:      10,
		Protein:       8.4,
		Carbohydrates: 10.1,
		Fat:           5.1,
	}

	jsonBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}

	return string(jsonBytes), nil
}

// AI chat answer
func FetchAIChat(message string, chatHistory []models.AIChatMessage) (string, error) {
	var chatHs []map[string]interface{}
	for _, hs := range chatHistory {
		chatHs = append(chatHs, map[string]interface{}{"role": hs.SenderRole, "content": hs.Content})
	}

	bodyRequest := map[string]interface{}{
		"content":      message,
		"chat_history": chatHs,
	}

	// Return AI {response :""}
	_ = bodyRequest

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

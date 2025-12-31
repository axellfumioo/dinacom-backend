package workers

import (
	"backend-dinakom/app/models"
	"backend-dinakom/app/types"
	"backend-dinakom/app/types/payload"
	"context"
	"encoding/json"
	"errors"
	"log"
	"time"

	"github.com/hibiken/asynq"
	"gorm.io/gorm"
)

// AI Fetching (ntar dipisah)
func fetchFoodScanAI(imageURL string) (string, error) {
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
func fetchAIChat(message string, chatHistory []models.AIChatMessage) (string, error) {
	var chatHs []map[string]interface{}
	for _, hs := range chatHistory {
		chatHs = append(chatHs, map[string]interface{}{"role": hs.SenderRole, "content": hs.Content})
	}

	bodyRequest := map[string]interface{}{
		"content":      message,
		"chat_history": chatHs,
	}
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

func FoodScanProcess(ctx context.Context, t asynq.Task, db *gorm.DB) error {
	var payload payload.FoodScanPayload
	if err := json.Unmarshal(t.Payload(), &payload); err != nil {
		return err
	}

	var fs models.FoodScan
	if err := db.Where("id = ? AND user_id = ?", &payload.FoodScanID, &payload.UserID).First(&fs).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.New("foodscan data not found")
		}
		return err
	}

	result, err := fetchFoodScanAI(payload.ImageURL)
	if err != nil {
		if err := db.Model(&fs).Update("Status", "FAILED").Error; err != nil {
			return errors.New("error when update foodScan status")
		}
		return err
	}

	var jsonResult types.AINutritionResponse
	if err := json.Unmarshal([]byte(result), &jsonResult); err != nil {
		return err
	}

	fsResult := models.FoodScanResult{
		FoodScanID: fs.ID,
		FoodNames:  jsonResult.FoodName,
		Calories:   float64(jsonResult.Calories),
		Protein:    jsonResult.Protein,
		Carbs:      jsonResult.Carbohydrates,
		Fat:        jsonResult.Fat,
	}

	if err := db.Create(&fsResult).Error; err != nil {
		if err := db.Model(&fs).Update("Status", "FAILED").Error; err != nil {
			return errors.New("error when updating foodScan status")

		}
		return err
	}

	userMeal := models.UserMeal{
		UserID:    payload.UserID,
		FoodNames: jsonResult.FoodName,
		Calories:  float64(jsonResult.Calories),
		Protein:   jsonResult.Protein,
		Carbs:     jsonResult.Carbohydrates,
		Fat:       jsonResult.Fat,
	}

	if err := db.Create(&userMeal).Error; err != nil {
		return errors.New("error when creating user_meal")
	}

	if err := db.Model(&fs).Update("Status", "SUCCESS").Error; err != nil {
		return errors.New("error when updating foodScan status")
	}

	log.Printf("Emit event WS for user %s, result: %s\n", fs.UserID, result)
	return nil
}

func AIChatProcess(ctx context.Context, t asynq.Task, db *gorm.DB) error {
	var payload payload.CreateAIMessagePayload
	if err := json.Unmarshal(t.Payload(), &payload); err != nil {
		return err
	}

	response, err := fetchAIChat(payload.Message, payload.ChatHistory)
	if err != nil {
		return err
	}

	var jsonResult types.AIChatResponse
	if err := json.Unmarshal([]byte(response), &jsonResult); err != nil {
		return err
	}

	// Create Chat
	message := &models.AIChatMessage{
		ChatID:     payload.ChatID,
		Content:    jsonResult.Message,
		Confidence: &jsonResult.Confidence,
		SenderRole: "AI",
	}
	if err := db.Create(&message).Error; err != nil {
		return err
	}

	log.Printf("Emit event WS for chatRoom %s", payload.ChatID)
	return nil
}

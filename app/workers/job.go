package workers

import (
	"backend-dinakom/app/dto/payload"
	"backend-dinakom/app/models"
	"backend-dinakom/app/socket"
	extservices "backend-dinakom/external/services"
	"backend-dinakom/external/types"

	"context"
	"encoding/json"
	"errors"
	"log"

	"github.com/hibiken/asynq"
	"gorm.io/gorm"
)

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

	result, err := extservices.FetchFoodScanAI(payload.ImageURL)
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

	socket.EmitToUser(fs.UserID, "refresh:foodscan", &fs)
	log.Printf("foodscan:process Done")
	return nil
}

func AIChatProcess(ctx context.Context, t asynq.Task, db *gorm.DB) error {
	var payload payload.CreateAIMessagePayload
	if err := json.Unmarshal(t.Payload(), &payload); err != nil {
		return err
	}

	response, err := extservices.FetchAIChat(payload.Message, payload.ChatHistory)
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

	socket.EmitToRoom(payload.ChatID, "refresh:room", &message.ChatID)
	log.Printf("aichat:proccess Done")
	return nil
}

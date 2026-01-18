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
	"gorm.io/gorm/clause"
)

func FoodScanProcess(ctx context.Context, t asynq.Task, db *gorm.DB) error {
	var payload payload.FoodScanPayload
	if err := json.Unmarshal(t.Payload(), &payload); err != nil {
		return err
	}

	var fs models.FoodScan
	if err := db.Where("id = ? AND user_id = ?", payload.FoodScanID, payload.UserID).First(&fs).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Printf("foodscan not found")
			return errors.New("foodscan data not found")
		}
		return err
	}

	result, err := extservices.FetchFoodScanAI(fs.ImageURL)
	if err != nil {
		log.Printf(err.Error())
		if err := db.Model(&fs).Update("Status", "FAILED").Error; err != nil {
			return errors.New("error when update foodScan status")
		}
		socket.EmitToUser(fs.UserID, "refresh:foodscan", &fs)
		return err
	}

	var jsonResult types.AINutritionResponse
	if err := json.Unmarshal([]byte(result), &jsonResult); err != nil {
		return err
	}

	fsResult := models.FoodScanResult{
		FoodScanID:   fs.ID,
		FoodName:     jsonResult.FoodName,
		FoodType:     jsonResult.FoodType,
		CaloriesKcal: float64(jsonResult.CaloriesKcal),
		ProteinG:     jsonResult.Nutrition.ProteinG,
		CarbsG:       jsonResult.Nutrition.CarbsG,
		FatG:         jsonResult.Nutrition.FatG,
		Confidence:   jsonResult.ConfidenceScore,
		Vitamins:     jsonResult.Vitamins,
	}

	if err := db.Create(&fsResult).Error; err != nil {
		if err := db.Model(&fs).Update("Status", "FAILED").Error; err != nil {
			return errors.New("error when updating foodScan status")

		}
		socket.EmitToUser(fs.UserID, "refresh:foodscan", &fs)
		return err
	}

	if err := db.Model(&fs).Update("Status", "SUCCESS").Error; err != nil {
		socket.EmitToUser(fs.UserID, "refresh:foodscan", &fs)
		return errors.New("error when updating foodScan status")
	}

	socket.EmitToUser(fs.UserID, "refresh:foodscan", &fs)
	log.Printf("foodscan:process Done")
	return nil
}

func AIChatProcess(ctx context.Context, t asynq.Task, db *gorm.DB) error {
	var payload payload.CreateAIMessagePayload
	if err := json.Unmarshal(t.Payload(), &payload); err != nil {
		return nil
	}

	errorMessage := &models.AIChatMessage{
		ChatID:     payload.ChatID,
		Content:    "Aku mengalami kendala saat memproses permintaan ini. Silakan kirim pesan lagi.",
		Confidence: nil,
		SenderRole: "ASSISTANT",
	}

	response, err := extservices.FetchAIChat(payload.Message, payload.ChatHistory)
	if err != nil {
		if err := db.Create(&errorMessage).Error; err != nil {
			return err
		}
		socket.EmitToRoom(payload.ChatID, "refresh:room", &errorMessage.ChatID)
		return nil
	}

	var jsonResult types.AIChatResponse
	if err := json.Unmarshal([]byte(response), &jsonResult); err != nil {
		if err := db.Create(&errorMessage).Error; err != nil {
			return err
		}
		socket.EmitToRoom(payload.ChatID, "refresh:room", &errorMessage.ChatID)
		return nil
	}

	// Create Chat
	message := &models.AIChatMessage{
		ChatID:     payload.ChatID,
		Content:    jsonResult.Message,
		Confidence: &jsonResult.Confidence,
		SenderRole: "ASSISTANT",
	}
	if err := db.Create(&message).Error; err != nil {
		return err
	}

	var sources []models.AIChatMessageSource
	for _, source := range jsonResult.Sources {
		sources = append(sources, models.AIChatMessageSource{Url: source.Url, Title: source.Title, Query: source.Query, MessageID: message.ID})
	}

	if err := db.Clauses(clause.OnConflict{DoNothing: true}).Create(&sources).Error; err != nil {
		return err
	}

	socket.EmitToRoom(payload.ChatID, "refresh:room", &message.ChatID)
	log.Printf("aichat:proccess Done")
	return nil
}

func AIDecisionProcess(ctx context.Context, t asynq.Task, db *gorm.DB) error {
	var payload payload.CreateAIDecisionPayload
	if err := json.Unmarshal(t.Payload(), &payload); err != nil {
		return err
	}

	decision := &models.AIDecision{
		Queries:     payload.Queries,
		NeedSearch:  payload.NeedSearch,
		RequestType: payload.RequestType,
		RiskLevel:   payload.RiskLevel,
	}

	if err := db.Create(&decision).Error; err != nil {
		return errors.New("failed to create decision:" + err.Error())
	}

	if payload.NeedSearch == true {
		log.Println("create need search")
	}

	log.Println("ai-decision:process Done")
	return nil
}

func AIInsightProcess(ctx context.Context, t asynq.Task, db *gorm.DB) error {
	var payload payload.AIInsightPayload
	if err := json.Unmarshal(t.Payload(), &payload); err != nil {
		return nil
	}

	response, err := extservices.FetchAIInsight(payload)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	var jsonResult types.AIInsightResponse
	if err := json.Unmarshal([]byte(response), &jsonResult); err != nil {
		log.Println(err.Error())
		return nil
	}

	Insight := models.AIInsight{
		UserID:            payload.User.ID,
		HealthScore:       jsonResult.HealthScore,
		PersonalAIInsight: jsonResult.PersonalAIInsight,
		AIImportantNotice: jsonResult.AIImportantNotice,
		Confidence:        jsonResult.ConfidenceLevel,
	}
	if err := db.Create(&Insight).Error; err != nil {
		log.Println(err.Error())
		return err
	}

	socket.EmitToUser(payload.User.ID, "refresh:dashboard", Insight);
	return nil
}

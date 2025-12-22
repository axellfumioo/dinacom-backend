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

func fetchAI(imageURL string) (string, error) {
	// Simulasi fetch ke AI
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

	result, err := fetchAI(payload.ImageURL)
	if err != nil {
		if err := db.Model(&fs).Update("Status", "FAILED").Error; err != nil {
			return err
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
			return err
		}
		return err
	}

	if err := db.Model(&fs).Update("Status", "SUCCESS").Error; err != nil {
		return err
	}

	log.Printf("Emit event WS for user %d, result: %s\n", fs.UserID, result)
	return nil
}

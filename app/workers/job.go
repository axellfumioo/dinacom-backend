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

type JobHandler interface {
	FoodScanProcess() error
}

type jobHandler struct {
	ctx *context.Context
	t   *asynq.Task
	db  *gorm.DB
}

func NewJobHandler(db gorm.DB, ctx context.Context, t *asynq.Task) JobHandler {
	return &jobHandler{db: &db, ctx: &ctx, t: t}
}

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

func (s *jobHandler) FoodScanProcess() error {
	var payload payload.FoodScanPayload
	if err := json.Unmarshal(s.t.Payload(), &payload); err != nil {
		return err
	}

	var fs models.FoodScan
	if err := s.db.Where("id = ? AND user_id = ?", payload.FoodScanID, payload.UserID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.New("foodscan data not found")
		}
		return err
	}

	result, err := fetchAI(payload.ImageURL)
	if err != nil {
		s.db.Model(&fs).Update("Status", "Failed")
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

	if err := s.db.Create(&fsResult).Error; err != nil {
		return err
	}

	if err := s.db.Model(&fs).Updates(&models.FoodScan{Status: "SUCCESS"}).Error; err != nil {
		return err
	}

	log.Printf("Emit event WS for user %d, result: %s\n", fs.UserID, result)
	return nil
}

package workers

import (
	"backend-dinakom/configs"
	"context"
	"log"

	"github.com/hibiken/asynq"
	"gorm.io/gorm"
)

func StartWorker(db *gorm.DB) {
	server := configs.NewAsynqServer()
	mux := asynq.NewServeMux()

	// Foodscan handler
	mux.HandleFunc("foodscan:process", func(ctx context.Context, t *asynq.Task) error {
		return FoodScanProcess(ctx, *t, db)
	})

	// AI Chat Answer handler
	mux.HandleFunc("aichat:process", func(ctx context.Context, t *asynq.Task) error {
		return AIChatProcess(ctx, *t, db)
	})

	// AI Decision handler
	mux.HandleFunc("ai-decision:process", func(ctx context.Context, t *asynq.Task) error {
		return AIDecisionProcess(ctx, *t, db)
	})
	
	if err := server.Run(mux); err != nil {
		log.Fatal(err)
	}
}

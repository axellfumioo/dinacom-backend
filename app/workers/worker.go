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

	mux.HandleFunc("foodscan:process", func(ctx context.Context, t *asynq.Task) error {
		return FoodScanProcess(ctx, *t, db)
	})
	if err := server.Run(mux); err != nil {
		log.Fatal(err)
	}
}

package workers

import (
	"backend-dinakom/configs"
	"log"

	"github.com/hibiken/asynq"
	"gorm.io/gorm"
)

func StartWorker(db *gorm.DB) {
	server := configs.NewAsynqServer()
	mux := asynq.NewServeMux()

	if err := server.Run(mux); err != nil {
		log.Fatal(err)
	}
}

package configs

import "github.com/hibiken/asynq"

func NewAsynqServer() *asynq.Server {
	server := asynq.NewServer(
		asynq.RedisClientOpt{Addr: AppConfig.Redis.ADDRESS},
		asynq.Config{
			Concurrency: AppConfig.Redis.CONCURENCY,
		},
	)
	return server
}

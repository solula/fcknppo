package asynq_server

import (
	"github.com/redis/go-redis/v9"
)

type existingRedisClientOpt struct {
	client *redis.Client
}

func (r existingRedisClientOpt) MakeRedisClient() interface{} {
	return r.client
}

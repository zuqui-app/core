package services

import (
	"github.com/redis/go-redis/v9"

	"zuqui-core/internal"
)

var RedisClient *redis.Client

func init() {
	opt, _ := redis.ParseURL(internal.Env.UPSTASH_REDIS_URI)
	RedisClient = redis.NewClient(opt)
}

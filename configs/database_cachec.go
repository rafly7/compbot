package configs

import (
	"fmt"

	"github.com/go-redis/redis/v8"
)

func DatabaseCache() *redis.Client {
	redisHost := LoadEnv("REDIS_HOST")
	redisPort := LoadEnv("REDIS_PORT")
	redisPass := LoadEnv("REDIS_PASS")
	addr := fmt.Sprintf("%s:%s", redisHost, redisPort)
	return redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: redisPass,
		DB:       0,
	})
}

package jotnar

import (
	"github.com/go-redis/redis/v7"
)

type Redis struct {
	redisSet map[string]*redis.Client
}

var redisInstance *Redis

func initRedis() {
	redisInstance = &Redis{make(map[string]*redis.Client)}

	redisInstance.redisSet["main"] = redis.NewClient(&redis.Options{
		Addr:       RedisConfig["main"].Addr,
		Password:   RedisConfig["main"].Password,
		DB:         RedisConfig["main"].DB,
		MaxRetries: RedisConfig["main"].MaxRetries,
	})

	if RedisConfig["salve"] != nil {
		redisInstance.redisSet["salve"] = redis.NewClient(&redis.Options{
			Addr:       RedisConfig["salve"].Addr,
			Password:   RedisConfig["salve"].Password,
			DB:         RedisConfig["salve"].DB,
			MaxRetries: RedisConfig["salve"].MaxRetries,
		})
	}
}

func ReadRedis() *redis.Client {
	if sc, ok := redisInstance.redisSet["slave"]; ok {
		return sc
	}
	return redisInstance.redisSet["main"]
}

func WriteRedis() *redis.Client {
	return redisInstance.redisSet["main"]
}

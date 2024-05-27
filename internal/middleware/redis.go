package middleware

import (
	"fmt"
	"log"
	"os"

	"github.com/azusachino/ficus/global"
	"github.com/go-redis/redis"
)

var RedisClient *redis.Client

func InitRedis() {
	redisHost := os.Getenv(global.PG_HOST)
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", redisHost, os.Getenv(global.REDIS_PORT)),
		Password: os.Getenv(global.REDIS_PASS),
	})

	_, err := RedisClient.Ping().Result()
	if err != nil {
		log.Fatal("unbale to connect to Redis", err)
	}

	log.Println("connected to Redis server", redisHost)
}

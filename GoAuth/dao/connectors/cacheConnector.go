package connectors

import (
	"github.com/go-redis/redis"
	"fmt"
)

var client *redis.Client

//Connect to redis using provided credentials
func InitializeCache(redisURI string) {
	client = redis.NewClient(&redis.Options{
		Addr:     redisURI,
		PoolSize:  1000,
		MaxRetries: 5,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()
	if err == nil && len(pong) > 0 {
		fmt.Println("Cache Connection Successful")
	} else {
		fmt.Errorf("Cache Connection Failed")
	}
}

func GetCacheClient() *redis.Client{
	return client
}
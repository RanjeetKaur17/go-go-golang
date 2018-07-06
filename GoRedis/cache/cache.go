package cache

import (
	"github.com/go-redis/redis"
	"encoding/json"
	"time"
)

var redisClient *redis.Client

func InitializeRedis() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		PoolSize:  100,
		MaxRetries: 2,
		Password: "",
		DB:       0,
	})

	ping, err := redisClient.Ping().Result()
	if err == nil && len(ping) > 0 {
		println("Connected to Redis")
	} else {
		println("Redis Connection Failed")
	}
}

func GetValue(key string) (interface{}, error) {
	var deserializedValue interface{}
	serializedValue, err := redisClient.Get(key).Result()
	json.Unmarshal([]byte(serializedValue), &deserializedValue)
	return deserializedValue, err
}

func SetValue(key string, value interface{}) (bool, error) {
	serializedValue, _ := json.Marshal(value)
	err := redisClient.Set(key, string(serializedValue), 0).Err()
	return true, err
}

func SetValueWithTTL(key string, value interface{}, ttl int) (bool, error) {
	serializedValue, _ := json.Marshal(value)
	err := redisClient.Set(key, string(serializedValue), time.Duration(ttl)*time.Second).Err()
	return true, err
}

func RPush(key string, valueList []string) (bool, error) {
	err := redisClient.RPush(key, valueList).Err()
	return true, err;
}

func RpushWithTTL(key string, valueList []string, ttl int) (bool, error) {
	err := redisClient.RPush(key, valueList, ttl).Err()
	return true, err;
}

func LRange(key string) (bool, error) {
	err := redisClient.LRange(key, 0, -1).Err()
	return true, err;
}

func ListLength(key string) (int64)  {
	return redisClient.LLen(key).Val()
}

func Publish(channel string, message string)  {
	redisClient.Publish(channel, message)
}

func GetKeyListByPattern(pattern string) []string {
	return redisClient.Keys(pattern).Val()
}

func IncrementValue(key string) (int64)  {
	return redisClient.Incr(key).Val()
}

func DelKey(key string) (error) {
	return redisClient.Del(key).Err()
}
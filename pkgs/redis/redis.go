package redis

import (
	"context"
	"encoding/json"
	"fmt"

	redis "github.com/go-redis/redis/v8"
)

var (
	clientRedis *redis.Client
	ctx         = context.Background()
)

func init() {
	clientRedis = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
}

func Get(key string) (string, error) {
	val, err := clientRedis.Get(ctx, key).Result()
	if err != nil {
		return val, err
	}
	return val, nil
}

func GetAll(domain string) ([]interface{}, error) {
	pattern := fmt.Sprintf("*" + domain + "*")
	val, err := clientRedis.Keys(ctx, pattern).Result()
	if err != nil {
		return nil, err
	}
	teams, err := clientRedis.MGet(ctx, val...).Result()
	if err != nil {
		return nil, err
	}
	return teams, nil
}

func Set(key string, value interface{}) error {
	jsonValue, err := json.Marshal(value)
	if err != nil {
		return err
	}
	if err := clientRedis.Set(ctx, key, jsonValue, 0).Err(); err != nil {
		return err
	}
	return nil
}

func Delete(key string) (int64, error) {
	result, err := clientRedis.Del(ctx, key).Result()
	if err != nil {
		return result, err
	}
	if result == 0 {
		return result, err
	}
	return result, nil
}

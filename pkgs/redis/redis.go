package redis

import (
	"context"
	"encoding/json"

	redis "github.com/go-redis/redis/v8"
)

var (
	clientRedis *redis.Client
	ctx         = context.Background()
)

const myhash_teams = "teams"

func init() {
	clientRedis = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
}

func HashGet(team_id string) (string, error) {
	val, err := clientRedis.HGet(ctx, myhash_teams, team_id).Result()
	if err != nil {
		return val, err
	}
	return val, nil
}

func HashGetAll() (map[string]string, error) {
	result, err := clientRedis.HGetAll(ctx, myhash_teams).Result()
	if err != nil {
		return nil, err
	}
	return result, nil
}

func HashSet(team_id string, value interface{}) error {
	jsonValue, err := json.Marshal(value)
	if err != nil {
		return err
	}
	if err := clientRedis.HSet(ctx, myhash_teams, team_id, jsonValue).Err(); err != nil {
		return err
	}
	return nil
}

func HDel(team_id string) (int64, error) {
	result, err := clientRedis.HDel(ctx, myhash_teams, team_id).Result()
	if err != nil {
		return result, err
	}
	if result == 0 {
		return result, err
	}
	return result, nil
}

func Get(key string) (string, error) {
	val, err := clientRedis.Get(ctx, key).Result()
	if err != nil {
		return val, err
	}
	return val, nil
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

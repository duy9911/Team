package publisher

import (
	"context"
	"encoding/json"
	"errors"

	redis "github.com/go-redis/redis/v8"
)

var redisPublisher = redis.NewClient(&redis.Options{
	Addr: "localhost:6379",
})

func Publishing(newchannl string, value interface{}) error {
	var ctx = context.Background()

	payload, err := json.Marshal(value)
	if err != nil {
		return errors.New("error marshal before publishing")
	}

	if err := redisPublisher.Publish(ctx, "add_staff", payload).Err(); err != nil {
		return errors.New("error publishing")
	}
	return nil
}

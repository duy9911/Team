package subscribe

import (
	"context"
	"encoding/json"

	redis "github.com/go-redis/redis/v8"
)

var redisClient = redis.NewClient(&redis.Options{
	Addr: "localhost:6379",
})

func Subscribing(channel string) (interface{}, error) {

	var ctx = context.Background()
	subscriber := redisClient.Subscribe(ctx, "validate_pipeline")
	msg, err := subscriber.ReceiveMessage(ctx)
	if err != nil {
		panic(err)
	}
	var payload string
	if err := json.Unmarshal([]byte(msg.Payload), &payload); err != nil {
		panic(err)
	}
	return payload, nil

}

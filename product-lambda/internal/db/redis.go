package db

import (
	"context"

	"github.com/go-redis/redis/v8"
)

func InitRedis(addr string) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: addr,
	})
}

var Ctx = context.Background()

// context.Background() is a default, empty context:

// It has no timeout
// It can’t be cancelled
// It has no values

// A blank sheet of paper that you pass to other functions, which can add their own rules
// (like timeout, cancellation, etc.)

package rock

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type red struct {
	client *redis.Client
}

func (r *red) Get(ctx context.Context, key string) (err error) {
	r.client.Get(ctx, key)
}

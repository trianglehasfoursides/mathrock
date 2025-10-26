package rock

import "context"

var Rock store

type store interface {
	Get(ctx context.Context, key string) error
	Set(ctx context.Context, key string, val string) error
	Del(ctx context.Context, key string) error
}

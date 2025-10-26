package rock

import (
	"context"
	"errors"
	"sync"
)

type rizz struct {
	cache sync.Map
}

func (r *rizz) Get(_ context.Context, key string) (val string, err error) {
	value, exist := r.cache.Load(key)
	if !exist {
		return "", errors.New("")
	}

	val = value.(string)
	return val, nil
}

func (r *rizz) Set(_ context.Context, key string, val string) (err error) {
}

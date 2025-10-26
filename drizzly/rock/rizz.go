package rock

import (
	"context"
	"errors"

	"github.com/dgraph-io/ristretto/v2"
)

type rizz struct {
	cache *ristretto.Cache[string, string]
}

func (r *rizz) Get(_ context.Context, key string) (val string, err error) {
	val, exist := r.cache.Get(key)
	if !exist {
		return "", errors.New("")
	}

	return val, nil
}

func (r *rizz) Set(_ context.Context, key string, val string) (err error) {
}

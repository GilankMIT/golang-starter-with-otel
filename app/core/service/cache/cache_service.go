package cache

import "context"

type CacheService interface {
	SetNamespace(string) CacheAccessor
}

type CacheAccessor interface {
	GetString(ctx context.Context, key string) string
	SetString(ctx context.Context, key string, val string)
	Clean(ctx context.Context) error
}

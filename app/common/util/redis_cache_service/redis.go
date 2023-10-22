package redis_cache_service

import (
	"context"
	"encoding/json"
	"github.com/redis/go-redis/v9"
	"go-boilerplate/common/util/logutil"
	"go-boilerplate/core/service/cache"
)

type redisCache struct {
	redisClient *redis.Client
}

type redisCacheAccessor struct {
	namespace   string
	redisClient *redis.Client
}

func NewRedisCache(redisClient *redis.Client) cache.CacheService {
	return &redisCache{
		redisClient: redisClient,
	}
}

func (r redisCache) SetNamespace(s string) cache.CacheAccessor {
	return redisCacheAccessor{namespace: s, redisClient: r.redisClient}
}

func (r redisCacheAccessor) GetString(ctx context.Context, key string) string {
	var cacheValMap map[string]string

	data, err := r.redisClient.Get(ctx, r.namespace).Bytes()
	if err != nil {
		return ""
	}

	err = json.Unmarshal(data, &cacheValMap)
	if err != nil {
		logutil.LogError(ctx, "failed when accessing cache data", err.Error())
		return ""
	}

	return cacheValMap[key]
}

func (r redisCacheAccessor) SetString(ctx context.Context, key string, val string) {
	var cacheValMap = make(map[string]string)

	data, err := r.redisClient.Get(ctx, r.namespace).Bytes()
	if err != nil {
		if err == redis.Nil {
			cacheValMap[key] = val
		} else {
			logutil.LogError(ctx, "failed when accessing cache data", err.Error())
			return
		}
	} else {
		err = json.Unmarshal(data, &cacheValMap)
		if err != nil {
			logutil.LogError(ctx, "failed when marshalling cache data", err.Error())
			return
		}
		cacheValMap[key] = val
	}

	cacheValJson, err := json.Marshal(cacheValMap)
	if err != nil {
		logutil.LogError(ctx, "failed when marshalling cache data", err.Error())
		return
	}

	r.redisClient.Set(ctx, r.namespace, cacheValJson, 0)
}

func (r redisCacheAccessor) Clean(ctx context.Context) error {
	return r.redisClient.Del(ctx, r.namespace).Err()
}

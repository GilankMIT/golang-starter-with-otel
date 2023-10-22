package bigcache_client

import (
	"context"
	"encoding/json"
	"github.com/allegro/bigcache"
	"go-boilerplate/common/util/logutil"
	"go-boilerplate/core/service/cache"
)

type bigCache struct {
	bigCacheClient *bigcache.BigCache
}

type bigCacheAccessor struct {
	namespace      string
	bigCacheClient *bigcache.BigCache
}

func NewBigCache(bigCacheClient *bigcache.BigCache) cache.CacheService {
	return &bigCache{
		bigCacheClient: bigCacheClient,
	}
}

func (b bigCache) SetNamespace(s string) cache.CacheAccessor {
	return bigCacheAccessor{
		namespace:      s,
		bigCacheClient: b.bigCacheClient,
	}
}

func (r bigCacheAccessor) GetString(ctx context.Context, key string) string {
	var cacheValMap map[string]string

	data, err := r.bigCacheClient.Get(r.namespace)
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

func (r bigCacheAccessor) SetString(ctx context.Context, key string, val string) {
	var cacheValMap = make(map[string]string)

	data, err := r.bigCacheClient.Get(r.namespace)
	if err != nil {
		if err == bigcache.ErrEntryNotFound {
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

	err = r.bigCacheClient.Set(r.namespace, cacheValJson)
	if err != nil {
		logutil.LogError(ctx, "failed when setting cache data", err.Error())
		return
	}

}

func (r bigCacheAccessor) Clean(ctx context.Context) error {
	return r.bigCacheClient.Delete(r.namespace)
}

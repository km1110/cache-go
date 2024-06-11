package adapters

import (
	"time"

	"github.com/patrickmn/go-cache"
)

type MemoryCacheAdapter struct {
	client *cache.Cache
}

func NewMemoryCacheAdapter() *MemoryCacheAdapter {
	return &MemoryCacheAdapter{
		client: cache.New(30*time.Second, 2*time.Second),
	}
}

func (adapter *MemoryCacheAdapter) Get(key string) string {
	val, found := adapter.client.Get(key)
	if !found {
		return ""
	}
	return val.(string)
}

func (adapter *MemoryCacheAdapter) Set(key string, value string) {
	adapter.client.Set(key, value, cache.DefaultExpiration)
}

func (adapter *MemoryCacheAdapter) Delete(key string) {
	adapter.client.Delete(key)
}

package main

import (
	"fmt"
	"time"

	"github.com/km1110/cache-go/adapters"
	"github.com/km1110/cache-go/interfaces"
)

func main() {
	memory_cache_adapter := adapters.NewMemoryCacheAdapter()
	redis_adapter := adapters.NewRedisAdapter()

	totalTime := 0

	for {
		value := interfaces.GetData(memory_cache_adapter, "key")
		if value == "" {
			fmt.Println("Cache miss so get from redis")
			value = interfaces.GetData(redis_adapter, "key")
			interfaces.SetData(memory_cache_adapter, "key", value)
		} else {
			fmt.Println("Cache hit")
		}
		fmt.Println("Value: ", value)

		time.Sleep(3 * time.Second)

		totalTime += 3
		fmt.Println("Total time: ", totalTime)
	}
}

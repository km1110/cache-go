package main

import (
	"fmt"
	"time"

	"github.com/patrickmn/go-cache"
)

func main() {
	c := cache.New(30*time.Second, 2*time.Minute)
	c.Set("key", "value", cache.DefaultExpiration)

	val, found := c.Get("key")
	if found {
		fmt.Println("Found key:", val)
	} else {
		fmt.Println("Key not found")
	}

	time.Sleep(20 * time.Second)
	fmt.Println("After 1 minute")
	_, found = c.Get("key")
	if found {
		fmt.Println("Found key:", val)
	} else {
		fmt.Println("Key not found")
	}
}

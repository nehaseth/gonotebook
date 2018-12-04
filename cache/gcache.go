package main

import (
	"github.com/bluele/gcache"
	"fmt"
	"time"
	"github.com/pkg/errors"
)

type CacheKeyStr string

func main() {
	gc := gcache.New(100).Expiration(time.Second * 5).LRU().LoaderFunc(func(key interface{}) (interface{}, error) {
		fmt.Println("Loading in cache")
		if _, ok := key.(CacheKeyStr); !ok {
			fmt.Println("Invalid key")
			return nil, errors.New("Invalid key")
		}
		return "ok", nil
	}).Build()


	//Test without single set statement :
	gc.Set("key", "ok")
	validKey := CacheKeyStr("test")
	value, _ := gc.Get(validKey)
	fmt.Println("Get:", value)

	// Wait for value to expire
	time.Sleep(time.Second*10)

	value, _ = gc.Get(validKey)
	fmt.Println("Get:", value)
	value, err := gc.Get("key")
	if err != nil {
		panic(err)
	}
	fmt.Println("Get:", value)
}

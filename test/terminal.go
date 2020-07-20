package main

import (
	"fmt"
	"github.com/fideism/golang-wechat/cache"
)

func main() {
	opts := &cache.RedisOpts{
		Host: "127.0.0.1:6379",
	}
	redis := cache.NewRedis(opts)

	fmt.Println(redis)
}

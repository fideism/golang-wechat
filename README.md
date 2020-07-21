# WeChat SDK for Go

## 目录
- [缓存](./cache)
- [基础信息(access_token，js_ticket)](./credential)
- [公众号](./officialaccount)
- [小程序](./miniprogram)

## 快速入门

### 缓存

```go
package main

import	"github.com/fideism/golang-wechat/cache"

func main() {
    redis := &cache.RedisOpts{
            Host:        "127.0.0.1:6379",
            Password:    "111111",
            Database:    1,
            MaxIdle:     5, //最大等待连接中的数量
            MaxActive:   3, //最大连接数据库连接数
            IdleTimeout: 1, //客户端的idle
    }

    cache := cache.NewRedis(redis)
}
```

### 公众号

详细方法见[公众号参考](./officialaccount/README.md)

```go
package main

import (
	"fmt"
	wechat "github.com/fideism/golang-wechat"
	"github.com/fideism/golang-wechat/cache"
	"github.com/fideism/golang-wechat/officialaccount"
	offConfig "github.com/fideism/golang-wechat/officialaccount/config"
)

func main() {
    redis := &cache.RedisOpts{
        Host:     "127.0.0.1:6379",
    }
    
    config := &offConfig.Config{
        AppID:          "xxx",
        AppSecret:      "xxx",
        Token:          "xxx",
        EncodingAESKey: "xxx",
        Cache:          cache.NewRedis(redis),
    }
    
    // 初始化wechat实例，分别调用对应功能模块
    wechat := wechat.NewWechat()
    officail := wechat.GetOfficialAccount(config)
    
    // 单独获得officailAccount实例
    // officail := officialaccount.NewOfficialAccount(config)
    
    token, err := officail.GetAccessToken()
    if err != nil {
        panic(err)
    }
    
    fmt.Println(token)
}

```

### Based On
[silenceper/wechat](https://github.com/silenceper/wechat) 

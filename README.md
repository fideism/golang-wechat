# WeChat SDK for Go

:clap::clap::clap: Golang Wechat SDK 

[![golang](https://img.shields.io/badge/Language-Go-green.svg?style=flat)](https://golang.org)
[![MIT license](http://img.shields.io/badge/license-MIT-brightgreen.svg)](http://opensource.org/licenses/MIT)

## 目录
- [缓存](./cache)
- [基础信息(access_token，js_ticket)](./credential)
- [公众号](./officialaccount)
- [小程序](./miniprogram)
- [开放平台](./openplatform)
- [支付](./pay)

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

### 参数

`Params util.Params`

```go
import "github.com/fideism/golang-wechat/util"

// Params map[string]interface{}
type Params map[string]interface{}

// Set 设置值
func (p Params) Set(k string, v interface{})

// Get 获取值
func (p Params) Get(k string) (v interface{})

// GetString 强制获取k对应的v string类型
func (p Params) GetString(k string) string

// Exists 判断是否存在
func (p Params) Exists(k string) bool

//具体使用
p := util.Params{
    "openid": "xx",
}

//alse can
p.Set("notify_url", "https://github.com/fideism/golang-wechat")
```

### 日志
默认记录`debug`级别日志

可以通过设置系统`LOG_LEVEL`来控制日志记录

### 版本说明

- V1.0.0 初始版本

### Based On :thumbsup:
[silenceper/wechat](https://github.com/silenceper/wechat) 

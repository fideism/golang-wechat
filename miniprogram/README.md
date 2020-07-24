# 微信小程序

[官方文档](https://developers.weixin.qq.com/miniprogram/dev/framework/)

## 包说明

-   analysis 数据分析相关 API

## 快速入门

```go
package main

import (
	wechat "github.com/fideism/golang-wechat"
	"github.com/fideism/golang-wechat/cache"
	"github.com/fideism/golang-wechat/miniprogram"
	miniConfig "github.com/fideism/golang-wechat/miniprogram/config"
)

func main() {
	redis := &cache.RedisOpts{
		Host: "*.*.*.*:port",
	}

	config := &miniConfig.Config{
		AppID:     `appid`,
		AppSecret: `appsecret`,
		Cache:     cache.NewRedis(redis),
	}

	// 初始化wechat实例，分别调用对应功能模块
	wechat := wechat.NewWechat()
	miniprogram := wechat.GetMiniProgram(config)

	// 单独获得officailAccount实例
	// miniprogram := miniprogram.NewMiniProgram(config)
}
```

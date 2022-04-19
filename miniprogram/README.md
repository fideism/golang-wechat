# 微信小程序

[官方文档](https://developers.weixin.qq.com/miniprogram/dev/framework/)

## 包说明

- analysis 数据分析相关 API
- auth 授权
- url 小程序url link
- qrcode 小程序二维码

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
		Host:     "127.0.0.1:6379",
		Database: 0,
	}

	config := &miniConfig.Config{
		AppID:     `xxxx`,
		AppSecret: `xxx`,
		Cache:     cache.NewRedis(redis),
	}

	mini := miniprogram.NewMiniProgram(config)

	fmt.Println(mini)

	rst, err := mini.GetAuth().Code2Session(`093MOqHa19WedC0AnXHa1oTu2B4MOqHi`)
	fmt.Println(err)
	fmt.Println(rst)
}
```

# 微信公众号

[官方文档](https://developers.weixin.qq.com/doc/offiaccount/Getting_Started/Overview.html)

## 目录
- [快速入门](#快速入门)
- [Token](#Token)
- [获取微信服务器IP地址](#获取微信服务器IP地址)
- [清理接口调用频次](#清理接口调用频次)
- [带参数二维码](#带参数二维码)
- [用户](#用户)
- [菜单](#菜单)
- [oauth2网页授权](#oauth2网页授权)
- [素材管理](#素材管理)
- [js-sdk配置](#js-sdk配置)
- [群发消息](#群发消息)
- [模板消息接口](#模板消息接口)
- [获取智能设备的实例](#获取智能设备的实例)
- [数据统计](#数据统计)


## 快速入门

```go
package main

import (
	"github.com/fideism/golang-wechat/cache"
	"github.com/fideism/golang-wechat/officialaccount"
	"github.com/fideism/golang-wechat/officialaccount/config"
)

func main() {
	//设置全局cache，也可以单独为每个操作实例设置
	redis := &cache.RedisOpts{
		Host:     "127.0.0.1:6379",
		Database: 1,
	}

	config := &config.Config{
		AppID:          "xxxx",
		AppSecret:      "xxxx",
		Token:          "xxxx",
		EncodingAESKey: "xxx",
		Cache:          cache.NewRedis(redis),
	}

	officail := officialaccount.NewOfficialAccount(config)
}
```

`content` 获取
```go
func (officialAccount *OfficialAccount) GetContext() *context.Context
```

### Token

```
token, err := officail.GetAccessToken()
```

自定义`token`获取方式
```go
func (officialAccount *OfficialAccount) SetAccessTokenHandle(accessTokenHandle credential.AccessTokenHandle)
```

### 获取微信服务器IP地址

`callback` IP地址

```go
ips, err := officail.GetBasic().GetCallbackIP()
```

`API`接口 IP地址
```go
ips, err := officail.GetBasic().GetAPIDomainIP()
```

### 清理接口调用频次

```go
err:=officialAccount.GetBasic().ClearQuota()
```

### 带参数二维码

- 生成请求参数
```go
import "github.com/fideism/golang-wechat/officialaccount/basic"

//临时 有效时间(秒) 场景值
request := basic.NewTmpQrRequest(300, "scene")
//永久 场景值
request := basic.NewLimitQrRequest("scene")

// QrActionScene QR_SCENE为临时的整型参数值
// QrActionStrScene QR_STR_SCENE为临时的字符串参数值
// QrActionLimitScene QR_LIMIT_SCENE为永久的整型参数值
// QrActionLimitStrScene QR_LIMIT_STR_SCENE为永久的字符串参数值
//手动组装request信息
req := basic.Request{
		ExpireSeconds: 0,
		ActionName:    basic.QrActionScene,
	}
req.ActionInfo.Scene.SceneStr = "scene"
req.ActionInfo.Scene.SceneID = 123
```

- 通过ticket换取二维码
```go
ticket, err := basic.GetQRTicket(req)

qrcode_url := basic.ShowQRCode(ticket)
```

### 用户
```go
func (officialAccount *OfficialAccount) GetUser() *user.User

```
- 用户列表

```go
// 所有用户openids列表
func (user *User) ListAllUserOpenIDs() ([]string, error)

// 批量拉取1000个用户信息
func (user *User) ListUserOpenIDs(nextOpenid ...string) (*OpenidList, error)
```

- 获取用户基本信息(UnionID机制)
```go
func (user *User) GetUserInfo(openID string) (userInfo *Info, err error) 
```

- 对指定用户设置备注名
```go
func (user *User) UpdateRemark(openID, remark string) (err error)
```

### 菜单
```go
func (officialAccount *OfficialAccount) GetMenu() *menu.Menu
```

### oauth2网页授权
```go
func (officialAccount *OfficialAccount) GetOauth() *oauth.Oauth
```

### 素材管理
```go
func (officialAccount *OfficialAccount) GetMaterial() *material.Material
```

### js-sdk配置
```go
func (officialAccount *OfficialAccount) GetJs() *js.Js
```

### 群发消息
```go
func (officialAccount *OfficialAccount) GetBroadcast() *broadcast.Broadcast
```

### 模板消息接口
```go
func (officialAccount *OfficialAccount) GetTemplate() *message.Template
```

### 获取智能设备的实例
```go
func (officialAccount *OfficialAccount) GetDevice() *device.Device
```

### 数据统计
```go
func (officialAccount *OfficialAccount) GetDataCube() *datacube.DataCube
```

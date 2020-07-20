package config

import (
	"github.com/fideism/golang-wechat/cache"
)

//Config config for 小程序
type Config struct {
	AppID     string `json:"app_id"`     //appid
	AppSecret string `json:"app_secret"` //appsecret
	Cache     cache.Cache
}

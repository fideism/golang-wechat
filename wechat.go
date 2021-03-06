package wechat

import (
	"github.com/fideism/golang-wechat/cache"
	"github.com/fideism/golang-wechat/miniprogram"
	miniConfig "github.com/fideism/golang-wechat/miniprogram/config"
	"github.com/fideism/golang-wechat/officialaccount"
	offConfig "github.com/fideism/golang-wechat/officialaccount/config"
	"github.com/fideism/golang-wechat/openplatform"
	openConfig "github.com/fideism/golang-wechat/openplatform/config"
	"github.com/fideism/golang-wechat/pay"
	"github.com/fideism/golang-wechat/pay/config"
)

// Wechat struct
type Wechat struct {
	cache cache.Cache
}

// NewWechat init
func NewWechat() *Wechat {
	return &Wechat{}
}

//SetCache 设置cache
func (w *Wechat) SetCache(cache cache.Cache) {
	w.cache = cache
}

//GetOfficialAccount 获取微信公众号实例
func (w *Wechat) GetOfficialAccount(c *offConfig.Config) *officialaccount.OfficialAccount {
	if c.Cache == nil {
		c.Cache = w.cache
	}
	return officialaccount.NewOfficialAccount(c)
}

// GetMiniProgram 获取小程序的实例
func (w *Wechat) GetMiniProgram(c *miniConfig.Config) *miniprogram.MiniProgram {
	if c.Cache == nil {
		c.Cache = w.cache
	}
	return miniprogram.NewMiniProgram(c)
}

// GetPay 获取微信支付的实例
func (w *Wechat) GetPay(c *config.Config) *pay.Pay {
	return pay.NewPay(c)
}

// GetOpenPlatform 获取微信开放平台的实例
func (w *Wechat) GetOpenPlatform(c *openConfig.Config) *openplatform.OpenPlatform {
	return openplatform.NewOpenPlatform(c)
}

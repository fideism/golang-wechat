package wechat

import (
<<<<<<< HEAD
	"os"

	"github.com/fideism/golang-wechat/cache"
=======
	"github.com/fideism/golang-wechat/cache"
	logger "github.com/fideism/golang-wechat/log"
>>>>>>> ab3c5799693c669ca46f2f759fc28f67a35e865c
	"github.com/fideism/golang-wechat/miniprogram"
	miniConfig "github.com/fideism/golang-wechat/miniprogram/config"
	"github.com/fideism/golang-wechat/officialaccount"
	offConfig "github.com/fideism/golang-wechat/officialaccount/config"
	"github.com/fideism/golang-wechat/openplatform"
	openConfig "github.com/fideism/golang-wechat/openplatform/config"
	"github.com/fideism/golang-wechat/pay"
	payConfig "github.com/fideism/golang-wechat/pay/config"
	log "github.com/sirupsen/logrus"
<<<<<<< HEAD
=======
	"os"
>>>>>>> ab3c5799693c669ca46f2f759fc28f67a35e865c
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
<<<<<<< HEAD
	log.SetFormatter(&log.TextFormatter{})
=======
	log.SetFormatter(&logger.LogstashFormatter{
		Channel: "WeChat",
	})
>>>>>>> ab3c5799693c669ca46f2f759fc28f67a35e865c

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.DebugLevel)
}

// Wechat struct
type Wechat struct {
	cache cache.Cache
}

// NewWechat init
func NewWechat() *Wechat {
	return &Wechat{}
}

//SetCache 设置cache
<<<<<<< HEAD
func (wc *Wechat) SetCache(cahce cache.Cache) {
	wc.cache = cahce
}

//GetOfficialAccount 获取微信公众号实例
func (wc *Wechat) GetOfficialAccount(cfg *offConfig.Config) *officialaccount.OfficialAccount {
	if cfg.Cache == nil {
		cfg.Cache = wc.cache
	}
	return officialaccount.NewOfficialAccount(cfg)
}

// GetMiniProgram 获取小程序的实例
func (wc *Wechat) GetMiniProgram(cfg *miniConfig.Config) *miniprogram.MiniProgram {
	if cfg.Cache == nil {
		cfg.Cache = wc.cache
	}
	return miniprogram.NewMiniProgram(cfg)
}

// GetPay 获取微信支付的实例
func (wc *Wechat) GetPay(cfg *payConfig.Config) *pay.Pay {
	return pay.NewPay(cfg)
}

// GetOpenPlatform 获取微信开放平台的实例
func (wc *Wechat) GetOpenPlatform(cfg *openConfig.Config) *openplatform.OpenPlatform {
	return openplatform.NewOpenPlatform(cfg)
=======
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
func (w *Wechat) GetPay(c *payConfig.Config) *pay.Pay {
	return pay.NewPay(c)
}

// GetOpenPlatform 获取微信开放平台的实例
func (w *Wechat) GetOpenPlatform(c *openConfig.Config) *openplatform.OpenPlatform {
	return openplatform.NewOpenPlatform(c)
>>>>>>> ab3c5799693c669ca46f2f759fc28f67a35e865c
}

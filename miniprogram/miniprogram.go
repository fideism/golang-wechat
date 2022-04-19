package miniprogram

import (
	"github.com/fideism/golang-wechat/credential"
	"github.com/fideism/golang-wechat/miniprogram/analysis"
	"github.com/fideism/golang-wechat/miniprogram/auth"
	"github.com/fideism/golang-wechat/miniprogram/config"
	"github.com/fideism/golang-wechat/miniprogram/context"
	"github.com/fideism/golang-wechat/miniprogram/encryptor"
	"github.com/fideism/golang-wechat/miniprogram/qrcode"
	"github.com/fideism/golang-wechat/miniprogram/subscribe"
	"github.com/fideism/golang-wechat/miniprogram/tcb"
	"github.com/fideism/golang-wechat/miniprogram/url"
)

//MiniProgram 微信小程序相关API
type MiniProgram struct {
	ctx *context.Context
}

//NewMiniProgram 实例化小程序API
func NewMiniProgram(cfg *config.Config) *MiniProgram {
	defaultAkHandle := credential.NewDefaultAccessToken(cfg.AppID, cfg.AppSecret, credential.CacheKeyMiniProgramPrefix, cfg.Cache)
	ctx := &context.Context{
		Config:            cfg,
		AccessTokenHandle: defaultAkHandle,
	}
	return &MiniProgram{ctx}
}

//SetAccessTokenHandle 自定义access_token获取方式
func (miniProgram *MiniProgram) SetAccessTokenHandle(accessTokenHandle credential.AccessTokenHandle) {
	miniProgram.ctx.AccessTokenHandle = accessTokenHandle
}

// GetContext get Context
func (miniProgram *MiniProgram) GetContext() *context.Context {
	return miniProgram.ctx
}

// GetEncryptor  小程序加解密
func (miniProgram *MiniProgram) GetEncryptor() *encryptor.Encryptor {
	return encryptor.NewEncryptor(miniProgram.ctx)
}

//GetAuth 登录/用户信息相关接口
func (miniProgram *MiniProgram) GetAuth() *auth.Auth {
	return auth.NewAuth(miniProgram.ctx)
}

//GetAnalysis 数据分析
func (miniProgram *MiniProgram) GetAnalysis() *analysis.Analysis {
	return analysis.NewAnalysis(miniProgram.ctx)
}

//GetQRCode 小程序码相关API
func (miniProgram *MiniProgram) GetQRCode() *qrcode.QRCode {
	return qrcode.NewQRCode(miniProgram.ctx)
}

//GetTcb 小程序云开发API
func (miniProgram *MiniProgram) GetTcb() *tcb.Tcb {
	return tcb.NewTcb(miniProgram.ctx)
}

//GetSubscribe 小程序订阅消息
func (miniProgram *MiniProgram) GetSubscribe() *subscribe.Subscribe {
	return subscribe.NewSubscribe(miniProgram.ctx)
}

//GetURL 小程序链接
func (miniProgram *MiniProgram) GetURL() *url.URL {
	return url.NewURL(miniProgram.ctx)
}

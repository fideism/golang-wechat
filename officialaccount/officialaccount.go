package officialaccount

import (
	"net/http"

	"github.com/fideism/golang-wechat/credential"
	"github.com/fideism/golang-wechat/officialaccount/basic"
	"github.com/fideism/golang-wechat/officialaccount/broadcast"
	"github.com/fideism/golang-wechat/officialaccount/card"
	"github.com/fideism/golang-wechat/officialaccount/config"
	"github.com/fideism/golang-wechat/officialaccount/context"
	"github.com/fideism/golang-wechat/officialaccount/datacube"
	"github.com/fideism/golang-wechat/officialaccount/device"
	"github.com/fideism/golang-wechat/officialaccount/js"
	"github.com/fideism/golang-wechat/officialaccount/material"
	"github.com/fideism/golang-wechat/officialaccount/menu"
	"github.com/fideism/golang-wechat/officialaccount/message"
	"github.com/fideism/golang-wechat/officialaccount/oauth"
	"github.com/fideism/golang-wechat/officialaccount/server"
	"github.com/fideism/golang-wechat/officialaccount/user"
)

//OfficialAccount 微信公众号相关API
type OfficialAccount struct {
	ctx *context.Context
}

//NewOfficialAccount 实例化公众号API
func NewOfficialAccount(cfg *config.Config) *OfficialAccount {
	defaultAkHandle := credential.NewDefaultAccessToken(cfg.AppID, cfg.AppSecret, credential.CacheKeyOfficialAccountPrefix, cfg.Cache)
	ctx := &context.Context{
		Config:            cfg,
		AccessTokenHandle: defaultAkHandle,
	}
	return &OfficialAccount{ctx: ctx}
}

//SetAccessTokenHandle 自定义access_token获取方式
func (o *OfficialAccount) SetAccessTokenHandle(accessTokenHandle credential.AccessTokenHandle) {
	o.ctx.AccessTokenHandle = accessTokenHandle
}

// GetContext get Context
func (o *OfficialAccount) GetContext() *context.Context {
	return o.ctx
}

// GetBasic qr/url 相关配置
func (o *OfficialAccount) GetBasic() *basic.Basic {
	return basic.NewBasic(o.ctx)
}

// GetMenu 菜单管理接口
func (o *OfficialAccount) GetMenu() *menu.Menu {
	return menu.NewMenu(o.ctx)
}

// GetServer 消息管理：接收事件，被动回复消息管理
func (o *OfficialAccount) GetServer(req *http.Request, writer http.ResponseWriter) *server.Server {
	srv := server.NewServer(o.ctx)
	srv.Request = req
	srv.Writer = writer
	return srv
}

//GetAccessToken 获取access_token
func (o *OfficialAccount) GetAccessToken() (string, error) {
	return o.ctx.GetAccessToken()
}

// GetOauth oauth2网页授权
func (o *OfficialAccount) GetOauth() *oauth.Oauth {
	return oauth.NewOauth(o.ctx)
}

// GetMaterial 素材管理
func (o *OfficialAccount) GetMaterial() *material.Material {
	return material.NewMaterial(o.ctx)
}

// GetJs js-sdk配置
func (o *OfficialAccount) GetJs() *js.Js {
	return js.NewJs(o.ctx)
}

// GetUser 用户管理接口
func (o *OfficialAccount) GetUser() *user.User {
	return user.NewUser(o.ctx)
}

// GetTemplate 模板消息接口
func (o *OfficialAccount) GetTemplate() *message.Template {
	return message.NewTemplate(o.ctx)
}

// GetDevice 获取智能设备的实例
func (o *OfficialAccount) GetDevice() *device.Device {
	return device.NewDevice(o.ctx)
}

//GetBroadcast 群发消息
//TODO 待完善
func (o *OfficialAccount) GetBroadcast() *broadcast.Broadcast {
	return broadcast.NewBroadcast(o.ctx)
}

//GetDataCube 数据统计
func (o *OfficialAccount) GetDataCube() *datacube.DataCube {
	return datacube.NewCube(o.ctx)
}

// GetCard 卡券
func (o *OfficialAccount) GetCard() *card.Card {
	return card.NewCard(o.ctx)
}

// GetMessage 消息管理
func (o *OfficialAccount) GetMessage() *message.Manager {
	return message.NewMessageManager(o.ctx)
}

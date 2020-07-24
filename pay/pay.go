package pay

import (
	"github.com/fideism/golang-wechat/pay/config"
	"github.com/fideism/golang-wechat/pay/notify"
	"github.com/fideism/golang-wechat/pay/order"
	"github.com/fideism/golang-wechat/pay/refund"
	"github.com/fideism/golang-wechat/pay/server"
)

//Pay 微信支付相关API
type Pay struct {
	cfg *config.Config
}

//NewPay 实例化微信支付相关API
func NewPay(cfg *config.Config) *Pay {
	return &Pay{cfg}
}

// GetOrder  下单
func (pay *Pay) GetOrder() *order.Order {
	return order.NewOrder(pay.cfg)
}

// GetNotify  通知
func (pay *Pay) GetNotify() *notify.Notify {
	return notify.NewNotify(pay.cfg)
}

// GetRefund  退款
func (pay *Pay) GetRefund() *refund.Refund {
	return refund.NewRefund(pay.cfg)
}

// GetServer Get Server
func (pay *Pay) GetServer() *server.Server {
	return server.NewServer(pay.cfg)
}

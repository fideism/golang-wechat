package order

import (
	"github.com/fideism/golang-wechat/pay/base"
	"github.com/fideism/golang-wechat/pay/config"
)

// Order 下单相关实例
type Order struct {
	*config.Config
}

// NewOrder 获取下单实例
func NewOrder(c *config.Config) *Order {
	return &Order{c}
}

// Unify 统一下单
func (order *Order) Unify(params base.Params) (*base.Response, error) {
	if !params.Exists("spbill_create_ip") {
		params.Set("spbill_create_ip", base.GetIP())
	}

	params.Sign(order.Config)

	xmlStr, err := base.PostWithoutCert(base.GetUnifyURL(order.Config), params)
	if err != nil {
		return nil, err
	}

	return base.ProcessResponseXML(xmlStr)
}

// MicroPay 付款码支付
func (order *Order) MicroPay(params base.Params) (*base.Response, error) {
	if !params.Exists("spbill_create_ip") {
		params.Set("spbill_create_ip", base.GetIP())
	}

	params.Sign(order.Config)

	xmlStr, err := base.PostWithoutCert(base.GetMicroPayURL(order.Config), params)
	if err != nil {
		return nil, err
	}

	return base.ProcessResponseXML(xmlStr)
}

// Query 查询订单
func (order *Order) Query(params base.Params) (*base.Response, error) {
	params.Sign(order.Config)

	xmlStr, err := base.PostWithoutCert(base.GetOrderQueryURL(order.Config), params)
	if err != nil {
		return nil, err
	}

	return base.ProcessResponseXML(xmlStr)
}

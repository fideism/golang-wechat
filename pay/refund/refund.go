package refund

import (
	"github.com/fideism/golang-wechat/pay/base"
	"github.com/fideism/golang-wechat/pay/config"
	"github.com/fideism/golang-wechat/util"
)

// Refund struct extends context
type Refund struct {
	*config.Config
}

// NewRefund return an instance of refund package
func NewRefund(cfg *config.Config) *Refund {
	return &Refund{cfg}
}

// Refund 退款
func (refund *Refund) Refund(params util.Params, certCfg config.Cert) (*base.Response, error) {
	params = base.Sign(params, refund.Config)

	tsl, err := base.CertTLSConfig(refund.Config.MchID, certCfg)
	if err != nil {
		return nil, err
	}

	xmlStr, err := base.PostWithTSL(base.GetRefundURL(refund.Config), params, tsl)
	if err != nil {
		return nil, err
	}

	return base.ProcessResponseXML(xmlStr)
}

// Query 查询退款
func (refund *Refund) Query(params util.Params) (*base.Response, error) {
	params = base.Sign(params, refund.Config)

	xmlStr, err := base.PostWithoutCert(base.GetRefundQueryURL(refund.Config), params)
	if err != nil {
		return nil, err
	}

	return base.ProcessResponseXML(xmlStr)
}

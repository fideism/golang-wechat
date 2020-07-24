package notify

import (
	"github.com/fideism/golang-wechat/pay/base"
	"github.com/fideism/golang-wechat/pay/config"
)

//Notify 回调
type Notify struct {
	*config.Config
}

//NewNotify new
func NewNotify(cfg *config.Config) *Notify {
	return &Notify{cfg}
}

// 通知成功
func (n *Notify) Success() (*base.Response, error) {
	var params base.Response

	params.ReturnCode = "Success"
	params.ReturnMsg = "OK"

	return &params, nil
}

// 通知不成功
func (n *Notify) Fail(errMsg string) (*base.Response, error) {
	var params base.Response

	params.ReturnCode = "Success"
	params.ReturnMsg = errMsg

	return &params, nil
}

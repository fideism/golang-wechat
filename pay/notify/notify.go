package notify

import (
	"encoding/xml"

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

// PayNotifyXML 支付通知解析
type PayNotifyXML struct {
	Appid         string `xml:"appid" json:"appid"`
	Attach        string `xml:"attach" json:"attach"`
	BankType      string `xml:"bank_type" json:"bank_type"`
	CashFee       string `xml:"cash_fee" json:"cash_fee"`
	FeeType       string `xml:"fee_type" json:"fee_type"`
	IsSubscribe   string `xml:"is_subscribe" json:"is_subscribe"`
	MchID         string `xml:"mch_id" json:"mch_id"`
	NonceStr      string `xml:"nonce_str" json:"nonce_str"`
	Openid        string `xml:"openid" json:"openid"`
	OutTradeNo    string `xml:"out_trade_no" json:"out_trade_no"`
	ResultCode    string `xml:"result_code" json:"result_code"`
	ReturnCode    string `xml:"return_code" json:"return_code"`
	Sign          string `xml:"sign" json:"sign"`
	SubMchID      string `xml:"sub_mch_id" json:"sub_mch_id"`
	TimeEnd       string `xml:"time_end" json:"time_end"`
	TotalFee      string `xml:"total_fee" json:"total_fee"`
	TradeType     string `xml:"trade_type" json:"trade_type"`
	TransactionID string `xml:"transaction_id" json:"transaction_id"`
}

// AnalysisPayNotify 解析支付通知回调
func (n *Notify) AnalysisPayNotify(context []byte) (*PayNotifyXML, error) {
	var response PayNotifyXML
	if err := xml.Unmarshal(context, &response); nil != err {
		return nil, err
	}

	return &response, nil
}

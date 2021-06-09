package notify

import (
	"crypto/aes"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"encoding/xml"
	"errors"
	"fmt"
	"strings"

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

// Success 通知成功
func (n *Notify) Success() (*base.Response, error) {
	var params base.Response

	params.ReturnCode = "Success"
	params.ReturnMsg = "OK"

	return &params, nil
}

// Fail 通知不成功
func (n *Notify) Fail(errMsg string) (*base.Response, error) {
	var params base.Response

	params.ReturnCode = "Success"
	params.ReturnMsg = errMsg

	return &params, nil
}

// PayNotify 支付通知解析
type PayNotify struct {
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
func (n *Notify) AnalysisPayNotify(context []byte) (*PayNotify, error) {
	var response PayNotify
	if err := xml.Unmarshal(context, &response); nil != err {
		return nil, err
	}

	return &response, nil
}

// RefundNotifyXML 退款xml
type RefundNotifyXML struct {
	ReturnCode string `xml:"return_code,omitempty" json:"return_code,omitempty"`
	ReturnMsg  string `xml:"return_msg,omitempty" json:"return_msg,omitempty"`
	Appid      string `xml:"appid,omitempty" json:"appid,omitempty"`
	SubAppid   string `xml:"sub_appid,omitempty" json:"sub_appid,omitempty"`
	MchId      string `xml:"mch_id,omitempty" json:"mch_id,omitempty"`
	SubMchId   string `xml:"sub_mch_id,omitempty" json:"sub_mch_id,omitempty"`
	NonceStr   string `xml:"nonce_str,omitempty" json:"nonce_str,omitempty"`
	ReqInfo    string `xml:"req_info,omitempty" json:"req_info,omitempty"`
}

// RefundNotifyDetail 退款详细信息
type RefundNotifyDetail struct {
	TransactionId       string `xml:"transaction_id,omitempty" json:"transaction_id,omitempty"`
	OutTradeNo          string `xml:"out_trade_no,omitempty" json:"out_trade_no,omitempty"`
	RefundId            string `xml:"refund_id,omitempty" json:"refund_id,omitempty"`
	OutRefundNo         string `xml:"out_refund_no,omitempty" json:"out_refund_no,omitempty"`
	TotalFee            string `xml:"total_fee,omitempty" json:"total_fee,omitempty"`
	SettlementTotalFee  string `xml:"settlement_total_fee,omitempty" json:"settlement_total_fee,omitempty"`
	RefundFee           string `xml:"refund_fee,omitempty" json:"refund_fee,omitempty"`
	SettlementRefundFee string `xml:"settlement_refund_fee,omitempty" json:"settlement_refund_fee,omitempty"`
	RefundStatus        string `xml:"refund_status,omitempty" json:"refund_status,omitempty"`
	SuccessTime         string `xml:"success_time,omitempty" json:"success_time,omitempty"`
	RefundRecvAccout    string `xml:"refund_recv_accout,omitempty" json:"refund_recv_accout,omitempty"`
	RefundAccount       string `xml:"refund_account,omitempty" json:"refund_account,omitempty"`
	RefundRequestSource string `xml:"refund_request_source,omitempty" json:"refund_request_source,omitempty"`
}

// RefundNotify 退款解析返回
type RefundNotify struct {
	XML    RefundNotifyXML    `json:"xml"`
	Detail RefundNotifyDetail `json:"detail"`
}

// AnalysisRefundNotify 解析退款通知回调
func (n *Notify) AnalysisRefundNotify(context []byte) (*RefundNotify, error) {
	var responseDetail RefundNotifyDetail
	var responseXML RefundNotifyXML
	if err := xml.Unmarshal(context, &responseXML); nil != err {
		return nil, err
	}

	if len(responseXML.ReqInfo) == 0 {
		return nil, errors.New("refund xml req info nil")
	}

	encryption, encryptionErr := base64.StdEncoding.DecodeString(responseXML.ReqInfo)
	if encryptionErr != nil {
		return nil, fmt.Errorf("req info decode err:%s", encryptionErr.Error())
	}

	md := md5.New()
	md.Write([]byte(n.Key))
	apiKey := strings.ToLower(hex.EncodeToString(md.Sum(nil)))
	if len(encryption)%aes.BlockSize != 0 {
		return nil, errors.New("encryption data error")
	}

	block, blockErr := aes.NewCipher([]byte(apiKey))
	if blockErr != nil {
		return nil, fmt.Errorf("block api key err:%s", blockErr.Error())
	}

	blockSize := block.BlockSize()

	cipherErr := func(input, output []byte) error {
		if len(output)%blockSize != 0 {
			return errors.New("cipher: input blocks error")
		}
		if len(input) < len(output) {
			return errors.New("cipher: output blocks error")
		}
		for len(output) > 0 {
			block.Decrypt(input, output[:blockSize])
			output = output[blockSize:]
			input = input[blockSize:]
		}
		return nil
	}(encryption, encryption)
	if cipherErr != nil {
		return nil, errors.New("cipher error")
	}

	var encryptionData []byte
	// 过滤补全byte
	length := len(encryption)
	byteLength := int(encryption[length-1])
	if byteLength <= 16 {
		encryptionData = encryption[:(length - byteLength)]
	} else {
		encryptionData = encryption
	}

	if err := xml.Unmarshal(encryptionData, &responseDetail); err != nil {
		return nil, fmt.Errorf("xml.Unmarshal err:%s", err.Error())
	}

	return &RefundNotify{
		XML:    responseXML,
		Detail: responseDetail,
	}, nil
}

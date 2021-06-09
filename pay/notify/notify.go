package notify

import (
	"crypto/aes"
	"crypto/cipher"
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

// AnalysisRefundNotify 解析退款通知回调
func (n *Notify) AnalysisRefundNotify(context []byte) (*RefundNotifyXML, error) {
	var response RefundNotifyXML
	if err := xml.Unmarshal(context, &response); nil != err {
		return nil, err
	}

	return &response, nil
}

// DecryptRefundNotifyReqInfo 解密微信退款异步通知的加密数据
//	reqInfo：gopay.ParseRefundNotify() 方法获取的加密数据 req_info
//	apiKey：API秘钥值
//	返回参数refundNotify：RefundNotify请求的加密数据
//	返回参数err：错误信息
//	文档：https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_16&index=10
func DecryptRefundNotifyReqInfo(reqInfo, apiKey string) (refundNotify *RefundNotifyDetail, err error) {
	if len(reqInfo) == 0 || len(apiKey) == 0 {
		return nil, errors.New("reqInfo or apiKey is null")
	}
	var (
		encryptionB, bs []byte
		block           cipher.Block
		blockSize       int
	)
	if encryptionB, err = base64.StdEncoding.DecodeString(reqInfo); err != nil {
		return nil, err
	}
	h := md5.New()
	h.Write([]byte(apiKey))
	key := strings.ToLower(hex.EncodeToString(h.Sum(nil)))
	if len(encryptionB)%aes.BlockSize != 0 {
		return nil, errors.New("encryptedData is error")
	}
	if block, err = aes.NewCipher([]byte(key)); err != nil {
		return nil, err
	}
	blockSize = block.BlockSize()

	err = func(dst, src []byte) error {
		if len(src)%blockSize != 0 {
			return errors.New("crypto/cipher: input not full blocks")
		}
		if len(dst) < len(src) {
			return errors.New("crypto/cipher: output smaller than input")
		}
		for len(src) > 0 {
			block.Decrypt(dst, src[:blockSize])
			src = src[blockSize:]
			dst = dst[blockSize:]
		}
		return nil
	}(encryptionB, encryptionB)
	if err != nil {
		return nil, err
	}

	bs = PKCS7UnPadding(encryptionB)
	refundNotify = new(RefundNotifyDetail)
	if err = xml.Unmarshal(bs, refundNotify); err != nil {
		return nil, fmt.Errorf("xml.Unmarshal(%s)：%w", string(bs), err)
	}
	return
}

// 解密填充模式（去除补全码） PKCS7UnPadding
// 解密时，需要在最后面去掉加密时添加的填充byte
func PKCS7UnPadding(origData []byte) (bs []byte) {
	length := len(origData)
	unPaddingNumber := int(origData[length-1]) // 找到Byte数组最后的填充byte 数字
	if unPaddingNumber <= 16 {
		bs = origData[:(length - unPaddingNumber)] // 只截取返回有效数字内的byte数组
	} else {
		bs = origData
	}
	return
}

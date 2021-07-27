package base

import (
	"bytes"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"sort"
	"strings"

	logger "github.com/fideism/golang-wechat/log"
	"github.com/fideism/golang-wechat/pay/config"
	"github.com/fideism/golang-wechat/util"
	"github.com/sirupsen/logrus"
)

// Params map[string]interface{}
type Params util.Params

// AppendConfig 增加 基础config参数
func AppendConfig(p util.Params, config *config.Config) util.Params {
	p.Set("appid", config.AppID).
		Set("mch_id", config.MchID)

	return p
}

// SignParams 签名自定义params
func SignParams(p util.Params) string {
	var keys []string
	for k := range p {
		keys = append(keys, k)
	}

	// 由于切片的元素顺序是不固定，所以这里强制给切片元素加个顺序
	sort.Strings(keys)

	//创建字符缓冲
	var buf bytes.Buffer
	for _, k := range keys {
		buf.WriteString(k)
		buf.WriteString(`=`)
		buf.WriteString(p.GetString(k))
		buf.WriteString(`&`)
	}

	return makeSign(p, buf)
}

// Sign 生成签名
func Sign(p util.Params, config *config.Config) util.Params {
	p = AppendConfig(p, config)
	p.Set("nonce_str", util.RandomStr(32))

	if !p.Exists("sign_type") {
		p.Set("sign_type", "MD5")
	}

	// 创建切片
	var keys = make([]string, 0, len(p))
	// 遍历签名参数
	for k := range p {
		if k != "sign" && k != "key" { // 排除 sign, key 字段
			keys = append(keys, k)
		}
	}

	// 由于切片的元素顺序是不固定，所以这里强制给切片元素加个顺序
	sort.Strings(keys)

	//创建字符缓冲
	var buf bytes.Buffer
	for _, k := range keys {
		buf.WriteString(k)
		buf.WriteString(`=`)
		buf.WriteString(p.GetString(k))
		buf.WriteString(`&`)
	}

	// 加入apiKey作加密密钥
	buf.WriteString(`key=`)
	buf.WriteString(config.Key)

	p.Set("sign", makeSign(p, buf))

	logger.Entry().WithFields(logrus.Fields{
		"sign string": buf.String(),
		"params":      p,
	}).Info("支付签名加密")

	return p
}

func makeSign(params util.Params, buffer bytes.Buffer) string {
	var sign string

	if params.GetString("sign_type") == "MD5" {
		dataMd5 := md5.Sum(buffer.Bytes())

		sign = hex.EncodeToString(dataMd5[:]) //需转换成切片
	}

	if params.GetString("sign_type") == "HMAC-SHA256" {
		h := hmac.New(sha256.New, []byte(params.GetString("key")))
		h.Write(buffer.Bytes())
		dataSha256 := h.Sum(nil)
		sign = hex.EncodeToString(dataSha256[:])
	}

	return strings.ToUpper(sign)
}

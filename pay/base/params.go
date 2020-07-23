package base

import (
	"bytes"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"github.com/fideism/golang-wechat/pay/config"
	"github.com/fideism/golang-wechat/util"
	"sort"
	"strconv"
	"strings"
)

// Params map[string]interface{}
type Params map[string]interface{}

// Set 设置值
func (p Params) Set(k string, v interface{}) Params {
	p[k] = v

	return p
}

// Get 获取值
func (p Params) Get(k string) (v interface{}) {
	v, _ = p[k]

	return
}

// GetString 强制获取k对应的v string类型
func (p Params) GetString(k string) string {
	v, _ := p[k]

	return InterfaceToString(v)
}

// Exists 判断是否存在
func (p Params) Exists(k string) bool {
	_, ok := p[k]

	return ok
}

// AppendConfig 增加 基础config参数
func (p Params) AppendConfig(config *config.Config) Params {
	p.Set("appid", config.AppID).
		Set("mch_id", config.MchID)

	return p
}

// Sign 生成签名
func (p Params) Sign(config *config.Config) Params {
	p.AppendConfig(config)
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

	return p
}

func makeSign(params Params, buffer bytes.Buffer) string {
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

// InterfaceToString 不定类型强制装换为string
func InterfaceToString(value interface{}) string {
	var key string
	if value == nil {
		return key
	}

	switch value.(type) {
	case float64:
		ft := value.(float64)
		key = strconv.FormatFloat(ft, 'f', -1, 64)
	case float32:
		ft := value.(float32)
		key = strconv.FormatFloat(float64(ft), 'f', -1, 64)
	case int:
		it := value.(int)
		key = strconv.Itoa(it)
	case uint:
		it := value.(uint)
		key = strconv.Itoa(int(it))
	case int8:
		it := value.(int8)
		key = strconv.Itoa(int(it))
	case uint8:
		it := value.(uint8)
		key = strconv.Itoa(int(it))
	case int16:
		it := value.(int16)
		key = strconv.Itoa(int(it))
	case uint16:
		it := value.(uint16)
		key = strconv.Itoa(int(it))
	case int32:
		it := value.(int32)
		key = strconv.Itoa(int(it))
	case uint32:
		it := value.(uint32)
		key = strconv.Itoa(int(it))
	case int64:
		it := value.(int64)
		key = strconv.FormatInt(it, 10)
	case uint64:
		it := value.(uint64)
		key = strconv.FormatUint(it, 10)
	case string:
		key = value.(string)
	case []byte:
		key = string(value.([]byte))
	default:
		newValue, _ := json.Marshal(value)
		key = string(newValue)
	}

	return key
}

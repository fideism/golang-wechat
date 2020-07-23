# 微信支付

[官方文档](https://pay.weixin.qq.com/wiki/doc/api/index.html)



### 请求参数

支付内部调用参数传递类型`Params  map[string]interface{}` 
```go
import "github.com/fideism/golang-wechat/util"

// Params map[string]interface{}
type Params map[string]interface{}

// Set 设置值
func (p Params) Set(k string, v interface{})

// Get 获取值
func (p Params) Get(k string) (v interface{})

// GetString 强制获取k对应的v string类型
func (p Params) GetString(k string) string

// Exists 判断是否存在
func (p Params) Exists(k string) bool

//具体使用
p := util.Params{
		"openid":       "xx",
}

//alse can
p.Set("notify_url", "https://github.com/fideism/golang-wechat")
```

### 响应参数
```go
import "github.com/fideism/golang-wechat/pay/base"

type Response struct {
	ReturnCode string       `json:"return_code"`
	ReturnMsg  string       `json:"return_msg"`
	Data       base.Params  `json:"data"`
}
```

## 实例化支付对象

```go
import (
	"github.com/fideism/golang-wechat/pay"
	"github.com/fideism/golang-wechat/pay/config"
)

payment := pay.NewPay(&config.Config{
		Sandbox: false,
		AppID:   "wxd12a9416bb9b87fc",
		MchID:   "1480756832",
		Key:     "84e5161b71bec1ce9f3a104a2c602f6d",
	})
```

### 统一下单

`spbill_create_ip` 如果外部传入参数没有，后续调用会自动获取

`sign_type` 外部不传入，默认`MD5`

```go
p := util.Params{
		"out_trade_no": "202007230001",
		"total_fee":    1,
		"body":         "测试支付统一下单",
		"time_start":   "20200723091010",
		"time_expire":  "20200724091010",
		"trade_type":   "JSAPI",
	}

response, err := payment.GetOrder().Unify(p)

// 返回
//{FAIL appid和mch_id不匹配 map[return_code:FAIL return_msg:appid和mch_id不匹配]}
//{SUCCESS OK map[result_code:SUCCESS return_code:SUCCESS return_msg:OK  trade_type:JSAPI.......]}
```

- 下单类型
```go
//APP支付 JSAPI支付 扫码支付 H5支付 小程序支付
func (order *Order) Unify(params base.Params) (*base.Response, error)
// MicroPay 付款码支付
func (order *Order) MicroPay(params base.Params) (*base.Response, error)
```

### 查询订单
```go
params := util.Params{
    "transaction_id": "4200000235201812131594207984",
}

func (order *Order) Query(params base.Params) (*base.Response, error)

payment.GetOrder().Query()
```

### 退款

`order.`

```go
github.com/fideism/golang-wechat/pay/config

p := util.Params{
    "sub_mch_id":     "1512175241",
    "transaction_id": "4200000235201812131594207984",
    "out_refund_no":  "202007230001111",
    "total_fee":      1,
    "refund_fee":     1,
}

// 证书绝对路径
cert := config.Cert{
    Path: "/path/apiclient_cert.p12", 
}

func (order *Order) Refund(params base.Params, cert config.Cert) (*base.Response, error)

payment.GetRefund().Refund()
```

### 查询退款
```go
p := util.Params{
    "sub_mch_id": "1512175241",
    "refund_id":  "50000701192019070910499634214",
}

func (refund *Refund) Query(params base.Params) (*base.Response, error)

payment.GetRefund().Query()
```
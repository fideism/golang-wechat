# 微信公众号

[官方文档](https://developers.weixin.qq.com/doc/offiaccount/Getting_Started/Overview.html)

## 目录
- [快速入门](#快速入门)
- [Token](#Token)
- [获取微信服务器IP地址](#获取微信服务器IP地址)
- [清理接口调用频次](#清理接口调用频次)
- [带参数二维码](#带参数二维码)
- [用户](#用户)
- [菜单](#菜单)
- [oauth2网页授权](#oauth2网页授权)
- [素材管理](#素材管理)
- [卡券](#卡券)
- [js-sdk配置](#js-sdk配置)
- [群发消息](#群发消息)
- [模板消息接口](#模板消息接口)
- [获取智能设备的实例](#获取智能设备的实例)
- [数据统计](#数据统计)


## 快速入门

```go
package main

import (
	"github.com/fideism/golang-wechat/cache"
	"github.com/fideism/golang-wechat/officialaccount"
	"github.com/fideism/golang-wechat/officialaccount/config"
)

func main() {
	//设置全局cache，也可以单独为每个操作实例设置
	redis := &cache.RedisOpts{
		Host:     "127.0.0.1:6379",
		Database: 1,
	}

	config := &config.Config{
		AppID:          "xxxx",
		AppSecret:      "xxxx",
		Token:          "xxxx",
		EncodingAESKey: "xxx",
		Cache:          cache.NewRedis(redis),
	}

	officail := officialaccount.NewOfficialAccount(config)
}
```

`content` 获取
```go
func (officialAccount *OfficialAccount) GetContext() *context.Context
```

### Token

```
token, err := officail.GetAccessToken()
```

自定义`token`获取方式
```go
func (officialAccount *OfficialAccount) SetAccessTokenHandle(accessTokenHandle credential.AccessTokenHandle)
```

### 获取微信服务器IP地址

`callback` IP地址

```go
ips, err := officail.GetBasic().GetCallbackIP()
```

`API`接口 IP地址
```go
ips, err := officail.GetBasic().GetAPIDomainIP()
```

### 清理接口调用频次

```go
err:=officialAccount.GetBasic().ClearQuota()
```

### 带参数二维码

- 生成请求参数
```go
import "github.com/fideism/golang-wechat/officialaccount/basic"

//临时 有效时间(秒) 场景值
request := basic.NewTmpQrRequest(300, "scene")
//永久 场景值
request := basic.NewLimitQrRequest("scene")

// QrActionScene QR_SCENE为临时的整型参数值
// QrActionStrScene QR_STR_SCENE为临时的字符串参数值
// QrActionLimitScene QR_LIMIT_SCENE为永久的整型参数值
// QrActionLimitStrScene QR_LIMIT_STR_SCENE为永久的字符串参数值
//手动组装request信息
req := basic.Request{
		ExpireSeconds: 0,
		ActionName:    basic.QrActionScene,
	}
req.ActionInfo.Scene.SceneStr = "scene"
req.ActionInfo.Scene.SceneID = 123
```

- 通过ticket换取二维码
```go
ticket, err := basic.GetQRTicket(req)

qrcode_url := basic.ShowQRCode(ticket)
```

### 用户
```go
func (officialAccount *OfficialAccount) GetUser() *user.User

```
- 用户列表

```go
// 所有用户openids列表
func (user *User) ListAllUserOpenIDs() ([]string, error)

// 批量拉取1000个用户信息
func (user *User) ListUserOpenIDs(nextOpenid ...string) (*OpenidList, error)
```

- 获取用户基本信息(UnionID机制)
```go
func (user *User) GetUserInfo(openID string) (userInfo *Info, err error) 
```

- 对指定用户设置备注名
```go
func (user *User) UpdateRemark(openID, remark string) (err error)
```

### 菜单
```go
func (officialAccount *OfficialAccount) GetMenu() *menu.Menu
```

### oauth2网页授权
```go
func (officialAccount *OfficialAccount) GetOauth() *oauth.Oauth
```

### 素材管理
```go
func (officialAccount *OfficialAccount) GetMaterial() *material.Material
```

- 上传图片
```go
func (material *Material) UploadImage(filename string) (url string, err error)

// path 图片绝对路径
url, err := officail.GetMaterial().UploadImage(path)
```

### 卡券
```go
func (officialAccount *OfficialAccount) GetCard() *card.Card
```

- 颜色
```go

type Colors struct {
	Name string `json:"name"`
	Value string `json:"value"`
}

// res []Colors
res, err := officail.GetCard().GetColors()
```


- 卡券开放类目查询接口
```go
type Category struct {
	PrimaryCategoryID int64  `json:"primary_category_id"`
	CategoryName      string `json:"category_name"`
	SecondaryCategory []struct {
		SecondaryCategoryID     int64    `json:"secondary_category_id"`
		CategoryName            string   `json:"category_name"`
		NeedQualificationStuffs []string `json:"need_qualification_stuffs"`
		CanChoosePrepaidCard    int64    `json:"can_choose_prepaid_card"`
		CanChoosePaymentCard    int64    `json:"can_choose_payment_card"`
	} `json:"secondary_category"`
}

// res []Category
res, _ := officail.GetCard().GetApplyProtocol()
```

- 创建卡券

所有创建卡券信息都调用该接口就行，传入不同的`card_type` 以及对应卡券所需要的字段信息`map[string]interface{}{.....}`

```go
// 示例创建一张会员卡
attrs := map[string]interface{}{
    "background_pic_url" : cardImage,
    "prerogative" : "可参与丰富的会员专享活动，详情参看相关活动页面",
    "supply_bonus" : false,//显示积分
    "supply_balance" : false,//是否支持储值
    "wx_activate" : true,
    "custom_field1" : map[string]string{
        "name_type" : "FIELD_NAME_TYPE_COUPON",
        "url" : "https://github.com/fideism/golang-wechat",
        "name" : "优惠券",
    },
    "custom_field2" : map[string]string{
        "name_type" : "FIELD_NAME_TYPE_MILEAGE",
        "url" : "",
        "name" : "余额",
    },
    "base_info" : map[string]interface{}{
        "logo_url" : logoImage,
        "code_type" : "CODE_TYPE_NONE",
        "brand_name" : "Kparty",
        "title" : "会员卡",
        "color" : "Color030",
        "notice" : "结账时出示会员卡",
        "description" : "会员权益不可与其它优惠同享（详见活动页面说明）\n结账时出示会员卡，不错过会员权益！",
        "sku" : map[string]int64{
            "quantity": 100000000,
        },
        "date_info" : map[string]string{
            "type": "DATE_TYPE_PERMANENT",
        },
        "use_custom_code" : false,
        "get_limit" : 1,
        "can_give_friend" : false,
        "bind_openid" : false,
        "center_title" : "出示会员码",
        "center_sub_title" : "",
        "center_url" : "https://github.com/fideism/golang-wechat",
        "service_phone" : "028-123456789",
        "custom_url_name" : "领取优惠券",
        "custom_url" : "https://github.com/fideism/golang-wechat",
    },
}

// Groupon 团购券类型
// Groupon Type = "GROUPON"
// Cash 代金券类型
// Cash Type = "CASH"
// Discount 折扣券类型
// Discount Type = "DISCOUNT"
// Gift 兑换券类型
// Gift Type = "GIFT"
// GeneralCoupon 优惠券类型。
// GeneralCoupon Type = "GENERAL_COUPON"
// MemberCard 会员卡
// MemberCard Type = "MEMBER_CARD"
// GeneralCard 礼品卡
// GeneralCard Type = "GENERAL_CARD"
cardID, err := officail.GetCard().CreateCard(card.MemberCard, attrs)
```

- 查看卡券详情
```go
// 卡券ID  返回 map[string]interface{}
card, err := officail.GetCard().GetCard("xxxxx")
```

- 修改卡券信息
```go
// 卡券ID 卡券类型 卡券其他字段参考新增卡券接口
// 返回 是否提交审核，false为修改后不会重新提审，true为修改字段后重新提审，该卡券的状态变为审核中
func (card *Card) UpdateCard(cardID string, t Type, attrs interface{}) (check bool, err error)
```

- 删除卡券
```go
// 卡券ID 
func (card *Card) DeleteCard(cardID string) (err error)
```

- 批量查询卡券列表
```go
// BatchGetRequest 批量查询卡券列表 请求参数
type BatchGetRequest struct {
	Offset     int    `json:"offset"`
	Count      int    `json:"count"`
	StatusList Status `json:"status_list"`
}

// BatchCardList 批量查询卡券列表 返回结果
type BatchCardList struct {
	CardIdList []string `json:"card_id_list"`
	TotalNum   int      `json:"total_num"`
}

func (card *Card) BatchGet(req BatchGetRequest) (res BatchCardList, err error)
```

- 设置会员卡开卡字段接口
```go
// 示例
attrs := map[string]interface{}{
    "required_form": map[string]interface{}{
        "common_field_id_list": []string{"USER_FORM_INFO_FLAG_NAME", "USER_FORM_INFO_FLAG_MOBILE"},
    },
    "optional_form": map[string]interface{}{
        "common_field_id_list": []string{"USER_FORM_INFO_FLAG_BIRTHDAY", "USER_FORM_INFO_FLAG_SEX"},
    },
}

// cardid 卡ID  map[string]interface{} 详细字段信息
func (card *Card) SetActivateUserForm(cardID string, attrs map[string]interface{}) (err error)
```

- 创建卡券二维码
```go
attrs := map[string]interface{}{
    "action_name": "QR_CARD",
    "action_info": map[string]interface{}{
        "card": map[string]interface{}{
            "card_id":  "xxxxx",
            "outer_id": "go",
        },
    },
}

type Qrcode struct {
	Ticket        string `json:"ticket"`
	ExpireSeconds int64  `json:"expire_seconds"`
	URL           string `json:"url"`
	ShowQrcodeURL string `json:"show_qrcode_url"`
}

// map[string]interface{} 详细字段信息
func (card *Card) CreateCardQrcode(attr map[string]interface{}) (res Qrcode, err error)
```

### js-sdk配置
```go
func (officialAccount *OfficialAccount) GetJs() *js.Js
```

### 群发消息
```go
func (officialAccount *OfficialAccount) GetBroadcast() *broadcast.Broadcast
```

### 模板消息接口
```go
func (officialAccount *OfficialAccount) GetTemplate() *message.Template
```

### 获取智能设备的实例
```go
func (officialAccount *OfficialAccount) GetDevice() *device.Device
```

### 数据统计
```go
func (officialAccount *OfficialAccount) GetDataCube() *datacube.DataCube
```

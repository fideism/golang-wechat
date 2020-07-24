# 微信公众号

[官方文档](https://developers.weixin.qq.com/doc/offiaccount/Getting_Started/Overview.html)

## 目录
- [快速入门](#快速入门)
- [Token](#Token)
- [基础接口](#基础接口)
    - [获取微信服务器IP地址](#获取微信服务器IP地址)
    - [清理接口调用频次](#清理接口调用频次)
    - [带参数二维码](#带参数二维码)
    - [短链接](#短链接)
- [用户](#用户)
    - [用户列表](#用户列表)
    - [获取用户基本信息](#获取用户基本信息)
    - [对指定用户设置备注名](#对指定用户设置备注名)
    - [新增标签](#新增标签)
    - [标签列表](#标签列表)
    - [修改标签](#修改标签)
    - [删除标签](#删除标签)
    - [获取标签下粉丝列表](#获取标签下粉丝列表)
    
- [菜单](#菜单)
- [oauth2网页授权](#oauth2网页授权)
- [素材管理](#素材管理)
    - [上传图文消息内的图片](#上传图文消息内的图片)
    - [新增临时素材](#新增临时素材)
    - [获取临时素材](#获取临时素材)
    - [新增永久素材](#新增永久素材)
    - [永久图文素材](#永久图文素材)
    - [永久视频素材](#永久视频素材)
    - [其他类型永久素材](#其他类型永久素材)
    - [删除永久素材](#删除永久素材)
    - [删除永久素材](#删除永久素材)
    - [批量获取永久素材](#批量获取永久素材)
    - [素材总数](#素材总数)
- [卡券](#卡券)
    - [颜色](#颜色)
    - [卡券开放类目查询接口](#卡券开放类目查询接口)
    - [设置白名单](#设置白名单)
    - [创建卡券](#创建卡券)
    - [查看卡券详情](#查看卡券详情)
    - [修改卡券信息](#修改卡券信息)
    - [删除卡券](#删除卡券)
    - [批量查询卡券列表](#批量查询卡券列表)
    - [设置会员卡开卡字段接口](#设置会员卡开卡字段接口)
    - [创建卡券二维码](#创建卡券二维码)
- [消息](#消息)
- [群发消息](#群发消息)
    - [群发消息-发送](#发送群消息)
    - [群发消息-删除](#删除群消息)
    - [群发消息-预览](#预览群消息)
    - [群发消息-状态](#群消息状态)
    - [群发消息-速度](#群消息速度)
- [模板消息接口](#模板消息接口)
- [js-sdk配置](#js-sdk配置)
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

### 基础接口


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
ticket, err := officail.GetBasic().GetQRTicket(req)

qrcode_url := officail.GetBasic().ShowQRCode(ticket)
```

### 短链接
```go
url, err := officail.GetBasic().GetShortURL("https://github.com/fideism/golang-wechat")
```

### 用户
```go
func (officialAccount *OfficialAccount) GetUser() *user.User

```

### 用户列表

```go
// 所有用户openids列表
func (user *User) ListAllUserOpenIDs() ([]string, error)

// 批量拉取1000个用户信息
func (user *User) ListUserOpenIDs(nextOpenid ...string) (*OpenidList, error)
```

### 获取用户基本信息
```go
func (user *User) GetUserInfo(openID string) (userInfo *Info, err error) 
```

### 对指定用户设置备注名
```go
func (user *User) UpdateRemark(openID, remark string) (err error)
```

### 新增标签
```go
type Tag struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Count int64  `json:"count"`
}

func (user *User) CreateTag(name string) (tag Tag, err error)
```

### 标签列表
```go
type Tag struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Count int64  `json:"count"`
}

func (user *User) TagList() (tags []Tag, err error)
```

### 修改标签
```go
func (user *User) UpdateTag(tagID int, name string) (err error)
```

### 删除标签
```go
func (user *User) DeleteTag(tagID int) (err error)
```

### 获取标签下粉丝列表
```go
type TagUser struct {
	Count int `json:"count"`
	Data  struct {
		Openid []string `json:"openid"`
	} `json:"data"`
	NextOpenID string `json:"next_openid"`
}

// tagid 标签ID  next_openid
func (user *User) TagUserList(tagID int, openid string) (res TagUser, err error)
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

### 上传图文消息内的图片
```go
func (material *Material) UploadImage(filename string) (url string, err error)

// path 图片绝对路径
url, err := officail.GetMaterial().UploadImage(path)
```

### 新增临时素材
```go
//MediaTypeImage 媒体文件:图片
MediaTypeImage MediaType = "image"
//MediaTypeVoice 媒体文件:声音
MediaTypeVoice MediaType = "voice"
//MediaTypeVideo 媒体文件:视频
MediaTypeVideo MediaType = "video"
//MediaTypeThumb 媒体文件:缩略图
MediaTypeThumb MediaType = "thumb"
//Media 临时素材上传返回信息
type Media struct {
    ErrCode int64  `json:"errcode"`
    ErrMsg  string `json:"errmsg"`
    Type         MediaType `json:"type"`
    MediaID      string    `json:"media_id"`
    ThumbMediaID string    `json:"thumb_media_id"`
    CreatedAt    int64     `json:"created_at"`
}

func (material *Material) MediaUpload(mediaType MediaType, filename string) (media Media, err error)
```

### 获取临时素材
```go
func (material *Material) GetMediaURL(mediaID string) (mediaURL string, err error)
```

### 永久图文素材
```go
//永久图文素材
type Article struct {
	Title            string `json:"title"`
	ThumbMediaID     string `json:"thumb_media_id"`
	ThumbURL         string `json:"thumb_url"`
	Author           string `json:"author"`
	Digest           string `json:"digest"`
	ShowCoverPic     int    `json:"show_cover_pic"`
	Content          string `json:"content"`
	ContentSourceURL string `json:"content_source_url"`
	URL              string `json:"url"`
	DownURL          string `json:"down_url"`
}
func (material *Material) AddNews(articles []*Article) (mediaID string, err error)

//获取
func (material *Material) GetNews(id string) ([]*Article, error)

// UpdateNews 更新永久图文素材
func (material *Material) UpdateNews(article *Article, mediaID string, index int64) (err error)
```

### 永久视频素材
```go
//resAddMaterial 永久性素材上传返回的结果
type resAddMaterial struct {
    ErrCode int64  `json:"errcode"`
    ErrMsg  string `json:"errmsg"`
    MediaID string `json:"media_id"`
    URL     string `json:"url"`
}

func (material *Material) AddVideo(filename, title, introduction string) (mediaID string, url string, err error)
```

### 其他类型永久素材
```go
//MediaTypeImage 媒体文件:图片
MediaTypeImage MediaType = "image"
//MediaTypeVoice 媒体文件:声音
MediaTypeVoice MediaType = "voice"
//MediaTypeThumb 媒体文件:缩略图
MediaTypeThumb MediaType = "thumb"
//AddMaterial 上传永久性素材（处理视频需要单独上传）
func (material *Material) AddMaterial(mediaType MediaType, filename string) (mediaID string, url string, err error)
```

### 删除永久素材
```go
func (material *Material) DeleteMaterial(mediaID string) error
```

### 批量获取永久素材
```go
//PermanentMaterialTypeImage 永久素材图片类型（image）
PermanentMaterialTypeImage PermanentMaterialType = "image"
//PermanentMaterialTypeVideo 永久素材视频类型（video）
PermanentMaterialTypeVideo PermanentMaterialType = "video"
//PermanentMaterialTypeVoice 永久素材语音类型 （voice）
PermanentMaterialTypeVoice PermanentMaterialType = "voice"
//PermanentMaterialTypeNews 永久素材图文类型（news）
PermanentMaterialTypeNews PermanentMaterialType = "news"
func (material *Material) BatchGetMaterial(permanentMaterialType PermanentMaterialType, offset, count int64) (list ArticleList, err error)
```

### 素材总数
```go
// ResMaterialCount 素材总数
type ResMaterialCount struct {
    ErrCode int64  `json:"errcode"`
    ErrMsg  string `json:"errmsg"`
    VoiceCount int64 `json:"voice_count"` // 语音总数量
    VideoCount int64 `json:"video_count"` // 视频总数量
    ImageCount int64 `json:"image_count"` // 图片总数量
    NewsCount  int64 `json:"news_count"`  // 图文总数量
}
func (material *Material) GetMaterialCount() (res ResMaterialCount, err error)
```

### 卡券
```go
func (officialAccount *OfficialAccount) GetCard() *card.Card
```

### 颜色
```go

type Colors struct {
	Name string `json:"name"`
	Value string `json:"value"`
}

// res []Colors
res, err := officail.GetCard().GetColors()
```


### 卡券开放类目查询接口
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

### 设置白名单
```go
func (card *Card) SetWhiteListByOpenid(openids []string) (err error)

func (card *Card) SetWhiteListByUsername(names []string) (err error)
```

### 创建卡券

所有创建卡券信息都调用该接口就行，传入不同的`card_type` 以及对应卡券所需要的字段信息`util.Params`

```go
// 示例创建一张会员卡
"github.com/fideism/golang-wechat/util"

attrs := util.Params{
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

### 查看卡券详情
```go
// 卡券ID  返回 util.Params
card, err := officail.GetCard().GetCard("xxxxx")
```

### 修改卡券信息
```go
// 卡券ID 卡券类型 卡券其他字段参考新增卡券接口
// 返回 是否提交审核，false为修改后不会重新提审，true为修改字段后重新提审，该卡券的状态变为审核中
func (card *Card) UpdateCard(cardID string, t Type, attrs util.Params) (check bool, err error)
```

### 删除卡券
```go
// 卡券ID 
func (card *Card) DeleteCard(cardID string) (err error)
```

### 批量查询卡券列表
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

### 设置会员卡开卡字段接口
```go
// 示例
attrs := util.Params{
    "required_form": map[string]interface{}{
        "common_field_id_list": []string{"USER_FORM_INFO_FLAG_NAME", "USER_FORM_INFO_FLAG_MOBILE"},
    },
    "optional_form": map[string]interface{}{
        "common_field_id_list": []string{"USER_FORM_INFO_FLAG_BIRTHDAY", "USER_FORM_INFO_FLAG_SEX"},
    },
}

// cardid 卡ID  map[string]interface{} 详细字段信息
func (card *Card) SetActivateUserForm(cardID string, attrs util.Params) (err error)
```

### 创建卡券二维码
```go
attrs := util.Params{
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
func (card *Card) CreateCardQrcode(attr util.Params) (res Qrcode, err error)
```

### 消息

客服接口，发送单个用户消息

```go
func (o *OfficialAccount) GetMessage() *message.Manager

//获取实例
message := officail.GetMessage()

// 消息格式
// 快捷构造消息类型实例
func CustomerTextMessage(toUser, text string) *CustomerMessage
func CustomerImgMessage(toUser, mediaID string) *CustomerMessage
func CustomerVoiceMessage(toUser, mediaID string) *CustomerMessage
func CustomerWxCardMessage(toUser, cardID string) *CustomerMessage

// 其他类型消息
message.CustomerMessage
```

- 文本消息
```go
text := message.NewCustomerTextMessage("openid", "hello")
err := officail.GetMessage().Send(text)

err = officail.GetMessage().Send(&message.CustomerMessage{
		ToUser:  openid,
		Msgtype: message.MsgTypeText,
		Text: &message.MediaText{
			Content: "word",
		},
	})
```

其他类型消息
```go
//CustomerMessage  客服消息
type CustomerMessage struct {
	ToUser          string                `json:"touser"`                    //接受者OpenID
	Msgtype         MsgType               `json:"msgtype"`                   //客服消息类型
	Text            *MediaText            `json:"text,omitempty"`            //可选
	Image           *MediaResource        `json:"image,omitempty"`           //可选
	Voice           *MediaResource        `json:"voice,omitempty"`           //可选
	Video           *MediaVideo           `json:"video,omitempty"`           //可选
	Music           *MediaMusic           `json:"music,omitempty"`           //可选
	News            *MediaNews            `json:"news,omitempty"`            //可选
	Mpnews          *MediaResource        `json:"mpnews,omitempty"`          //可选
	Wxcard          *MediaWxcard          `json:"wxcard,omitempty"`          //可选
	Msgmenu         *MediaMsgmenu         `json:"msgmenu,omitempty"`         //可选
	Miniprogrampage *MediaMiniprogrampage `json:"miniprogrampage,omitempty"` //可选
}

//MsgTypeText 表示文本消息
MsgTypeText MsgType = "text"
//MsgTypeImage 表示图片消息
MsgTypeImage = "image"
//MsgTypeVoice 表示语音消息
MsgTypeVoice = "voice"
//MsgTypeVideo 表示视频消息
MsgTypeVideo = "video"
//MsgTypeShortVideo 表示短视频消息[限接收]
MsgTypeShortVideo = "shortvideo"
//MsgTypeLocation 表示坐标消息[限接收]
MsgTypeLocation = "location"
//MsgTypeLink 表示链接消息[限接收]
MsgTypeLink = "link"
//MsgTypeMusic 表示音乐消息[限回复]
MsgTypeMusic = "music"
//MsgTypeNews 表示图文消息[限回复]
MsgTypeNews = "news"
//MsgTypeTransfer 表示消息消息转发到客服
MsgTypeTransfer = "transfer_customer_service"
//MsgTypeEvent 表示事件推送消息
MsgTypeEvent = "event"
//MsgTypeWxcard 卡券消息
MsgTypeWxcard = "wxcard"
```

### js-sdk配置
```go
func (officialAccount *OfficialAccount) GetJs() *js.Js
```

### 群发消息
```go
func (officialAccount *OfficialAccount) GetBroadcast() *broadcast.Broadcast

// 消息发送用户
type User struct {
	TagID  int64
	OpenID []string
}
//user 为nil，表示全员发送
//&User{TagID:2} 根据tag发送
//&User{OpenID:[]string("xxx","xxx")} 根据openid发送


//Result 群发返回结果
type Result struct {
    ErrCode int64  `json:"errcode"`
    ErrMsg  string `json:"errmsg"`
    MsgID     int64 `json:"msg_id"`
    MsgDataID int64 `json:"msg_data_id"`
}
```

### 发送群消息
```go
//组装数据发送
import github.com/fideism/golang-wechat/util
p := util.Params{
    "filter": map[string]interface{}{
        "is_to_all": false,
        "tag_id":    100,
    },
    "text": map[string]string{
        "content": "11111111111111111111111111111111",
    },
    "msgtype": "text",
}
// SendAll 根据标签进行群发
// https://api.weixin.qq.com/cgi-bin/message/mass/sendall
func (broadcast *Broadcast) SendAll(params util.Params) (res *Result, err error)
// SendByOpenID 根据标签进行群发
// https://api.weixin.qq.com/cgi-bin/message/mass/send
func (broadcast *Broadcast) SendByOpenID(params util.Params) (res *Result, err error)
```

### 删除群消息
```go
func (broadcast *Broadcast) Delete(msgID int64, articleIDx int64) error
```

### 预览群消息
```go
import github.com/fideism/golang-wechat/util

p := util.Params{
    "text": map[string]string{
        "content": "world",
    },
    "msgtype": "text",
}

err := broad.PreviewOpenid("openid", p)

func (broadcast *Broadcast) PreviewWxName(name string, params util.Params) (err error)
func (broadcast *Broadcast) PreviewOpenid(openid string, params util.Params) (err error)
```

### 群消息状态
```go
type GetResult struct {
    ErrCode int64  `json:"errcode"`
    ErrMsg  string `json:"errmsg"`
    MsgID     int64  `json:"msg_id"`
    MsgStatus string `json:"msg_status"`
}

func (broadcast *Broadcast) GetMass(msgID int64) (res *GetResult, err error)
```

### 群消息速度
```go
func (broadcast *Broadcast) Speed(speed, realspeed int64) (err error)
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

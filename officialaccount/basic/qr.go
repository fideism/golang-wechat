package basic

import (
	"encoding/json"
	"fmt"
	"reflect"
	"time"

	"github.com/fideism/golang-wechat/util"
)

const (
	qrCreateURL = "https://api.weixin.qq.com/cgi-bin/qrcode/create?access_token=%s"
	getQRImgURL = "https://mp.weixin.qq.com/cgi-bin/showqrcode?ticket=%s"
)

// QrActionName action_name 二维码类型
type QrActionName string

const (
	// QrActionScene QR_SCENE为临时的整型参数值
	QrActionScene QrActionName = "QR_SCENE"
	// QrActionStrScene QR_STR_SCENE为临时的字符串参数值
	QrActionStrScene QrActionName = "QR_STR_SCENE"
	// QrActionLimitScene QR_LIMIT_SCENE为永久的整型参数值
	QrActionLimitScene QrActionName = "QR_LIMIT_SCENE"
	// QrActionLimitStrScene QR_LIMIT_STR_SCENE为永久的字符串参数值
	QrActionLimitStrScene QrActionName = "QR_LIMIT_STR_SCENE"
)

// Request 临时二维码
type Request struct {
	ExpireSeconds int64        `json:"expire_seconds,omitempty"`
	ActionName    QrActionName `json:"action_name"`
	ActionInfo    struct {
		Scene struct {
			SceneStr string `json:"scene_str,omitempty"`
			SceneID  int    `json:"scene_id,omitempty"`
		} `json:"scene"`
	} `json:"action_info"`
}

// Ticket 二维码ticket
type Ticket struct {
	util.CommonError `json:",inline"`
	Ticket           string `json:"ticket"`
	ExpireSeconds    int64  `json:"expire_seconds"`
	URL              string `json:"url"`
}

// GetQRTicket 获取二维码 Ticket
func (basic *Basic) GetQRTicket(tq *Request) (t *Ticket, err error) {
	accessToken, err := basic.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(qrCreateURL, accessToken)
	response, err := util.PostJSON(uri, tq)
	if err != nil {
		err = fmt.Errorf("get qr ticket failed, %s", err)
		return
	}

	t = new(Ticket)
	err = json.Unmarshal(response, &t)
	if err != nil {
		return
	}

	return
}

// ShowQRCode 通过ticket换取二维码
func ShowQRCode(tk *Ticket) string {
	return fmt.Sprintf(getQRImgURL, tk.Ticket)
}

// NewTmpQrRequest 新建临时二维码请求实例
func NewTmpQrRequest(exp time.Duration, scene interface{}) *Request {
	tq := &Request{
		ExpireSeconds: int64(exp.Seconds()),
	}
	switch reflect.ValueOf(scene).Kind() {
	case reflect.String:
		tq.ActionName = QrActionStrScene
		tq.ActionInfo.Scene.SceneStr = scene.(string)
	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64:
		tq.ActionName = QrActionScene
		tq.ActionInfo.Scene.SceneID = scene.(int)
	}

	return tq
}

// NewLimitQrRequest 新建永久二维码请求实例
func NewLimitQrRequest(scene interface{}) *Request {
	tq := &Request{}
	switch reflect.ValueOf(scene).Kind() {
	case reflect.String:
		tq.ActionName = QrActionLimitStrScene
		tq.ActionInfo.Scene.SceneStr = scene.(string)
	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64:
		tq.ActionName = QrActionLimitScene
		tq.ActionInfo.Scene.SceneID = scene.(int)
	}

	return tq
}

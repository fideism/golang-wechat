package broadcast

import (
	"fmt"

	"github.com/fideism/golang-wechat/officialaccount/context"
	"github.com/fideism/golang-wechat/util"
)

const (
	sendURLByTag    = "https://api.weixin.qq.com/cgi-bin/message/mass/sendall"
	sendURLByOpenID = "https://api.weixin.qq.com/cgi-bin/message/mass/send"
	deleteSendURL   = "https://api.weixin.qq.com/cgi-bin/message/mass/delete"
	previewURL      = "https://api.weixin.qq.com/cgi-bin/message/mass/preview"
	getURL          = "https://api.weixin.qq.com/cgi-bin/message/mass/get"
	speedURL        = "https://api.weixin.qq.com/cgi-bin/message/mass/speed/get"
)

//Broadcast 群发消息
type Broadcast struct {
	*context.Context
}

//NewBroadcast new
func NewBroadcast(ctx *context.Context) *Broadcast {
	return &Broadcast{ctx}
}

//Result 群发返回结果
type Result struct {
	util.CommonError
	MsgID     int64 `json:"msg_id"`
	MsgDataID int64 `json:"msg_data_id"`
}

// SendAll 根据标签进行群发
func (broadcast *Broadcast) SendAll(params util.Params) (res *Result, err error) {
	var token string
	token, err = broadcast.GetAccessToken()
	if err != nil {
		return
	}

	url := fmt.Sprintf("%s?access_token=%s", sendURLByTag, token)

	var response []byte
	response, err = util.PostJSON(url, params)
	if err != nil {
		return
	}

	res = new(Result)

	err = util.DecodeWithError(response, res, "SendAll")

	return
}

// SendByOpenID 根据标签进行群发
func (broadcast *Broadcast) SendByOpenID(params util.Params) (res *Result, err error) {
	var token string
	token, err = broadcast.GetAccessToken()
	if err != nil {
		return
	}

	url := fmt.Sprintf("%s?access_token=%s", sendURLByOpenID, token)

	var response []byte
	response, err = util.PostJSON(url, params)
	if err != nil {
		return
	}

	res = new(Result)

	err = util.DecodeWithError(response, res, "SendByOpenID")

	return
}

//Delete 删除群发消息
func (broadcast *Broadcast) Delete(msgID int64, articleIDx int64) error {
	ak, err := broadcast.GetAccessToken()
	if err != nil {
		return err
	}
	req := map[string]interface{}{
		"msg_id":      msgID,
		"article_idx": articleIDx,
	}
	url := fmt.Sprintf("%s?access_token=%s", deleteSendURL, ak)
	data, err := util.PostJSON(url, req)
	if err != nil {
		return err
	}
	return util.DecodeWithCommonError(data, "Delete")
}

// PreviewWxName wxname微信号的预览
func (broadcast *Broadcast) PreviewWxName(name string, params util.Params) (err error) {
	var token string
	token, err = broadcast.GetAccessToken()
	if err != nil {
		return
	}
	url := fmt.Sprintf("%s?access_token=%s", previewURL, token)
	params.Set("towxname", name)
	data, err := util.PostJSON(url, params)

	err = util.DecodeWithCommonError(data, "PreviewOpenid")

	return
}

//PreviewOpenid openid微信号的预览
func (broadcast *Broadcast) PreviewOpenid(openid string, params util.Params) (err error) {
	var token string
	token, err = broadcast.GetAccessToken()
	if err != nil {
		return
	}
	url := fmt.Sprintf("%s?access_token=%s", previewURL, token)
	params.Set("touser", openid)
	data, err := util.PostJSON(url, params)

	err = util.DecodeWithCommonError(data, "PreviewOpenid")

	return
}

// GetResult 状态查询返回
type GetResult struct {
	util.CommonError
	MsgID     int64  `json:"msg_id"`
	MsgStatus string `json:"msg_status"`
}

// GetMass 查询状态
func (broadcast *Broadcast) GetMass(msgID int64) (res *GetResult, err error) {
	var token string
	token, err = broadcast.GetAccessToken()
	if err != nil {
		return
	}

	url := fmt.Sprintf("%s?access_token=%s", getURL, token)

	var response []byte
	response, err = util.PostJSON(url, map[string]int64{
		"msg_id": msgID,
	})
	if err != nil {
		return
	}

	res = new(GetResult)

	err = util.DecodeWithError(response, res, "GetMass")

	return
}

// Speed 速度控制
func (broadcast *Broadcast) Speed(speed, realspeed int64) (err error) {
	var token string
	token, err = broadcast.GetAccessToken()
	if err != nil {
		return
	}

	url := fmt.Sprintf("%s?access_token=%s", speedURL, token)

	var response []byte
	response, err = util.PostJSON(url, map[string]int64{
		"speed":     speed,
		"realspeed": realspeed,
	})
	if err != nil {
		return
	}

	err = util.DecodeWithCommonError(response, "Speed")
	return
}

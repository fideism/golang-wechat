package auth

import (
	"fmt"

	"github.com/fideism/golang-wechat/miniprogram/context"
	"github.com/fideism/golang-wechat/util"
)

const (
	code2SessionURL   = "https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code"
	getPaidUnionIDURL = "https://api.weixin.qq.com/wxa/getpaidunionid?access_token=%s&openid=%s"
)

//Auth 登录/用户信息
type Auth struct {
	*context.Context
}

//NewAuth new auth
func NewAuth(ctx *context.Context) *Auth {
	return &Auth{ctx}
}

// ResCode2Session 登录凭证校验的返回结果
type ResCode2Session struct {
	util.CommonError

	OpenID     string `json:"openid"`      // 用户唯一标识
	SessionKey string `json:"session_key"` // 会话密钥
	UnionID    string `json:"unionid"`     // 用户在开放平台的唯一标识符，在满足UnionID下发条件的情况下会返回
}

//Code2Session 登录凭证校验。
func (auth *Auth) Code2Session(jsCode string) (result ResCode2Session, err error) {
	urlStr := fmt.Sprintf(code2SessionURL, auth.AppID, auth.AppSecret, jsCode)
	var response []byte
	response, err = util.HTTPGet(urlStr)
	if err != nil {
		return
	}

	err = util.DecodeWithError(response, &result, "Code2Session")
	return
}

// ResGetPaidUnionID 支付后获取用户unionid的返回结果
type ResGetPaidUnionID struct {
	util.CommonError

	UnionID string `json:"unionid"` // 用户在开放平台的唯一标识符，在满足UnionID下发条件的情况下会返回
}

//GetPaidUnionID 用户支付完成后，获取该用户的 UnionId，无需用户授权
func (auth *Auth) GetPaidUnionID(p util.Params) (result ResGetPaidUnionID, err error) {
	token, err := auth.Context.GetAccessToken()
	if err != nil {
		return
	}

	urlStr := paidUnionIDURL(p, token)

	var response []byte
	response, err = util.HTTPGet(urlStr)
	if err != nil {
		return
	}

	err = util.DecodeWithError(response, &result, "Code2Session")
	return
}

func paidUnionIDURL(p util.Params, token string) string {
	urlStr := fmt.Sprintf(getPaidUnionIDURL, token, p.GetString(`openid`))
	if p.Exists(`transaction_id`) {
		urlStr += `&transaction_id=` + p.GetString(`transaction_id`)
	}

	if p.Exists(`out_trade_no`) && p.Exists(`mch_id`) {
		urlStr += `&out_trade_no=` + p.GetString(`out_trade_no`) +
			`&mch_id=` + p.GetString(`mch_id`)
	}

	return urlStr
}

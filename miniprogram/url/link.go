package url

import (
	"fmt"

	"github.com/fideism/golang-wechat/miniprogram/context"
	"github.com/fideism/golang-wechat/util"
)

// 文档地址 https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/url-link/urllink.generate.html

//generateURLLink url_link.generate
const generateURLLink = `https://api.weixin.qq.com/wxa/generate_urllink?access_token=%s`

// URL struct
type URL struct {
	*context.Context
}

// NewURL 实例
func NewURL(ctx *context.Context) *URL {
	return &URL{ctx}
}

// GenerateRequest 请求
type GenerateRequest struct {
	EnvVersion     string `json:"env_version"`
	Path           string `json:"path"`
	Query          string `json:"query"`
	IsExpire       bool   `json:"is_expire"`
	ExpireType     int8   `json:"expire_type"`
	ExpireInterval int8   `json:"expire_interval"`
}

// GenerateResponse 响应
type GenerateResponse struct {
	util.CommonError
	URLLink string `json:"url_link"`
}

// Generate 生成URL Link
func (u *URL) Generate(req GenerateRequest) (*GenerateResponse, error) {
	accessToken, tokenErr := u.GetAccessToken()
	fmt.Println(accessToken)
	if tokenErr != nil {
		return nil, tokenErr
	}
	uri := fmt.Sprintf(generateURLLink, accessToken)
	response, err := util.PostJSON(uri, req)
	if err != nil {
		return nil, err
	}

	var result GenerateResponse
	err = util.DecodeWithError(response, &result, "Generate")

	return &result, nil
}

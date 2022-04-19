package card

import (
	"fmt"
	"strings"

	"github.com/fideism/golang-wechat/officialaccount/context"
	"github.com/fideism/golang-wechat/util"
)

// CouponType 券类型
type CouponType string

const (
	// Groupon 团购券类型
	Groupon CouponType = "GROUPON"
	// Cash 代金券类型
	Cash CouponType = "CASH"
	// Discount 折扣券类型
	Discount CouponType = "DISCOUNT"
	// Gift 兑换券类型
	Gift CouponType = "GIFT"
	// GeneralCoupon 通用券。
	GeneralCoupon CouponType = "GENERAL_COUPON"
	// MemberCard 会员卡
	MemberCard CouponType = "MEMBER_CARD"
	// GeneralCard 礼品卡
	GeneralCard CouponType = "GENERAL_CARD"
)

// Status 卡券状态
type Status string

const (
	// StatusNotVerify 待审核
	StatusNotVerify Status = "CARD_STATUS_NOT_VERIFY"
	// StatusVerifyFail 审核失败
	StatusVerifyFail Status = "CARD_STATUS_VERIFY_FAIL"
	// StatusVerifyOk 审核成功
	StatusVerifyOk Status = "CARD_STATUS_VERIFY_OK"
	// StatusDelete 卡券被商户删除
	StatusDelete Status = "CARD_STATUS_DELETE"
	// StatusDispatch 在公众平台投放过的卡券
	StatusDispatch Status = "CARD_STATUS_DISPATCH"
)

// Card 卡券
type Card struct {
	*context.Context
}

//NewCard init
func NewCard(context *context.Context) *Card {
	card := new(Card)
	card.Context = context

	return card
}

// CreateCard 创建卡券
func (card *Card) CreateCard(t CouponType, attrs util.Params) (cardID string, err error) {
	var token string
	token, err = card.GetAccessToken()
	if err != nil {
		return
	}
	url := fmt.Sprintf(createCardURL, token)

	params := map[string]interface{}{
		"card": map[string]interface{}{
			"card_type":                t,
			strings.ToLower(string(t)): attrs,
		},
	}

	var response []byte
	response, err = util.PostJSON(url, params)
	if err != nil {
		return
	}

	var res struct {
		util.CommonError
		CardID string `json:"card_id"`
	}

	err = util.DecodeWithError(response, &res, "CreateCard")
	if err != nil {
		return
	}

	cardID = res.CardID

	return
}

// UpdateCard 修改卡券信息
func (card *Card) UpdateCard(cardID string, t CouponType, attrs util.Params) (check bool, err error) {
	var token string
	token, err = card.GetAccessToken()
	if err != nil {
		return
	}
	url := fmt.Sprintf(updateURL, token)

	params := map[string]interface{}{
		"card_id":                  cardID,
		strings.ToLower(string(t)): attrs,
	}

	var response []byte
	response, err = util.PostJSON(url, params)
	if err != nil {
		return
	}

	var res struct {
		util.CommonError
		SendCheck bool `json:"send_check"`
	}

	err = util.DecodeWithError(response, &res, "UpdateCard")
	if err != nil {
		return
	}

	check = res.SendCheck

	return
}

// GetCard 获取卡券详情
func (card *Card) GetCard(cardID string) (res util.Params, err error) {
	var token string
	token, err = card.GetAccessToken()
	if err != nil {
		return
	}
	url := fmt.Sprintf(getCardURL, token)

	var response []byte
	response, err = util.PostJSON(url, map[string]interface{}{
		"card_id": cardID,
	})
	if err != nil {
		return
	}

	var cardInfo struct {
		util.CommonError
		Card util.Params `json:"card"`
	}

	err = util.DecodeWithError(response, &cardInfo, "GetCard")
	if err != nil {
		return
	}

	res = cardInfo.Card

	return
}

// BatchGetRequest 批量查询卡券列表 请求参数
type BatchGetRequest struct {
	Offset     int    `json:"offset"`
	Count      int    `json:"count"`
	StatusList Status `json:"status_list"`
}

// BatchCardList 批量查询卡券列表 返回结果
type BatchCardList struct {
	CardIDList []string `json:"card_id_list"`
	TotalNum   int      `json:"total_num"`
}

// BatchGet 批量查询卡券列表
func (card *Card) BatchGet(req BatchGetRequest) (res BatchCardList, err error) {
	var token string
	token, err = card.GetAccessToken()
	if err != nil {
		return
	}
	url := fmt.Sprintf(batchGetURL, token)

	var response []byte
	response, err = util.PostJSON(url, req)
	if err != nil {
		return
	}

	var items struct {
		util.CommonError
		BatchCardList
	}

	err = util.DecodeWithError(response, &items, "CreateCard")
	if err != nil {
		return
	}

	res = BatchCardList{
		CardIDList: items.CardIDList,
		TotalNum:   items.TotalNum,
	}

	return
}

// DeleteCard 删除卡券
func (card *Card) DeleteCard(cardID string) (err error) {
	var token string
	token, err = card.GetAccessToken()
	if err != nil {
		return
	}
	url := fmt.Sprintf(deleteURL, token)

	var response []byte
	response, err = util.PostJSON(url, map[string]interface{}{
		"card_id": cardID,
	})
	if err != nil {
		return
	}

	err = util.DecodeWithCommonError(response, "DeleteCard")

	return
}

// Qrcode 卡券二维码
type Qrcode struct {
	Ticket        string `json:"ticket"`
	ExpireSeconds int64  `json:"expire_seconds"`
	URL           string `json:"url"`
	ShowQrcodeURL string `json:"show_qrcode_url"`
}

// CreateCardQrcode 创建卡券二维码
func (card *Card) CreateCardQrcode(attr util.Params) (res Qrcode, err error) {
	var token string
	token, err = card.GetAccessToken()
	if err != nil {
		return
	}
	url := fmt.Sprintf(createQrcodeURL, token)

	var response []byte
	response, err = util.PostJSON(url, attr)
	if err != nil {
		return
	}

	var qrcode struct {
		util.CommonError
		Qrcode
	}
	err = util.DecodeWithError(response, &qrcode, "GetColors")
	if err != nil {
		return
	}

	res = Qrcode{
		Ticket:        qrcode.Ticket,
		ExpireSeconds: qrcode.ExpireSeconds,
		URL:           qrcode.URL,
		ShowQrcodeURL: qrcode.ShowQrcodeURL,
	}

	return
}

// GetHTML 图文消息群发卡券
func (card *Card) GetHTML(cardID string) (content string, err error) {
	var token string
	token, err = card.GetAccessToken()
	if err != nil {
		return
	}
	url := fmt.Sprintf(getHTMLURL, token)

	var response []byte
	response, err = util.PostJSON(url, map[string]interface{}{
		"card_id": cardID,
	})
	if err != nil {
		return
	}

	var res struct {
		util.CommonError
		Content string `json:"content"`
	}

	err = util.DecodeWithError(response, &res, "GetHtml")
	if err != nil {
		return
	}

	content = res.Content

	return
}

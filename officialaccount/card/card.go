package card

import (
	"fmt"
	"github.com/fideism/golang-wechat/officialaccount/context"
	"github.com/fideism/golang-wechat/util"
	"strings"
)

// Type 券类型
type Type string

const (
	// Groupon 团购券类型
	Groupon Type = "GROUPON"
	// Cash 代金券类型
	Cash Type = "CASH"
	// Discount 折扣券类型
	Discount Type = "DISCOUNT"
	// Gift 兑换券类型
	Gift Type = "GIFT"
	// GeneralCoupon 通用券。
	GeneralCoupon Type = "GENERAL_COUPON"
	// MemberCard 会员卡
	MemberCard Type = "MEMBER_CARD"
	// GeneralCard 礼品卡
	GeneralCard Type = "GENERAL_CARD"
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

// Colors 卡券颜色
type Colors struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// GetColors 获取卡券颜色
func (card *Card) GetColors() (colors []Colors, err error) {
	var token string
	token, err = card.GetAccessToken()
	if err != nil {
		return
	}
	uri := fmt.Sprintf("%s?access_token=%s", colorURL, token)

	var response []byte
	response, err = util.HTTPGet(uri)
	if err != nil {
		return
	}

	var res struct {
		util.CommonError
		Colors []Colors `json:"colors"`
	}
	err = util.DecodeWithError(response, &res, "GetColors")
	if err != nil {
		return
	}

	colors = res.Colors

	return
}

// Category 卡券类目信息
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

// GetApplyProtocol 卡券开放类目查询接口
func (card *Card) GetApplyProtocol() (c []Category, err error) {
	var token string
	token, err = card.GetAccessToken()
	if err != nil {
		return
	}
	uri := fmt.Sprintf("%s?access_token=%s", applyProtocolURL, token)

	var response []byte
	response, err = util.HTTPGet(uri)
	if err != nil {
		return
	}

	var res struct {
		util.CommonError
		Category []Category `json:"category"`
	}

	err = util.DecodeWithError(response, &res, "GetColors")
	if err != nil {
		return
	}

	c = res.Category

	return
}

// CreateCard 创建卡券
func (card *Card) CreateCard(t Type, attrs interface{}) (cardID string, err error) {
	var token string
	token, err = card.GetAccessToken()
	if err != nil {
		return
	}
	uri := fmt.Sprintf("%s?access_token=%s", createCardURL, token)

	params := map[string]interface{}{
		"card": map[string]interface{}{
			"card_type":                t,
			strings.ToLower(string(t)): attrs,
		},
	}

	var response []byte
	response, err = util.PostJSON(uri, params)
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
func (card *Card) UpdateCard(cardID string, t Type, attrs interface{}) (check bool, err error) {
	var token string
	token, err = card.GetAccessToken()
	if err != nil {
		return
	}
	uri := fmt.Sprintf("%s?access_token=%s", updateURL, token)

	params := map[string]interface{}{
		"card_id":                  cardID,
		strings.ToLower(string(t)): attrs,
	}

	var response []byte
	response, err = util.PostJSON(uri, params)
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
func (card *Card) GetCard(cardID string) (res map[string]interface{}, err error) {
	var token string
	token, err = card.GetAccessToken()
	if err != nil {
		return
	}
	uri := fmt.Sprintf("%s?access_token=%s", getCardURL, token)

	var response []byte
	response, err = util.PostJSON(uri, map[string]interface{}{
		"card_id": cardID,
	})
	if err != nil {
		return
	}

	var cardInfo struct {
		util.CommonError
		Card map[string]interface{} `json:"card"`
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
	uri := fmt.Sprintf("%s?access_token=%s", batchGetURL, token)

	var response []byte
	response, err = util.PostJSON(uri, req)
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
	uri := fmt.Sprintf("%s?access_token=%s", deleteURL, token)

	var response []byte
	response, err = util.PostJSON(uri, map[string]interface{}{
		"card_id": cardID,
	})
	if err != nil {
		return
	}

	var res struct {
		util.CommonError
	}

	err = util.DecodeWithError(response, &res, "DeleteCard")
	if err != nil {
		return
	}

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
func (card *Card) CreateCardQrcode(attr map[string]interface{}) (res Qrcode, err error) {
	var token string
	token, err = card.GetAccessToken()
	if err != nil {
		return
	}
	uri := fmt.Sprintf("%s?access_token=%s", createQrcodeURL, token)

	var response []byte
	response, err = util.PostJSON(uri, attr)
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

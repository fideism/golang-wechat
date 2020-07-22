package card

import (
	"fmt"
	"github.com/fideism/golang-wechat/util"
)

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
	url := fmt.Sprintf(colorURL, token)

	var response []byte
	response, err = util.HTTPGet(url)
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
	url := fmt.Sprintf(applyProtocolURL, token)

	var response []byte
	response, err = util.HTTPGet(url)
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

// SetWhiteListByOpenid 通过openid设置白名单
func (card *Card) SetWhiteListByOpenid(openids []string) (err error) {
	var token string
	token, err = card.GetAccessToken()
	if err != nil {
		return
	}
	url := fmt.Sprintf(whitelistURL, token)

	var response []byte
	response, err = util.PostJSON(url, map[string]interface{}{
		"openid": openids,
	})
	if err != nil {
		return
	}

	err = util.DecodeWithCommonError(response, "SetWhiteListByOpenid")

	return
}

// SetWhiteListByUsername 通过username设置白名单
func (card *Card) SetWhiteListByUsername(names []string) (err error) {
	var token string
	token, err = card.GetAccessToken()
	if err != nil {
		return
	}
	url := fmt.Sprintf(whitelistURL, token)

	var response []byte
	response, err = util.PostJSON(url, map[string]interface{}{
		"username": names,
	})
	if err != nil {
		return
	}

	err = util.DecodeWithCommonError(response, "SetWhiteListByUsername")

	return
}

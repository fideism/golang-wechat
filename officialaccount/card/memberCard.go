package card

import (
	"fmt"
	"github.com/fideism/golang-wechat/util"
)

// SetActivateUserForm 设置开卡字段接口
func (card *Card) SetActivateUserForm(cardID string, attrs util.Params) (err error) {
	var token string
	token, err = card.GetAccessToken()
	if err != nil {
		return
	}
	uri := fmt.Sprintf("%s?access_token=%s", memberSetActivateFieldURL, token)

	attrs["card_id"] = cardID

	var response []byte
	response, err = util.PostJSON(uri, attrs)
	if err != nil {
		return
	}
	var res struct {
		util.CommonError
	}

	err = util.DecodeWithError(response, &res, "SetActivateUserForm")

	return
}

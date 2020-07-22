package user

import (
	"fmt"
	"net/url"

	"github.com/fideism/golang-wechat/officialaccount/context"
	"github.com/fideism/golang-wechat/util"
)

const (
	userInfoURL     = "https://api.weixin.qq.com/cgi-bin/user/info?access_token=%s&openid=%s&lang=zh_CN"
	updateRemarkURL = "https://api.weixin.qq.com/cgi-bin/user/info/updateremark?access_token=%s"
	userListURL     = "https://api.weixin.qq.com/cgi-bin/user/get"
	createTagURL    = "https://api.weixin.qq.com/cgi-bin/tags/create?access_token=%s"
	tagListURL      = "https://api.weixin.qq.com/cgi-bin/tags/get?access_token=%s"
	updateTagURL    = "https://api.weixin.qq.com/cgi-bin/tags/update?access_token=%s"
	deleteTagURL    = "https://api.weixin.qq.com/cgi-bin/tags/delete?access_token=%s"
	tagUserListURL  = "https://api.weixin.qq.com/cgi-bin/user/tag/get?access_token=%s"
)

//User 用户管理
type User struct {
	*context.Context
}

//NewUser 实例化
func NewUser(context *context.Context) *User {
	user := new(User)
	user.Context = context
	return user
}

//Info 用户基本信息
type Info struct {
	util.CommonError

	Subscribe      int32   `json:"subscribe"`
	OpenID         string  `json:"openid"`
	Nickname       string  `json:"nickname"`
	Sex            int32   `json:"sex"`
	City           string  `json:"city"`
	Country        string  `json:"country"`
	Province       string  `json:"province"`
	Language       string  `json:"language"`
	Headimgurl     string  `json:"headimgurl"`
	SubscribeTime  int32   `json:"subscribe_time"`
	UnionID        string  `json:"unionid"`
	Remark         string  `json:"remark"`
	GroupID        int32   `json:"groupid"`
	TagIDList      []int32 `json:"tagid_list"`
	SubscribeScene string  `json:"subscribe_scene"`
	QrScene        int     `json:"qr_scene"`
	QrSceneStr     string  `json:"qr_scene_str"`
}

// OpenidList 用户列表
type OpenidList struct {
	Total int `json:"total"`
	Count int `json:"count"`
	Data  struct {
		OpenIDs []string `json:"openid"`
	} `json:"data"`
	NextOpenID string `json:"next_openid"`
}

//GetUserInfo 获取用户基本信息
func (user *User) GetUserInfo(openID string) (userInfo *Info, err error) {
	var accessToken string
	accessToken, err = user.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(userInfoURL, accessToken, openID)
	var response []byte
	response, err = util.HTTPGet(uri)
	if err != nil {
		return
	}
	userInfo = new(Info)
	err = util.DecodeWithError(response, userInfo, "GetUserInfo")
	if err != nil {
		return
	}
	if userInfo.ErrCode != 0 {
		err = fmt.Errorf("GetUserInfo Error , errcode=%d , errmsg=%s", userInfo.ErrCode, userInfo.ErrMsg)
		return
	}
	return
}

// UpdateRemark 设置用户备注名
func (user *User) UpdateRemark(openID, remark string) (err error) {
	var accessToken string
	accessToken, err = user.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(updateRemarkURL, accessToken)
	var response []byte
	response, err = util.PostJSON(uri, map[string]string{"openid": openID, "remark": remark})
	if err != nil {
		return
	}

	return util.DecodeWithCommonError(response, "UpdateRemark")
}

// ListUserOpenIDs 返回用户列表
func (user *User) ListUserOpenIDs(nextOpenid ...string) (*OpenidList, error) {
	accessToken, err := user.GetAccessToken()
	if err != nil {
		return nil, err
	}

	uri, _ := url.Parse(userListURL)
	q := uri.Query()
	q.Set("access_token", accessToken)
	if len(nextOpenid) > 0 && nextOpenid[0] != "" {
		q.Set("next_openid", nextOpenid[0])
	}
	uri.RawQuery = q.Encode()

	response, err := util.HTTPGet(uri.String())
	if err != nil {
		return nil, err
	}

	userlist := new(OpenidList)
	err = util.DecodeWithCustomerStruct(response, userlist, "ListUserOpenIDs")
	if err != nil {
		return nil, err
	}

	return userlist, nil
}

// ListAllUserOpenIDs 返回所有用户OpenID列表
func (user *User) ListAllUserOpenIDs() ([]string, error) {
	nextOpenid := ""
	var openids []string
	count := 0
	for {
		ul, err := user.ListUserOpenIDs(nextOpenid)
		if err != nil {
			return nil, err
		}
		openids = append(openids, ul.Data.OpenIDs...)
		count += ul.Count
		if ul.Total > count {
			nextOpenid = ul.NextOpenID
		} else {
			return openids, nil
		}
	}
}

// Tag 标签
type Tag struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Count int64  `json:"count"`
}

// CreateTag 创建标签
func (user *User) CreateTag(name string) (tag Tag, err error) {
	var accessToken string
	accessToken, err = user.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(createTagURL, accessToken)
	var response []byte
	response, err = util.PostJSON(uri, map[string]interface{}{
		"tag": map[string]string{
			"name": name,
		},
	})

	if err != nil {
		return
	}

	var res struct {
		util.CommonError
		Tag Tag `json:"tag"`
	}

	err = util.DecodeWithError(response, &res, "CreateTag")
	if err != nil {
		return
	}

	tag = res.Tag

	return
}

// TagList 标签列表
func (user *User) TagList() (tags []Tag, err error) {
	var accessToken string
	accessToken, err = user.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(tagListURL, accessToken)
	var response []byte
	response, err = util.HTTPGet(uri)

	if err != nil {
		return
	}

	var res struct {
		util.CommonError
		Tags []Tag `json:"tags"`
	}

	err = util.DecodeWithError(response, &res, "TagList")
	if err != nil {
		return
	}

	tags = res.Tags

	return
}

// UpdateTag 修改标签
func (user *User) UpdateTag(tagID int, name string) (err error) {
	var accessToken string
	accessToken, err = user.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(updateTagURL, accessToken)
	var response []byte
	response, err = util.PostJSON(uri, map[string]interface{}{
		"tag": map[string]interface{}{
			"id":   tagID,
			"name": name,
		},
	})

	if err != nil {
		return
	}

	err = util.DecodeWithCommonError(response, "UpdateTag")
	if err != nil {
		return
	}

	return
}

// DeleteTag 删除标签
func (user *User) DeleteTag(tagID int) (err error) {
	var accessToken string
	accessToken, err = user.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(deleteTagURL, accessToken)
	var response []byte
	response, err = util.PostJSON(uri, map[string]interface{}{
		"tag": map[string]int{
			"id": tagID,
		},
	})

	if err != nil {
		return
	}

	err = util.DecodeWithCommonError(response, "DeleteTag")
	if err != nil {
		return
	}

	return
}

// TagUser 标签用户
type TagUser struct {
	Count int `json:"count"`
	Data  struct {
		Openid []string `json:"openid"`
	} `json:"data"`
	NextOpenID string `json:"next_openid"`
}

// TagUserList 获取标签下粉丝列表
func (user *User) TagUserList(tagID int, openid string) (res TagUser, err error) {
	var accessToken string
	accessToken, err = user.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(tagUserListURL, accessToken)
	var response []byte
	response, err = util.PostJSON(uri, map[string]interface{}{
		"tagid":       tagID,
		"next_openid": openid,
	})

	if err != nil {
		return
	}

	err = util.DecodeWithCustomerStruct(response, &res, "TagUserList")
	if err != nil {
		return
	}

	return
}

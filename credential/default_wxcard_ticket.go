package credential

import (
	"fmt"
	"sync"
	"time"

	"github.com/fideism/golang-wechat/cache"
	"github.com/fideism/golang-wechat/util"
)

//获取ticket的url
const getWxCardTicketURL = "https://api.weixin.qq.com/cgi-bin/ticket/getticket?access_token=%s&type=wx_card"

//DefaultWxCardTicket 默认获取wx_card ticket方法
type DefaultWxCardTicket struct {
	appID          string
	cacheKeyPrefix string
	cache          cache.Cache
	//jsAPITicket 读写锁 同一个AppID一个
	wxCardAPITicketLock *sync.Mutex
}

//NewDefaultWxCardTicket new
func NewDefaultWxCardTicket(appID string, cacheKeyPrefix string, cache cache.Cache) WxCardTicketHandle {
	return &DefaultWxCardTicket{
		appID:               appID,
		cache:               cache,
		cacheKeyPrefix:      cacheKeyPrefix,
		wxCardAPITicketLock: new(sync.Mutex),
	}
}

// ResWxCardTicket 请求wx_card_tikcet返回结果
type ResWxCardTicket struct {
	util.CommonError

	Ticket    string `json:"ticket"`
	ExpiresIn int64  `json:"expires_in"`
}

//GetTicket 获取jsapi_ticket
func (wxCard *DefaultWxCardTicket) GetTicket(accessToken string) (ticketStr string, err error) {
	wxCard.wxCardAPITicketLock.Lock()
	defer wxCard.wxCardAPITicketLock.Unlock()

	//先从cache中取
	jsAPITicketCacheKey := fmt.Sprintf("%s_jsapi_ticket_%s", wxCard.cacheKeyPrefix, wxCard.appID)
	val := wxCard.cache.Get(jsAPITicketCacheKey)
	if val != nil {
		ticketStr = val.(string)
		return
	}
	var ticket ResWxCardTicket
	ticket, err = GetWxCardTicketFromServer(accessToken)
	if err != nil {
		return
	}
	expires := ticket.ExpiresIn - 1500
	err = wxCard.cache.Set(jsAPITicketCacheKey, ticket.Ticket, time.Duration(expires)*time.Second)
	ticketStr = ticket.Ticket
	return
}

//GetWxCardTicketFromServer 从服务器中获取ticket
func GetWxCardTicketFromServer(accessToken string) (ticket ResWxCardTicket, err error) {
	var response []byte
	url := fmt.Sprintf(getWxCardTicketURL, accessToken)
	response, err = util.HTTPGet(url)
	if err != nil {
		return
	}

	err = util.DecodeWithError(response, &ticket, `GetWxCardTicketFromServer`)
	return
}

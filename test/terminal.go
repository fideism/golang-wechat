package main

import (
	"fmt"
	"github.com/fideism/golang-wechat/officialaccount/basic"

	"github.com/fideism/golang-wechat/cache"
	"github.com/fideism/golang-wechat/officialaccount"
	offConfig "github.com/fideism/golang-wechat/officialaccount/config"
)

var officail *officialaccount.OfficialAccount

func main() {
	//设置全局cache，也可以单独为每个操作实例设置
	redis := &cache.RedisOpts{
		Host:     "127.0.0.1:6379",
		Database: 1,
	}

	config := &offConfig.Config{
		AppID:          "wxa3936c080457a87f",
		AppSecret:      "0359fc119d448a86ebf4fd280f157507",
		Token:          "4ea9c8a0657f922a0824d44167ba3052",
		EncodingAESKey: "Kt5Gk0PJTvx7TddPnUIQtSYQiX8elSB9V9ktvcjCeRF",
		Cache:          cache.NewRedis(redis),
	}

	officail = officialaccount.NewOfficialAccount(config)

	token, err := officail.GetAccessToken()
	if err != nil {
		panic(err)
	}

	fmt.Println(token)

	//basicOfficail()

	//qrOfficial()

	userOfficial()
}

func userOfficial() {
	users := officail.GetUser()

	items, _ := users.ListUserOpenIDs("oyJUpv0bKZkR02M34QkU1UK2_IjA")
	fmt.Println(items)

	user, _ := users.GetUserInfo("oyJUpv0bKZkR02M34QkU1UK2_IjA")

	fmt.Println(user)

}

func basicOfficail() {
	basic := officail.GetBasic()

	fmt.Println(basic.GetAPIDomainIP())

	fmt.Println(basic.GetCallbackIP())
}

func qrOfficial() {
	req := &basic.Request{
		ExpireSeconds: 0,
		ActionName:    basic.QrActionScene,
	}
	req.ActionInfo.Scene.SceneStr = "scene"
	req.ActionInfo.Scene.SceneID = 123

	//req := basic.NewTmpQrRequest(300, "test")

	fmt.Println(req)

	b := officail.GetBasic()

	ticket, err := b.GetQRTicket(req)
	if nil != err {
		panic(err)
	}

	fmt.Println(ticket)

	url := basic.ShowQRCode(ticket)

	fmt.Println(url)
}

# 微信公众号

[官方文档](https://developers.weixin.qq.com/doc/offiaccount/Getting_Started/Overview.html)

## 快速入门

```go
import (
	"fmt"
	"github.com/fideism/golang-wechat/cache"
	"github.com/fideism/golang-wechat/officialaccount"
	offConfig "github.com/fideism/golang-wechat/officialaccount/config"
)

func main() {
	//设置全局cache，也可以单独为每个操作实例设置
	redis := &cache.RedisOpts{
		Host:     "127.0.0.1:6379",
		Database: 1,
	}

	config := &offConfig.Config{
		AppID:          "xxxx",
		AppSecret:      "xxxx",
		Token:          "xxxx",
		EncodingAESKey: "xxx",
		Cache:          cache.NewRedis(redis),
	}

	officail := officialaccount.NewOfficialAccount(config)

	token, err := officail.GetAccessToken()
	if err != nil {
		panic(err)
	}

	fmt.Println(token)
}
```
# WeChat SDK for Go

- [缓存](#缓存)
- [公众号](#公众号)
  - [Token](#Token) 

## 缓存

暂时只支持`redis`缓存配置

```go
import (
	"github.com/fideism/golang-wechat/cache"
)

func main() {
	//设置全局cache，也可以单独为每个操作实例设置
	redis := &cache.RedisOpts{
		Host:        "127.0.0.1:6379",
        Password:    "111111",
        Database:    1,
        MaxIdle:     5, //最大等待连接中的数量
        MaxActive:   3, //最大连接数据库连接数
        IdleTimeout: 1, //客户端的idle
	}
    
    cache := cache.NewRedis(redis)
}

```

## 公众号

### Token

### Based On
[silenceper/wechat](https://github.com/silenceper/wechat) 
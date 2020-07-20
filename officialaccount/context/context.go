package context

import (
	"github.com/fideism/golang-wechat/credential"
	"github.com/fideism/golang-wechat/officialaccount/config"
)

// Context struct
type Context struct {
	*config.Config
	credential.AccessTokenHandle
}

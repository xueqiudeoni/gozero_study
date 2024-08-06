package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	JwtAuth struct {
		AccessSecret string
	}
	WxMiniConf struct {
		Appid     string `json:AppId`
		Appsecret string `json:AppSecret`
	}
	UsercenterRpcConf zrpc.RpcClientConf
}

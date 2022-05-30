package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	Auth struct {
		AccessSecret string
		AccessExpire int64
	}
	LocalStore StoreServer
	UserRpc    zrpc.RpcClientConf
	VideoRpc   zrpc.RpcClientConf
}

type StoreServer struct {
	Path string
	Addr string
}

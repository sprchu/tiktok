package svc

import (
	"github.com/ByteDance-camp/TickTalk/api/config"
	"github.com/ByteDance-camp/TickTalk/api/user/internal/middleware"
	user "github.com/ByteDance-camp/TickTalk/user/rpc/service"

	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config         config.Config
	AuthMiddleware rest.Middleware
	UserRpc        user.Service
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:         c,
		AuthMiddleware: middleware.NewAuthMiddleware(c.Auth.AccessSecret).Handle,
		UserRpc:        user.NewService(zrpc.MustNewClient(c.UserRpc)),
	}
}

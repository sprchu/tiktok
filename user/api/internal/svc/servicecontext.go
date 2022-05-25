package svc

import (
	"github.com/ByteDance-camp/TickTalk/user/api/internal/config"
	"github.com/ByteDance-camp/TickTalk/user/api/internal/middleware"
	"github.com/ByteDance-camp/TickTalk/user/rpc/user"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config         config.Config
	AuthMiddleware rest.Middleware
	UserRpc        user.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:         c,
		AuthMiddleware: middleware.NewAuthMiddleware(c.Auth.AccessSecret).Handle,
		UserRpc:        user.NewUser(zrpc.MustNewClient(c.UserRpc)),
	}
}

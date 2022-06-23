package svc

import (
	"github.com/sprchu/tiktok/api/config"
	"github.com/sprchu/tiktok/api/user/internal/middleware"
	user "github.com/sprchu/tiktok/user/rpc/service"

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

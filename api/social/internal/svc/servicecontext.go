package svc

import (
	"github.com/sprchu/tiktok/api/config"
	"github.com/sprchu/tiktok/api/social/internal/middleware"
	socialsvc "github.com/sprchu/tiktok/social/rpc/service"

	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config         config.Config
	AuthMiddleware rest.Middleware
	SocialRpc      socialsvc.Service
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:         c,
		AuthMiddleware: middleware.NewAuthMiddleware(c.Auth.AccessSecret).Handle,
		SocialRpc:      socialsvc.NewService(zrpc.MustNewClient(c.SocialRpc)),
	}
}

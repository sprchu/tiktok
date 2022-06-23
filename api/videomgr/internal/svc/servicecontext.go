package svc

import (
	"github.com/sprchu/tiktok/api/config"
	"github.com/sprchu/tiktok/api/videomgr/internal/middleware"
	videosvc "github.com/sprchu/tiktok/videomgr/rpc/service"

	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config         config.Config
	AuthMiddleware rest.Middleware
	VideoRpc       videosvc.Service
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:         c,
		AuthMiddleware: middleware.NewAuthMiddleware(c.Auth.AccessSecret).Handle,
		VideoRpc:       videosvc.NewService(zrpc.MustNewClient(c.VideoRpc)),
	}
}

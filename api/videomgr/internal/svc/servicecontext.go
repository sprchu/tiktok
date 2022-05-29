package svc

import (
	"github.com/ByteDance-camp/TickTalk/api/config"
	"github.com/ByteDance-camp/TickTalk/api/videomgr/internal/middleware"
	videosvc "github.com/ByteDance-camp/TickTalk/videomgr/rpc/service"

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

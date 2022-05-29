package svc

import (
	"github.com/ByteDance-camp/TickTalk/api/config"
	"github.com/ByteDance-camp/TickTalk/api/videomgr/internal/middleware"
	"github.com/ByteDance-camp/TickTalk/videomgr/rpc/videoservice"

	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config         config.Config
	AuthMiddleware rest.Middleware
	VideoRpc       videoservice.VideoService
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:         c,
		AuthMiddleware: middleware.NewAuthMiddleware(c.Auth.AccessSecret).Handle,
		VideoRpc:       videoservice.NewVideoService(zrpc.MustNewClient(c.VideoRpc)),
	}
}

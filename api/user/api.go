package user

import (
	"github.com/ByteDance-camp/TickTalk/api/config"
	"github.com/ByteDance-camp/TickTalk/api/user/internal/handler"
	"github.com/ByteDance-camp/TickTalk/api/user/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func InitApi(cfg config.Config, server *rest.Server) {
	ctx := svc.NewServiceContext(cfg)
	handler.RegisterHandlers(server, ctx)
}

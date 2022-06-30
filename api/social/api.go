package social

import (
	"github.com/sprchu/tiktok/api/config"
	"github.com/sprchu/tiktok/api/social/internal/handler"
	"github.com/sprchu/tiktok/api/social/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func InitApi(cfg config.Config, server *rest.Server) {
	ctx := svc.NewServiceContext(cfg)
	handler.RegisterHandlers(server, ctx)
}

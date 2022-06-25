package svc

import (
	user "github.com/sprchu/tiktok/user/rpc/service"
	"github.com/sprchu/tiktok/videomgr/model"
	"github.com/sprchu/tiktok/videomgr/rpc/internal/config"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config     config.Config
	VideoModel model.VideoModel
	UserRpc    user.Service
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:     c,
		VideoModel: model.NewVideoModel(conn, c.CacheRedis),
		UserRpc:    user.NewService(zrpc.MustNewClient(c.UserRpc)),
	}
}

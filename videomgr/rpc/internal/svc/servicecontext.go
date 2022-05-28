package svc

import (
	"github.com/ByteDance-camp/TickTalk/user/rpc/user"
	"github.com/ByteDance-camp/TickTalk/videomgr/model"
	"github.com/ByteDance-camp/TickTalk/videomgr/rpc/internal/config"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config     config.Config
	VideoModel model.VideoModel
	UserRpc    user.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:     c,
		VideoModel: model.NewVideoModel(conn, c.CacheRedis),
		UserRpc:    user.NewUser(zrpc.MustNewClient(c.UserRpc)),
	}
}

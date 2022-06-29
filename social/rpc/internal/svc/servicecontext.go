package svc

import (
	"github.com/sprchu/tiktok/social/model"
	"github.com/sprchu/tiktok/social/rpc/internal/config"
	userModel "github.com/sprchu/tiktok/user/model"
	user "github.com/sprchu/tiktok/user/rpc/service"
	videoModel "github.com/sprchu/tiktok/videomgr/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config        config.Config
	RelationModel model.RelationModel
	FavoriteModel model.FavoriteModel
	CommentModel  model.CommentModel
	UserModel     userModel.UserModel
	VideoModel    videoModel.VideoModel
	UserRpc       user.Service
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:        c,
		RelationModel: model.NewRelationModel(conn, c.CacheRedis),
		FavoriteModel: model.NewFavoriteModel(conn, c.CacheRedis),
		CommentModel:  model.NewCommentModel(conn, c.CacheRedis),
		UserModel:     userModel.NewUserModel(conn, c.CacheRedis),
		VideoModel:    videoModel.NewVideoModel(conn, c.CacheRedis),
		UserRpc:       user.NewService(zrpc.MustNewClient(c.UserRpc)),
	}
}

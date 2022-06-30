package favorite

import (
	"context"

	"github.com/sprchu/tiktok/api/social/internal/middleware"
	"github.com/sprchu/tiktok/api/social/internal/svc"
	"github.com/sprchu/tiktok/api/social/internal/types"
	"github.com/sprchu/tiktok/servebase/errno"
	"github.com/sprchu/tiktok/social/rpc/types/social"

	"github.com/zeromicro/go-zero/core/logx"
)

type FavoriteActionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFavoriteActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FavoriteActionLogic {
	return &FavoriteActionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FavoriteActionLogic) FavoriteAction(req *types.FavoriteActionRequest) (resp *types.FavoriteActionResponse, err error) {
	resp = &types.FavoriteActionResponse{}

	uid, ok := l.ctx.Value(middleware.UID("uid")).(int64)
	if !ok {
		return resp, errno.NewErrNo(errno.PublishActionErrCode, "favorite action need uid")
	}
	_, err = l.svcCtx.SocialRpc.FavoriteAction(l.ctx, &social.FavoriteActionRequest{
		UserId:     uid,
		VideoId:    req.VideoId,
		ActionType: req.ActionType,
	})
	if err != nil {
		l.Logger.Errorf("Favorite action err: %w", err)
		return resp, errno.NewErrNo(errno.FavoriteActionErrCode, err.Error())
	}

	return resp, nil
}

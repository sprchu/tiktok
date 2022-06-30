package favorite

import (
	"context"

	"github.com/sprchu/tiktok/api/social/internal/logic"
	"github.com/sprchu/tiktok/api/social/internal/svc"
	"github.com/sprchu/tiktok/api/social/internal/types"
	"github.com/sprchu/tiktok/servebase/errno"
	"github.com/sprchu/tiktok/social/rpc/types/social"

	"github.com/zeromicro/go-zero/core/logx"
)

type FavoriteListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFavoriteListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FavoriteListLogic {
	return &FavoriteListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FavoriteListLogic) FavoriteList(req *types.FavoriteListRequest) (resp *types.FavoriteListResponse, err error) {
	resp = &types.FavoriteListResponse{}

	res, err := l.svcCtx.SocialRpc.FavoriteList(l.ctx, &social.FavoriteListRequest{UserId: req.UserId})
	if err != nil {
		l.Logger.Errorf("List favorite err: %w", err)
		return resp, errno.NewErrNo(errno.FavoriteListErrCode, err.Error())
	}

	videos := make([]types.Video, 0, len(res.VideoList))
	for _, v := range res.VideoList {
		videos = append(videos, *logic.VideoInfoResolver(v))
	}

	resp.VideoList = videos
	return resp, nil
}

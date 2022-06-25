package logic

import (
	"context"

	"github.com/sprchu/tiktok/user/rpc/types/user"
	"github.com/sprchu/tiktok/videomgr/rpc/internal/svc"
	"github.com/sprchu/tiktok/videomgr/rpc/types/videomanager"

	"github.com/zeromicro/go-zero/core/logx"
)

const maxVideosCount = 30

type FeedLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFeedLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FeedLogic {
	return &FeedLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FeedLogic) Feed(in *videomanager.FeedRequest) (*videomanager.FeedResponse, error) {
	videos, err := l.svcCtx.VideoModel.MGetLatest(l.ctx, maxVideosCount)
	if err != nil {
		return nil, err
	}

	res := make([]*videomanager.Video, 0, len(videos))
	for _, video := range videos {
		resp, err := l.svcCtx.UserRpc.UserInfo(l.ctx, &user.UserInfoRequest{
			UserId: video.UserId,
		})
		if err != nil {
			return nil, err
		}
		vd := videoResolver(&video)
		vd.Author = resp.UserInfo
		res = append(res, vd)
	}

	return &videomanager.FeedResponse{
		NextTime:  nil,
		VideoList: res,
	}, nil
}

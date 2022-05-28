package logic

import (
	"context"

	"github.com/ByteDance-camp/TickTalk/user/rpc/types/user"
	"github.com/ByteDance-camp/TickTalk/videomgr/rpc/internal/svc"
	"github.com/ByteDance-camp/TickTalk/videomgr/rpc/types/videomanager"

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
		vd.Author = userInfoResolver(resp.UserInfo)
		res = append(res, vd)
	}

	nextTime := new(int64)
	if len(videos) != 0 {
		timestamp := videos[0].CreateTime.Unix()
		nextTime = &timestamp
	}
	return &videomanager.FeedResponse{
		NextTime:  nextTime,
		VideoList: res,
	}, nil
}

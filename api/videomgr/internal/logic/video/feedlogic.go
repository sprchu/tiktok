package video

import (
	"context"
	"time"

	"github.com/sprchu/tiktok/api/videomgr/internal/svc"
	"github.com/sprchu/tiktok/api/videomgr/internal/types"
	"github.com/sprchu/tiktok/servebase/errno"
	"github.com/sprchu/tiktok/videomgr/rpc/types/videomanager"

	"github.com/zeromicro/go-zero/core/logx"
)

const timeLayout = "2006-01-02 15:04:05"

type FeedLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFeedLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FeedLogic {
	return &FeedLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FeedLogic) Feed(req *types.FeedRequest) (resp *types.FeedResponse, err error) {
	resp = &types.FeedResponse{}
	latestTime := time.Now().Unix()
	if req.LatestTime != "" {
		parsedTime, err := time.Parse(timeLayout, req.LatestTime)
		if err == nil {
			latestTime = parsedTime.Unix()
		}
	}

	res, err := l.svcCtx.VideoRpc.Feed(l.ctx, &videomanager.FeedRequest{
		LatestTime: &latestTime,
	})
	if err != nil {
		return resp, errno.NewErrNo(errno.FeedErrCode, err.Error())
	}

	videos := make([]types.Video, 0, len(res.VideoList))
	for _, vd := range res.VideoList {
		videos = append(videos, *videoResolver(vd))
	}

	if res.NextTime != nil {
		resp.NextTime = time.Unix(*res.NextTime, 0).Local().Format(timeLayout)
	}
	resp.VideoList = videos
	return resp, nil
}

package video

import (
	"context"

	"github.com/ByteDance-camp/TickTalk/servebase/errno"
	"github.com/ByteDance-camp/TickTalk/videomgr/api/internal/svc"
	"github.com/ByteDance-camp/TickTalk/videomgr/api/internal/types"
	"github.com/ByteDance-camp/TickTalk/videomgr/rpc/types/videomanager"

	"github.com/zeromicro/go-zero/core/logx"
)

type PublishListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPublishListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PublishListLogic {
	return &PublishListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PublishListLogic) PublishList(req *types.PublishListRequest) (resp *types.PublishListResponse, err error) {
	resp = &types.PublishListResponse{}
	res, err := l.svcCtx.VideoRpc.PublishList(l.ctx, &videomanager.PublishListRequest{
		UserId: req.UserID,
	})
	if err != nil {
		return resp, errno.NewErrNo(errno.PublishListErrCode, err.Error())
	}

	videos := make([]types.Video, 0, len(res.VideoList))
	for _, vd := range res.VideoList {
		videos = append(videos, *videoResolver(vd))
	}

	return &types.PublishListResponse{
		VideoList: videos,
	}, nil
}

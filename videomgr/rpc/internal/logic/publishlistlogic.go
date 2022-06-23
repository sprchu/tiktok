package logic

import (
	"context"

	"github.com/sprchu/tiktok/user/rpc/types/user"
	"github.com/sprchu/tiktok/videomgr/rpc/internal/svc"
	"github.com/sprchu/tiktok/videomgr/rpc/types/videomanager"

	"github.com/zeromicro/go-zero/core/logx"
)

type PublishListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPublishListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PublishListLogic {
	return &PublishListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PublishListLogic) PublishList(in *videomanager.PublishListRequest) (*videomanager.PublishListResponse, error) {
	videos, err := l.svcCtx.VideoModel.GetByUser(l.ctx, in.UserId)
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

	return &videomanager.PublishListResponse{
		VideoList: res,
	}, nil
}

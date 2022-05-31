package logic

import (
	"context"

	"github.com/ByteDance-camp/TickTalk/videomgr/model"
	"github.com/ByteDance-camp/TickTalk/videomgr/rpc/internal/svc"
	"github.com/ByteDance-camp/TickTalk/videomgr/rpc/types/videomanager"

	"github.com/zeromicro/go-zero/core/logx"
)

type PublishActionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPublishActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PublishActionLogic {
	return &PublishActionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PublishActionLogic) PublishAction(in *videomanager.PublishActionRequest) (*videomanager.PublishActionResponse, error) {
	_, err := l.svcCtx.VideoModel.Insert(l.ctx, &model.Video{
		Title:    in.Title,
		FileUrl:  in.Url,
		CoverUrl: in.Cover,
		UserId:   in.UserId,
	})
	if err != nil {
		return nil, err
	}

	return &videomanager.PublishActionResponse{}, nil
}

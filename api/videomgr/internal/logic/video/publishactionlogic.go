package video

import (
	"context"

	"github.com/ByteDance-camp/TickTalk/api/videomgr/internal/middleware"
	"github.com/ByteDance-camp/TickTalk/api/videomgr/internal/svc"
	"github.com/ByteDance-camp/TickTalk/api/videomgr/internal/types"
	"github.com/ByteDance-camp/TickTalk/servebase/errno"
	"github.com/ByteDance-camp/TickTalk/videomgr/rpc/types/videomanager"

	"github.com/zeromicro/go-zero/core/logx"
)

type PublishActionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPublishActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PublishActionLogic {
	return &PublishActionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PublishActionLogic) PublishAction(req *types.PublishActionRequest) (resp *types.PublishActionResponse, err error) {
	resp = &types.PublishActionResponse{}
	uid, ok := l.ctx.Value(middleware.UID("uid")).(int64)
	if !ok {
		return resp, errno.NewErrNo(errno.PublishActionErrCode, "publish action need uid")
	}

	_, err = l.svcCtx.VideoRpc.PublishAction(l.ctx, &videomanager.PublishActionRequest{
		UserId: uid,
		Title:  req.Title,
		Url:    req.Url,
		Cover:  req.Cover,
	})
	if err != nil {
		return resp, errno.NewErrNo(errno.PublishActionErrCode, err.Error())
	}

	return
}

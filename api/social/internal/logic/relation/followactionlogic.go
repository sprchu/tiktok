package relation

import (
	"context"

	"github.com/sprchu/tiktok/api/social/internal/middleware"
	"github.com/sprchu/tiktok/api/social/internal/svc"
	"github.com/sprchu/tiktok/api/social/internal/types"
	"github.com/sprchu/tiktok/servebase/errno"
	"github.com/sprchu/tiktok/social/rpc/types/social"

	"github.com/zeromicro/go-zero/core/logx"
)

type FollowActionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFollowActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FollowActionLogic {
	return &FollowActionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FollowActionLogic) FollowAction(req *types.FollowActionRequest) (resp *types.FollowActionResponse, err error) {
	resp = &types.FollowActionResponse{}

	uid, ok := l.ctx.Value(middleware.UID("uid")).(int64)
	if !ok {
		return resp, errno.NewErrNo(errno.PublishActionErrCode, "follow action need uid")
	}

	_, err = l.svcCtx.SocialRpc.FollowAction(l.ctx, &social.FollowActionRequest{
		UserId:     uid,
		ToUserId:   req.ToUserId,
		ActionType: req.ActionType,
	})
	if err != nil {
		l.Logger.Errorf("Follow action err: %w", err)
		return resp, errno.NewErrNo(errno.FollowActionErrCode, err.Error())
	}

	return resp, nil
}

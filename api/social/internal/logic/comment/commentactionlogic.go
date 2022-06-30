package comment

import (
	"context"

	"github.com/sprchu/tiktok/api/social/internal/logic"
	"github.com/sprchu/tiktok/api/social/internal/svc"
	"github.com/sprchu/tiktok/api/social/internal/types"
	"github.com/sprchu/tiktok/servebase/errno"
	"github.com/sprchu/tiktok/social/rpc/types/social"

	"github.com/zeromicro/go-zero/core/logx"
)

type CommentActionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCommentActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommentActionLogic {
	return &CommentActionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CommentActionLogic) CommentAction(req *types.CommentActionRequest) (resp *types.CommentActionResponse, err error) {
	resp = &types.CommentActionResponse{}

	if req.ActionType == 2 {
		_, err = l.svcCtx.SocialRpc.CommentAction(l.ctx, &social.CommentActionRequest{
			UserId:     req.UserId,
			VideoId:    req.VideoId,
			ActionType: req.ActionType,
			CommentId:  &req.CommentId,
		})
		if err != nil {
			l.Logger.Errorf("Delete comment err: %w", err)
			return resp, errno.NewErrNo(errno.CommentActionErrCode, err.Error())
		}
		return resp, nil
	}

	res, err := l.svcCtx.SocialRpc.CommentAction(l.ctx, &social.CommentActionRequest{
		UserId:      req.UserId,
		VideoId:     req.VideoId,
		ActionType:  req.ActionType,
		CommentText: &req.CommentText,
	})
	if err != nil {
		l.Logger.Errorf("Create comment err: %w", err)
		return resp, errno.NewErrNo(errno.CommentActionErrCode, err.Error())
	}

	resp.Comment = logic.CommentInfoResolver(res.Comment)

	return resp, nil
}

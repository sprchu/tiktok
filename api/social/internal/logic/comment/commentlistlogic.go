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

type CommentListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCommentListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommentListLogic {
	return &CommentListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CommentListLogic) CommentList(req *types.CommentListRequest) (resp *types.CommentListResponse, err error) {
	resp = &types.CommentListResponse{}

	res, err := l.svcCtx.SocialRpc.CommentList(l.ctx, &social.CommentListRequest{VideoId: req.VideoId})
	if err != nil {
		l.Logger.Errorf("List comment err: %w", err)
		return resp, errno.NewErrNo(errno.CommentListErrCode, err.Error())
	}

	comments := make([]types.Comment, 0, len(res.CommentList))
	for _, v := range res.CommentList {
		comments = append(comments, *logic.CommentInfoResolver(v))
	}

	resp.CommentList = comments

	return resp, nil
}

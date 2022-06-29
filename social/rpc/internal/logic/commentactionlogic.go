package logic

import (
	"context"

	"github.com/sprchu/tiktok/social/rpc/internal/svc"
	"github.com/sprchu/tiktok/social/rpc/types/social"
	"github.com/sprchu/tiktok/user/rpc/types/user"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/zeromicro/go-zero/core/logx"
)

type CommentActionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCommentActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommentActionLogic {
	return &CommentActionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CommentActionLogic) CommentAction(in *social.CommentActionRequest) (*social.CommentActionResponse, error) {
	if in.ActionType != 1 && in.ActionType != 2 {
		return nil, status.Error(codes.InvalidArgument, ErrInvalidArgument)
	}
	if in.ActionType == 2 {
		err := l.svcCtx.CommentModel.Uncomment(l.ctx, *in.CommentId, in.VideoId)
		if err != nil {
			l.Logger.Errorf("Delete comment error: %w", err)
			return nil, err
		}
		return &social.CommentActionResponse{}, nil
	}

	cid, err := l.svcCtx.CommentModel.Comment(l.ctx, in.UserId, in.VideoId, *in.CommentText)
	if err != nil {
		l.Logger.Errorf("Comment error: %w", err)
		return nil, err
	}
	comment, err := l.svcCtx.CommentModel.FindOne(l.ctx, cid)
	if err != nil {
		l.Logger.Errorf("Find comment error: %w", err)
		return nil, err
	}

	userInfoResp, err := l.svcCtx.UserRpc.UserInfo(l.ctx, &user.UserInfoRequest{
		UserId: comment.UserId,
	})
	if err != nil {
		l.Logger.Errorf("Get user info error: %w", err)
		return nil, err
	}

	return &social.CommentActionResponse{
		Comment: &social.Comment{
			Id:         comment.Id,
			User:       userInfoResp.UserInfo,
			Content:    comment.Content,
			CreateDate: comment.CreateTime.Format("2006-01-02 15:04:05"),
		},
	}, nil
}

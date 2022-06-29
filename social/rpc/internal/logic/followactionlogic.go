package logic

import (
	"context"

	"github.com/sprchu/tiktok/social/rpc/internal/svc"
	"github.com/sprchu/tiktok/social/rpc/types/social"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type FollowActionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFollowActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FollowActionLogic {
	return &FollowActionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FollowActionLogic) FollowAction(in *social.FollowActionRequest) (*social.FollowActionResponse, error) {
	if in.ActionType != 1 && in.ActionType != 2 {
		return nil, status.Error(codes.InvalidArgument, ErrInvalidArgument)
	}

	if in.ActionType == 2 {
		err := l.svcCtx.RelationModel.Unfollow(l.ctx, in.UserId, in.ToUserId)
		if err != nil {
			l.Logger.Errorf("Unfollow error: %v", err)
			return nil, err
		}
		return &social.FollowActionResponse{}, nil
	}

	err := l.svcCtx.RelationModel.Follow(l.ctx, in.UserId, in.ToUserId)
	if err != nil {
		l.Logger.Errorf("Follow error: %v", err)
		return nil, err
	}

	return &social.FollowActionResponse{}, nil
}

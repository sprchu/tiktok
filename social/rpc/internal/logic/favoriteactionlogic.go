package logic

import (
	"context"

	"github.com/sprchu/tiktok/social/rpc/internal/svc"
	"github.com/sprchu/tiktok/social/rpc/types/social"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/zeromicro/go-zero/core/logx"
)

type FavoriteActionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFavoriteActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FavoriteActionLogic {
	return &FavoriteActionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FavoriteActionLogic) FavoriteAction(in *social.FavoriteActionRequest) (*social.FavoriteActionResponse, error) {
	if in.ActionType != 1 && in.ActionType != 2 {
		return nil, status.Error(codes.InvalidArgument, ErrInvalidArgument)
	}

	if in.ActionType == 2 {
		err := l.svcCtx.FavoriteModel.Unstar(l.ctx, in.UserId, in.VideoId)
		if err != nil {
			l.Logger.Errorf("Unstar error: %v", err)
			return nil, err
		}
		return &social.FavoriteActionResponse{}, nil
	}

	err := l.svcCtx.FavoriteModel.Star(l.ctx, in.UserId, in.VideoId)
	if err != nil {
		l.Logger.Errorf("Star error: %v", err)
		return nil, err
	}

	return &social.FavoriteActionResponse{}, nil
}

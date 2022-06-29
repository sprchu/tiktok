package logic

import (
	"context"

	"github.com/sprchu/tiktok/social/rpc/internal/svc"
	"github.com/sprchu/tiktok/social/rpc/types/social"
	"github.com/sprchu/tiktok/user/rpc/types/user"
	"github.com/sprchu/tiktok/videomgr/rpc/types/videomanager"

	"github.com/zeromicro/go-zero/core/logx"
)

type FavoriteListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFavoriteListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FavoriteListLogic {
	return &FavoriteListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FavoriteListLogic) FavoriteList(in *social.FavoriteListRequest) (*social.FavoriteListResponse, error) {
	vids, err := l.svcCtx.FavoriteModel.ListFavorite(l.ctx, in.UserId)
	if err != nil {
		l.Logger.Errorf("List favorite video id error: %w", err)
		return nil, err
	}
	if len(vids) == 0 {
		return &social.FavoriteListResponse{}, nil
	}
	videos, err := l.svcCtx.VideoModel.MGetByIDs(l.ctx, vids)
	if err != nil {
		l.Logger.Errorf("MGet video error: %w", err)
		return nil, err
	}

	res := make([]*videomanager.Video, 0, len(videos))
	for _, v := range videos {
		isStar, err := l.svcCtx.FavoriteModel.IsStar(l.ctx, in.UserId, v.Id)
		if err != nil {
			l.Logger.Errorf("Get favorite relation error: %w", err)
			return nil, err
		}
		user, err := l.svcCtx.UserRpc.UserInfo(l.ctx, &user.UserInfoRequest{UserId: v.UserId})
		if err != nil {
			l.Logger.Errorf("Get user info error: %w", err)
			return nil, err
		}
		isFollow, err := l.svcCtx.RelationModel.IsFollow(l.ctx, in.UserId, v.UserId)
		if err != nil {
			l.Logger.Errorf("Get relation err: %w", err)
			return nil, err
		}
		user.UserInfo.IsFollow = isFollow

		res = append(res, &videomanager.Video{
			Id:            v.Id,
			Author:        user.UserInfo,
			PlayUrl:       v.FileUrl,
			CoverUrl:      v.CoverUrl,
			FavoriteCount: v.FavoriteCount,
			CommentCount:  v.CommentCount,
			IsFavorite:    isStar,
			Title:         v.Title,
		})
	}

	return &social.FavoriteListResponse{
		VideoList: res,
	}, nil
}

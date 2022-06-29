package logic

import (
	"context"

	"github.com/sprchu/tiktok/social/rpc/internal/svc"
	"github.com/sprchu/tiktok/social/rpc/types/social"
	"github.com/sprchu/tiktok/user/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type FollowListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFollowListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FollowListLogic {
	return &FollowListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FollowListLogic) FollowList(in *social.FollowListRequest) (*social.FollowListResponse, error) {
	uids, err := l.svcCtx.RelationModel.ListFollow(l.ctx, in.UserId)
	if err != nil {
		l.Logger.Errorf("List follow user id error: %w", err)
		return nil, err
	}
	if len(uids) == 0 {
		return &social.FollowListResponse{}, nil
	}
	users, err := l.svcCtx.UserModel.MGetByIDs(l.ctx, uids)
	if err != nil {
		l.Logger.Errorf("MGet user error: %w", err)
		return nil, err
	}

	res := make([]*user.UserInfo, 0, len(users))
	for _, v := range users {
		isFollow, err := l.svcCtx.RelationModel.IsFollow(l.ctx, in.UserId, v.Id)
		if err != nil {
			l.Logger.Errorf("Get follow relation error: %w", err)
			return nil, err
		}
		res = append(res, &user.UserInfo{
			Id:            v.Id,
			Name:          v.Username,
			FollowCount:   &v.FollowCount,
			FollowerCount: &v.FollowerCount,
			IsFollow:      isFollow,
		})
	}

	return &social.FollowListResponse{
		UserList: res,
	}, nil
}

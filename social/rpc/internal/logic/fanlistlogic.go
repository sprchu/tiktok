package logic

import (
	"context"

	"github.com/sprchu/tiktok/social/rpc/internal/svc"
	"github.com/sprchu/tiktok/social/rpc/types/social"
	"github.com/sprchu/tiktok/user/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type FanListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFanListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FanListLogic {
	return &FanListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FanListLogic) FanList(in *social.FanListRequest) (*social.FanListResponse, error) {
	uids, err := l.svcCtx.RelationModel.ListFollower(l.ctx, in.UserId)
	if err != nil {
		l.Logger.Errorf("List follower error: %w", err)
		return nil, err
	}
	if len(uids) == 0 {
		return &social.FanListResponse{}, nil
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

	return &social.FanListResponse{
		UserList: res,
	}, nil
}

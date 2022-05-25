package logic

import (
	"context"

	"github.com/ByteDance-camp/TickTalk/user/rpc/internal/svc"
	"github.com/ByteDance-camp/TickTalk/user/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserInfoLogic) UserInfo(in *user.UserInfoRequest) (*user.UserInfoResponse, error) {
	res, err := l.svcCtx.UserModel.FindOne(l.ctx, in.UserId)
	if err != nil {
		return nil, err
	}

	return &user.UserInfoResponse{
		UserInfo: &user.UserInfo{
			Id:            res.Id,
			Name:          res.Username,
			FollowCount:   &res.FollowCount,
			FollowerCount: &res.FollowerCount,
		},
	}, nil
}

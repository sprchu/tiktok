package user

import (
	"context"

	"github.com/ByteDance-camp/TickTalk/api/user/internal/svc"
	"github.com/ByteDance-camp/TickTalk/api/user/internal/types"
	"github.com/ByteDance-camp/TickTalk/servebase/errno"
	"github.com/ByteDance-camp/TickTalk/user/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo(req *types.UserInfoRequest) (resp *types.UserInfoResponse, err error) {
	resp = &types.UserInfoResponse{}
	rpcResp, err := l.svcCtx.UserRpc.UserInfo(l.ctx, &user.UserInfoRequest{
		UserId: req.UserId,
	})
	if err != nil {
		return resp, errno.NewErrNo(errno.GetUserInfoErrCode, err.Error())
	}

	resp.User = types.User{
		Id:            rpcResp.UserInfo.Id,
		Name:          rpcResp.UserInfo.Name,
		FollowCount:   *rpcResp.UserInfo.FollowCount,
		FollowerCount: *rpcResp.UserInfo.FollowerCount,
		IsFollow:      rpcResp.UserInfo.IsFollow,
	}
	return resp, nil
}

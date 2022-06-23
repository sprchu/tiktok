package user

import (
	"context"

	"github.com/sprchu/tiktok/api/user/internal/svc"
	"github.com/sprchu/tiktok/api/user/internal/types"
	"github.com/sprchu/tiktok/servebase"
	"github.com/sprchu/tiktok/servebase/errno"
	"github.com/sprchu/tiktok/user/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginRequest) (resp *types.LoginResponse, err error) {
	resp = &types.LoginResponse{}
	rpcResp, err := l.svcCtx.UserRpc.Login(l.ctx, &user.LoginRequest{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		return resp, errno.NewErrNo(errno.LoginErrCode, err.Error())
	}

	token, err := servebase.GenerateToken(
		l.svcCtx.Config.Auth.AccessSecret,
		l.svcCtx.Config.Auth.AccessExpire,
		rpcResp.UserId,
	)
	if err != nil {
		return resp, errno.NewErrNo(errno.AuthErrCode, err.Error())
	}

	resp.UserId = rpcResp.UserId
	resp.Token = token
	return resp, nil
}

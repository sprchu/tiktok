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

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterRequest) (resp *types.RegisterResponse, err error) {
	resp = &types.RegisterResponse{}
	rpcResp, err := l.svcCtx.UserRpc.Register(l.ctx, &user.RegisterRequest{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		return resp, errno.NewErrNo(errno.RegisterErrCode, err.Error())
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

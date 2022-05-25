package user

import (
	"context"

	"github.com/ByteDance-camp/TickTalk/servebase"
	"github.com/ByteDance-camp/TickTalk/servebase/errno"
	"github.com/ByteDance-camp/TickTalk/user/api/internal/svc"
	"github.com/ByteDance-camp/TickTalk/user/api/internal/types"
	"github.com/ByteDance-camp/TickTalk/user/rpc/types/user"

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

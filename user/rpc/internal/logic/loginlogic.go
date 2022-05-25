package logic

import (
	"context"
	"errors"

	"github.com/ByteDance-camp/TickTalk/user/rpc/internal/svc"
	"github.com/ByteDance-camp/TickTalk/user/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *user.LoginRequest) (*user.LoginResponse, error) {
	res, err := l.svcCtx.UserModel.FindOneByUsername(l.ctx, in.Username)
	if err != nil {
		return nil, err
	}
	if res.Password != in.Password {
		return nil, errors.New("invalid password")
	}

	return &user.LoginResponse{UserId: res.Id}, nil
}

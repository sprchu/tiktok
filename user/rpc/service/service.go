// Code generated by goctl. DO NOT EDIT!
// Source: user.proto

package service

import (
	"context"

	"github.com/sprchu/tiktok/user/rpc/types/user"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	LoginRequest     = user.LoginRequest
	LoginResponse    = user.LoginResponse
	RegisterRequest  = user.RegisterRequest
	RegisterResponse = user.RegisterResponse
	UserInfo         = user.UserInfo
	UserInfoRequest  = user.UserInfoRequest
	UserInfoResponse = user.UserInfoResponse

	Service interface {
		Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
		Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
		UserInfo(ctx context.Context, in *UserInfoRequest, opts ...grpc.CallOption) (*UserInfoResponse, error)
	}

	defaultService struct {
		cli zrpc.Client
	}
)

func NewService(cli zrpc.Client) Service {
	return &defaultService{
		cli: cli,
	}
}

func (m *defaultService) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	client := user.NewServiceClient(m.cli.Conn())
	return client.Login(ctx, in, opts...)
}

func (m *defaultService) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error) {
	client := user.NewServiceClient(m.cli.Conn())
	return client.Register(ctx, in, opts...)
}

func (m *defaultService) UserInfo(ctx context.Context, in *UserInfoRequest, opts ...grpc.CallOption) (*UserInfoResponse, error) {
	client := user.NewServiceClient(m.cli.Conn())
	return client.UserInfo(ctx, in, opts...)
}

// Code generated by goctl. DO NOT EDIT!
// Source: videomanager.proto

package videoservice

import (
	"context"

	"github.com/ByteDance-camp/TickTalk/videomgr/rpc/types/videomanager"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	FeedRequest           = videomanager.FeedRequest
	FeedResponse          = videomanager.FeedResponse
	PublishActionRequest  = videomanager.PublishActionRequest
	PublishActionResponse = videomanager.PublishActionResponse
	PublishListRequest    = videomanager.PublishListRequest
	PublishListResponse   = videomanager.PublishListResponse
	UserInfo              = videomanager.UserInfo
	Video                 = videomanager.Video

	VideoService interface {
		Feed(ctx context.Context, in *FeedRequest, opts ...grpc.CallOption) (*FeedResponse, error)
		PublishAction(ctx context.Context, in *PublishActionRequest, opts ...grpc.CallOption) (*PublishActionResponse, error)
		PublishList(ctx context.Context, in *PublishListRequest, opts ...grpc.CallOption) (*PublishListResponse, error)
	}

	defaultVideoService struct {
		cli zrpc.Client
	}
)

func NewVideoService(cli zrpc.Client) VideoService {
	return &defaultVideoService{
		cli: cli,
	}
}

func (m *defaultVideoService) Feed(ctx context.Context, in *FeedRequest, opts ...grpc.CallOption) (*FeedResponse, error) {
	client := videomanager.NewVideoServiceClient(m.cli.Conn())
	return client.Feed(ctx, in, opts...)
}

func (m *defaultVideoService) PublishAction(ctx context.Context, in *PublishActionRequest, opts ...grpc.CallOption) (*PublishActionResponse, error) {
	client := videomanager.NewVideoServiceClient(m.cli.Conn())
	return client.PublishAction(ctx, in, opts...)
}

func (m *defaultVideoService) PublishList(ctx context.Context, in *PublishListRequest, opts ...grpc.CallOption) (*PublishListResponse, error) {
	client := videomanager.NewVideoServiceClient(m.cli.Conn())
	return client.PublishList(ctx, in, opts...)
}
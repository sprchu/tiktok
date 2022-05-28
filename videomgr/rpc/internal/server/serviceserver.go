// Code generated by goctl. DO NOT EDIT!
// Source: videomanager.proto

package server

import (
	"context"

	"github.com/ByteDance-camp/TickTalk/videomgr/rpc/internal/logic"
	"github.com/ByteDance-camp/TickTalk/videomgr/rpc/internal/svc"
	"github.com/ByteDance-camp/TickTalk/videomgr/rpc/types/videomanager"
)

type ServiceServer struct {
	svcCtx *svc.ServiceContext
	videomanager.UnimplementedServiceServer
}

func NewServiceServer(svcCtx *svc.ServiceContext) *ServiceServer {
	return &ServiceServer{
		svcCtx: svcCtx,
	}
}

func (s *ServiceServer) Feed(ctx context.Context, in *videomanager.FeedRequest) (*videomanager.FeedResponse, error) {
	l := logic.NewFeedLogic(ctx, s.svcCtx)
	return l.Feed(in)
}

func (s *ServiceServer) PublishAction(ctx context.Context, in *videomanager.PublishActionRequest) (*videomanager.PublishActionResponse, error) {
	l := logic.NewPublishActionLogic(ctx, s.svcCtx)
	return l.PublishAction(in)
}

func (s *ServiceServer) PublishList(ctx context.Context, in *videomanager.PublishListRequest) (*videomanager.PublishListResponse, error) {
	l := logic.NewPublishListLogic(ctx, s.svcCtx)
	return l.PublishList(in)
}

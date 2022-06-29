package logic

import (
	"context"

	"github.com/sprchu/tiktok/social/rpc/internal/svc"
	"github.com/sprchu/tiktok/social/rpc/types/social"
	"github.com/sprchu/tiktok/user/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type CommentListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCommentListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommentListLogic {
	return &CommentListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CommentListLogic) CommentList(in *social.CommentListRequest) (*social.CommentListResponse, error) {
	comments, err := l.svcCtx.CommentModel.ListByVideo(l.ctx, in.VideoId)
	if err != nil {
		l.Logger.Errorf("List comment by video failed: %w", err)
		return nil, err
	}
	res := make([]*social.Comment, 0, len(comments))
	for _, v := range comments {
		user, err := l.svcCtx.UserRpc.UserInfo(l.ctx, &user.UserInfoRequest{UserId: v.UserId})
		if err != nil {
			l.Logger.Errorf("Get user info failed: %w", err)
			return nil, err
		}
		isFollow, err := l.svcCtx.RelationModel.IsFollow(l.ctx, in.UserId, v.UserId)
		if err != nil {
			l.Logger.Errorf("Get relation failed: %w", err)
			return nil, err
		}
		user.UserInfo.IsFollow = isFollow

		res = append(res, &social.Comment{
			Id:         v.Id,
			User:       user.UserInfo,
			Content:    v.Content,
			CreateDate: v.CreateTime.Format("2006-01-02 15:04:05"),
		})
	}

	return &social.CommentListResponse{
		CommentList: res,
	}, nil
}

package relation

import (
	"context"

	"github.com/sprchu/tiktok/api/social/internal/logic"
	"github.com/sprchu/tiktok/api/social/internal/svc"
	"github.com/sprchu/tiktok/api/social/internal/types"
	"github.com/sprchu/tiktok/servebase/errno"
	"github.com/sprchu/tiktok/social/rpc/types/social"

	"github.com/zeromicro/go-zero/core/logx"
)

type FollowListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFollowListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FollowListLogic {
	return &FollowListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FollowListLogic) FollowList(req *types.FollowListRequest) (resp *types.FollowListResponse, err error) {
	resp = &types.FollowListResponse{}

	res, err := l.svcCtx.SocialRpc.FollowList(l.ctx, &social.FollowListRequest{UserId: req.UserId})
	if err != nil {
		l.Logger.Errorf("List follow err: %w", err)
		return resp, errno.NewErrNo(errno.FollowListErrCode, err.Error())
	}

	users := make([]types.User, 0, len(res.UserList))
	for _, v := range res.UserList {
		users = append(users, *logic.UserInfoResolver(v))
	}
	resp.UserList = users

	return resp, nil
}

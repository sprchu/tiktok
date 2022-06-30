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

type FansListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFansListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FansListLogic {
	return &FansListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FansListLogic) FansList(req *types.FansListRequest) (resp *types.FansListResponse, err error) {
	resp = &types.FansListResponse{}

	res, err := l.svcCtx.SocialRpc.FanList(l.ctx, &social.FanListRequest{
		UserId: req.UserId,
	})
	if err != nil {
		l.Logger.Errorf("List fans error: %w", err)
		return resp, errno.NewErrNo(errno.FollowerListErrCode, err.Error())
	}

	users := make([]types.User, 0, len(res.UserList))
	for _, v := range res.UserList {
		users = append(users, *logic.UserInfoResolver(v))
	}

	resp.UserList = users
	return resp, nil
}

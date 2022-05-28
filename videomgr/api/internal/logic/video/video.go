package video

import (
	"github.com/ByteDance-camp/TickTalk/videomgr/api/internal/types"
	"github.com/ByteDance-camp/TickTalk/videomgr/rpc/types/videomanager"
)

func videoResolver(video *videomanager.Video) *types.Video {
	return &types.Video{
		Id:            video.Id,
		Author:        *userInfoResolver(video.Author),
		PlayUrl:       video.PlayUrl,
		CoverUrl:      video.CoverUrl,
		FavoriteCount: video.FavoriteCount,
		CommentCount:  video.CommentCount,
		IsFavorite:    video.IsFavorite,
	}
}

func userInfoResolver(userInfo *videomanager.UserInfo) *types.User {
	return &types.User{
		Id:            userInfo.Id,
		Name:          userInfo.Name,
		FollowCount:   *userInfo.FollowCount,
		FollowerCount: *userInfo.FollowerCount,
		IsFollow:      userInfo.IsFollow,
	}
}

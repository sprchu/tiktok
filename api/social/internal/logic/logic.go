package logic

import (
	"github.com/sprchu/tiktok/api/social/internal/types"
	"github.com/sprchu/tiktok/social/rpc/types/social"
	"github.com/sprchu/tiktok/user/rpc/types/user"
	"github.com/sprchu/tiktok/videomgr/rpc/types/videomanager"
)

func UserInfoResolver(user *user.UserInfo) *types.User {
	return &types.User{
		Id:            user.Id,
		Name:          user.Name,
		FollowCount:   *user.FollowCount,
		FollowerCount: *user.FollowerCount,
		IsFollow:      user.IsFollow,
	}
}

func VideoInfoResolver(video *videomanager.Video) *types.Video {
	return &types.Video{
		Id:            video.Id,
		Author:        *UserInfoResolver(video.Author),
		CoverUrl:      video.CoverUrl,
		PlayUrl:       video.PlayUrl,
		FavoriteCount: video.FavoriteCount,
		CommentCount:  video.CommentCount,
		IsFavorite:    video.IsFavorite,
	}
}

func CommentInfoResolver(comment *social.Comment) *types.Comment {
	return &types.Comment{
		Id:         comment.Id,
		User:       *UserInfoResolver(comment.User),
		Content:    comment.Content,
		CreateDate: comment.CreateDate,
	}
}

package logic

import (
	um "github.com/ByteDance-camp/TickTalk/user/rpc/types/user"
	"github.com/ByteDance-camp/TickTalk/videomgr/model"
	"github.com/ByteDance-camp/TickTalk/videomgr/rpc/types/videomanager"
)

func videoResolver(video *model.Video) *videomanager.Video {
	return &videomanager.Video{
		Id:       video.Id,
		PlayUrl:  video.FileUrl,
		CoverUrl: video.CoverUrl,
		Title:    video.Title,
	}
}

func userInfoResolver(userInfo *um.UserInfo) *videomanager.UserInfo {
	return &videomanager.UserInfo{
		Id:            userInfo.Id,
		Name:          userInfo.Name,
		FollowCount:   userInfo.FollowCount,
		FollowerCount: userInfo.FollowerCount,
		IsFollow:      userInfo.IsFollow,
	}
}

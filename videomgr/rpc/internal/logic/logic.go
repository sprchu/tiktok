package logic

import (
	"github.com/sprchu/tiktok/videomgr/model"
	"github.com/sprchu/tiktok/videomgr/rpc/types/videomanager"
)

func videoResolver(video *model.Video) *videomanager.Video {
	return &videomanager.Video{
		Id:       video.Id,
		PlayUrl:  video.FileUrl,
		CoverUrl: video.CoverUrl,
		Title:    video.Title,
	}
}

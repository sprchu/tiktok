// Code generated by goctl. DO NOT EDIT.
package types

type FeedRequest struct {
	LatestTime string `form:"latest_time,optional"`
	Token      string `form:"token,optional"`
}

type FeedResponse struct {
	StatusCode string  `json:"status_code"`
	StatusMsg  string  `json:"status_msg,omitempty"`
	NextTime   string  `json:"next_time,omitempty"`
	VideoList  []Video `json:"video_list,omitempty"`
}

type PublishActionRequest struct {
	Token string `form:"token" validate:"gt=0"`
	Title string `form:"title" validate:"gt=0"`
	Url   string `form:"file_url,optional"` // 只是方便调用 rpc，跟请求中的 data 字段无关
}

type PublishActionResponse struct {
	StatusCode string `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

type PublishListRequest struct {
	Token  string `form:"token" validate:"gt=0"`
	UserID int64  `form:"user_id" validate:"gt=0"`
}

type PublishListResponse struct {
	StatusCode string  `json:"status_code"`
	StatusMsg  string  `json:"status_msg,omitempty"`
	VideoList  []Video `json:"video_list,omitempty"`
}

type User struct {
	Id            int64  `json:"id"`
	Name          string `json:"name"`
	FollowCount   int64  `json:"follow_count"`
	FollowerCount int64  `json:"follower_count"`
	IsFollow      bool   `json:"is_follow"`
}

type Video struct {
	Id            int64  `json:"id,omitempty"`
	Author        User   `json:"author,omitempty"`
	CoverUrl      string `json:"cover_url,omitempty"`
	PlayUrl       string `json:"play_url,omitempty"`
	FavoriteCount int64  `json:"favorite_count,omitempty"`
	CommentCount  int64  `json:"comment_count,omitempty"`
	IsFavorite    bool   `json:"is_favorite,omitempty"`
}
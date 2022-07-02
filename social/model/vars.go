package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var ErrNotFound = sqlx.ErrNotFound

const (
	videoTable = "video"
	userTable  = "user"
)

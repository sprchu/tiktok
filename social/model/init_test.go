package model

import (
	"context"
	"log"
	"os"
	"testing"

	user "github.com/sprchu/tiktok/user/model"
	video "github.com/sprchu/tiktok/videomgr/model"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var (
	conn       sqlx.SqlConn
	cacheRedis cache.ClusterConf
	ctx        context.Context
	u1, u2, u3 user.User
	v1, v2     video.Video
)

func TestMain(m *testing.M) {
	conn = sqlx.NewMysql("test:tiktok@tcp(:9911)/tiktok?charset=utf8mb4&parseTime=True")
	cacheRedis = []cache.NodeConf{
		{
			RedisConf: redis.RedisConf{
				Host: "localhost:9912",
			},
			Weight: 100,
		},
	}
	ctx = context.Background()

	query := "insert ignore `user` (username) values (?), (?), (?)"
	_, err := conn.ExecCtx(ctx, query, "social_1", "social_2", "social_3")
	if err != nil {
		log.Fatal(err)
	}
	var users []user.User
	query = "select * from `user` where `username` in (?, ?, ?)"
	err = conn.QueryRowsCtx(ctx, &users, query, "social_1", "social_2", "social_3")
	if err != nil {
		log.Fatal(err)
	}
	if len(users) != 3 {
		log.Fatalf("expected 3 users, got %d", len(users))
	}
	u1, u2, u3 = users[0], users[1], users[2]

	query = "insert ignore `video` (title, file_url, user_id) values (?, ?, ?), (?, ?, ?)"
	_, err = conn.ExecCtx(ctx, query, "social_1", "", u1.Id, "social_2", "", u1.Id)
	if err != nil {
		log.Fatal(err)
	}
	var videos []video.Video
	query = "select * from `video` where `title` in (?, ?)"
	err = conn.QueryRowsCtx(ctx, &videos, query, "social_1", "social_2")
	if err != nil {
		log.Fatal(err)
	}
	if len(videos) < 2 {
		log.Fatalf("expected at least 2 videos, got %d", len(videos))
	}
	v1, v2 = videos[0], videos[1]

	os.Exit(m.Run())
}

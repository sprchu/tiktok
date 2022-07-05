package model

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var (
	conn        sqlx.SqlConn
	cacheRedis  cache.ClusterConf
	ctx         context.Context
	model       VideoModel
	videos      []Video
	uid         int64
	lastVideoId int64
)

const (
	videoCount = 10
	getLatest  = 3
)

func TestMGetLatest(t *testing.T) {
	assert := require.New(t)
	vds, err := model.MGetLatest(ctx, getLatest)
	assert.Nil(err)
	assert.Len(vds, getLatest)
}

func TestMGetByIDs(t *testing.T) {
	assert := require.New(t)
	vds, err := model.MGetByIDs(ctx, []int64{lastVideoId + 1, lastVideoId + 2, lastVideoId + 3})
	assert.Nil(err)
	assert.Len(vds, 3)
}

func TestGetByUser(t *testing.T) {
	assert := require.New(t)
	vds, err := model.GetByUser(ctx, uid)
	assert.Nil(err)
	assert.Len(vds, videoCount)
}

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
	model = NewVideoModel(conn, cacheRedis)
	res, err := conn.ExecCtx(ctx, "insert into `user` (username) values (?)", fmt.Sprintf("test_%d", rand.Int()))
	if err != nil {
		log.Fatal(err)
	}
	uid, err = res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}

	err = conn.QueryRowCtx(ctx, &lastVideoId, "select id from `video` order by id desc limit 1")
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < videoCount; i++ {
		vd := Video{
			Title:  fmt.Sprintf("title %d", i),
			UserId: uid,
		}
		videos = append(videos, vd)
		_, err := model.Insert(ctx, &vd)
		if err != nil {
			log.Fatal(err)
		}
	}

	os.Exit(m.Run())
}

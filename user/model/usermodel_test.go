package model

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var (
	conn       sqlx.SqlConn
	cacheRedis cache.ClusterConf
	ctx        context.Context
)

func TestMGetByIDs(t *testing.T) {
	assert := require.New(t)
	md := NewUserModel(conn, cacheRedis)

	user1 := User{
		Username: "test1",
		Password: "ps1",
	}
	user2 := User{
		Username: "test2",
		Password: "ps2",
	}

	res, err := md.Insert(ctx, &user1)
	assert.NoError(err)
	uid1, err := res.LastInsertId()
	assert.NoError(err)

	res, err = md.Insert(ctx, &user2)
	assert.NoError(err)
	uid2, err := res.LastInsertId()
	assert.NoError(err)

	users, err := md.MGetByIDs(ctx, []int64{uid1, uid2})
	assert.NoError(err)
	assert.Equal(
		[]User{user1, user2},
		[]User{
			{
				Username: users[0].Username,
				Password: users[0].Password,
			},
			{
				Username: users[1].Username,
				Password: users[1].Password,
			},
		},
	)
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
	os.Exit(m.Run())
}

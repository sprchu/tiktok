package model

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFollowAndUnfollow(t *testing.T) {
	assert := require.New(t)
	model := NewRelationModel(conn, cacheRedis)

	assert.Nil(model.Follow(ctx, u1.Id, u2.Id))
	isFollow, err := model.IsFollow(ctx, u1.Id, u2.Id)
	assert.Nil(err)
	assert.True(isFollow)

	assert.Nil(model.Unfollow(ctx, u1.Id, u2.Id))
	isFollow, err = model.IsFollow(ctx, u1.Id, u2.Id)
	assert.Nil(err)
	assert.False(isFollow)
}

func TestListFollow(t *testing.T) {
	assert := require.New(t)
	model := NewRelationModel(conn, cacheRedis)

	assert.Nil(model.Follow(ctx, u2.Id, u3.Id))
	uids, err := model.ListFollow(ctx, u2.Id)
	assert.Nil(err)
	assert.Contains(uids, u3.Id)
}

func TestListFollower(t *testing.T) {
	assert := require.New(t)
	model := NewRelationModel(conn, cacheRedis)

	assert.Nil(model.Follow(ctx, u3.Id, u1.Id))
	assert.Nil(model.Follow(ctx, u2.Id, u1.Id))
	uids, err := model.ListFollower(ctx, u1.Id)
	assert.Nil(err)
	assert.Contains(uids, u2.Id)
	assert.Contains(uids, u3.Id)
}

package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStarAndUnstar(t *testing.T) {
	assert := assert.New(t)
	model := NewFavoriteModel(conn, cacheRedis)

	assert.Nil(model.Star(ctx, u1.Id, v1.Id))
	isStar, err := model.IsStar(ctx, u1.Id, v1.Id)
	assert.Nil(err)
	assert.True(isStar)

	assert.Nil(model.Unstar(ctx, u1.Id, v1.Id))
	isStar, err = model.IsStar(ctx, u1.Id, v1.Id)
	assert.Nil(err)
	assert.False(isStar)
}

func TestListFavorite(t *testing.T) {
	assert := assert.New(t)
	model := NewFavoriteModel(conn, cacheRedis)

	assert.Nil(model.Star(ctx, u2.Id, v1.Id))
	assert.Nil(model.Star(ctx, u2.Id, v2.Id))

	favorites, err := model.ListFavorite(ctx, u2.Id)
	assert.Nil(err)
	assert.Contains(favorites, v1.Id)
	assert.Contains(favorites, v2.Id)
}

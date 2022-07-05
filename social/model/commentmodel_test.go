package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCommentAndUncomment(t *testing.T) {
	assert := assert.New(t)
	model := NewCommentModel(conn, cacheRedis)

	cid, err := model.Comment(ctx, u1.Id, v1.Id, "comment_1")
	assert.Nil(err)
	assert.NotEmpty(cid)

	assert.Nil(model.Uncomment(ctx, u1.Id, cid))
}

func TestListByVideo(t *testing.T) {
	assert := assert.New(t)
	model := NewCommentModel(conn, cacheRedis)

	cid1, err := model.Comment(ctx, u2.Id, v1.Id, "comment_1")
	assert.Nil(err)
	comment1, err := model.FindOne(ctx, cid1)
	assert.Nil(err)
	cid2, err := model.Comment(ctx, u3.Id, v1.Id, "comment_2")
	assert.Nil(err)
	comment2, err := model.FindOne(ctx, cid2)
	assert.Nil(err)

	comments, err := model.ListByVideo(ctx, v1.Id)
	assert.Nil(err)
	assert.Contains(comments, *comment1)
	assert.Contains(comments, *comment2)
}

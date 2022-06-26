package model

import (
	"context"
	"fmt"

	video "github.com/sprchu/tiktok/videomgr/model"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ CommentModel = (*customCommentModel)(nil)

type (
	// CommentModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCommentModel.
	CommentModel interface {
		commentModel
		Comment(ctx context.Context, userID, videoID int64, content string) (int64, error)
		Uncomment(ctx context.Context, id, videoID int64) error
		ListByVideo(ctx context.Context, videoID int64) ([]Comment, error)
	}

	customCommentModel struct {
		*defaultCommentModel
		videoModel video.VideoModel
	}
)

// NewCommentModel returns a model for the database table.
func NewCommentModel(conn sqlx.SqlConn, c cache.CacheConf) CommentModel {
	return &customCommentModel{
		defaultCommentModel: newCommentModel(conn, c),
		videoModel:          video.NewVideoModel(conn, c),
	}
}

func (m *customCommentModel) Comment(ctx context.Context, userID, videoID int64, content string) (int64, error) {
	var id int64
	err := m.CachedConn.TransactCtx(ctx, func(context.Context, sqlx.Session) error {
		vd, err := m.videoModel.FindOne(ctx, videoID)
		if err != nil {
			return err
		}
		vd.CommentCount++
		err = m.videoModel.Update(ctx, vd)
		if err != nil {
			return err
		}

		res, err := m.Insert(ctx, &Comment{
			UserId:  userID,
			VideoId: videoID,
			Content: content,
		})
		if err != nil {
			return err
		}
		id, err = res.LastInsertId()
		return err
	})
	if err != nil {
		return 0, err
	}

	return id, err
}

func (m *customCommentModel) Uncomment(ctx context.Context, id, videoID int64) error {
	return m.CachedConn.TransactCtx(ctx, func(context.Context, sqlx.Session) error {
		vd, err := m.videoModel.FindOne(ctx, videoID)
		if err != nil {
			return err
		}
		vd.CommentCount--
		err = m.videoModel.Update(ctx, vd)
		if err != nil {
			return err
		}

		return m.Delete(ctx, id)
	})
}

func (m *customCommentModel) ListByVideo(ctx context.Context, videoID int64) ([]Comment, error) {
	var comments []Comment
	query := fmt.Sprintf("select %s from %s where `video_id` = ?", commentRows, m.table)
	err := m.CachedConn.QueryRowsNoCacheCtx(ctx, &comments, query, videoID)
	switch err {
	case nil:
		return comments, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

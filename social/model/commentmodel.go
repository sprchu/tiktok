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
	err := m.CachedConn.TransactCtx(ctx, func(ctx context.Context, tx sqlx.Session) error {
		insertStat := fmt.Sprintf(
			"insert ignore %s (%s) values (?, ?, ?)",
			m.table,
			commentRowsExpectAutoSet,
		)
		res, err := tx.ExecCtx(ctx, insertStat, userID, videoID, content)
		if err != nil {
			return err
		}
		id, err = res.LastInsertId()
		if err != nil {
			return err
		}

		return updateCommentCount(ctx, tx, videoID, 1)
	})
	if err != nil {
		return 0, err
	}

	return id, err
}

func (m *customCommentModel) Uncomment(ctx context.Context, id, videoID int64) error {
	return m.CachedConn.TransactCtx(ctx, func(ctx context.Context, tx sqlx.Session) error {
		deleteStat := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		res, err := tx.ExecCtx(ctx, deleteStat, id)
		if err != nil {
			return err
		}
		affected, err := res.RowsAffected()
		if err != nil {
			return err
		}
		if affected == 0 {
			return nil
		}

		return updateCommentCount(ctx, tx, videoID, -1)
	})
}

func updateCommentCount(ctx context.Context, tx sqlx.Session, videoID int64, n int64) error {
	var vd video.Video
	query := fmt.Sprintf("select * from %s where `id` = ?", videoTable)
	err := tx.QueryRowCtx(ctx, &vd, query, videoID)
	if err != nil {
		return err
	}
	updataStat := fmt.Sprintf(
		"update %s set comment_count = %d where `id` = ?",
		videoTable,
		vd.CommentCount+n,
	)
	_, err = tx.ExecCtx(ctx, updataStat, videoID)
	return err
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

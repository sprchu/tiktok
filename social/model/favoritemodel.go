package model

import (
	"context"
	"fmt"

	video "github.com/sprchu/tiktok/videomgr/model"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ FavoriteModel = (*customFavoriteModel)(nil)

type (
	// FavoriteModel is an interface to be customized, add more methods here,
	// and implement the added methods in customFavoriteModel.
	FavoriteModel interface {
		favoriteModel
		Star(ctx context.Context, userID, videoID int64) error
		Unstar(ctx context.Context, userID, videoID int64) error
		IsStar(ctx context.Context, userID, videoID int64) (bool, error)
		ListFavorite(ctx context.Context, userID int64) ([]int64, error)
	}

	customFavoriteModel struct {
		*defaultFavoriteModel
		videoModel video.VideoModel
	}
)

// NewFavoriteModel returns a model for the database table.
func NewFavoriteModel(conn sqlx.SqlConn, c cache.CacheConf) FavoriteModel {
	return &customFavoriteModel{
		defaultFavoriteModel: newFavoriteModel(conn, c),
		videoModel:           video.NewVideoModel(conn, c),
	}
}

func (m *customFavoriteModel) Star(ctx context.Context, userID, videoID int64) error {
	return m.CachedConn.TransactCtx(ctx, func(ctx context.Context, tx sqlx.Session) error {
		insertStat := fmt.Sprintf("insert into %s (%s) values (?, ?)", m.table, favoriteRowsExpectAutoSet)
		_, err := tx.ExecCtx(ctx, insertStat, userID, videoID)
		if err != nil {
			return err
		}

		return updateFavoriteCount(ctx, tx, videoID, 1)
	})
}

func (m *customFavoriteModel) Unstar(ctx context.Context, userID, videoID int64) error {
	return m.CachedConn.TransactCtx(ctx, func(ctx context.Context, tx sqlx.Session) error {
		deleteStat := fmt.Sprintf("delete from %s where `user_id` = ? and `video_id` = ?", m.table)
		res, err := tx.ExecCtx(ctx, deleteStat, userID, videoID)
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

		return updateFavoriteCount(ctx, tx, videoID, -1)
	})
}

func updateFavoriteCount(ctx context.Context, tx sqlx.Session, videoID int64, n int64) error {
	var vd video.Video
	query := fmt.Sprintf("select * from %s where `id` = ?", videoTable)
	err := tx.QueryRowCtx(ctx, &vd, query, videoID)
	if err != nil {
		return err
	}
	updataStat := fmt.Sprintf(
		"update %s set favorite_count = %d where `id` = ?",
		videoTable,
		vd.FavoriteCount+n,
	)
	_, err = tx.ExecCtx(ctx, updataStat, videoID)
	return err
}

func (m *customFavoriteModel) IsStar(ctx context.Context, userID, videoID int64) (bool, error) {
	var cnt int64
	query := fmt.Sprintf("select count(*) from %s where `user_id` = ? and `video_id` = ?", m.table)
	err := m.CachedConn.QueryRowNoCacheCtx(ctx, &cnt, query, userID, videoID)
	if err != nil {
		return false, err
	}

	return cnt > 0, nil
}

func (m *customFavoriteModel) ListFavorite(ctx context.Context, userID int64) ([]int64, error) {
	var vids []int64
	query := fmt.Sprintf("select `video_id` from %s where `user_id` = ?", m.table)
	err := m.CachedConn.QueryRowsNoCacheCtx(ctx, &vids, query, userID)
	switch err {
	case nil:
		return vids, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

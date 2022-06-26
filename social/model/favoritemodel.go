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
	return m.CachedConn.TransactCtx(ctx, func(context.Context, sqlx.Session) error {
		vd, err := m.videoModel.FindOne(ctx, videoID)
		if err != nil {
			return err
		}
		vd.FavoriteCount++
		err = m.videoModel.Update(ctx, vd)
		if err != nil {
			return err
		}

		_, err = m.Insert(ctx, &Favorite{
			UserId:  userID,
			VideoId: videoID,
		})

		return err
	})
}

func (m *customFavoriteModel) Unstar(ctx context.Context, userID, videoID int64) error {
	return m.CachedConn.TransactCtx(ctx, func(ctx context.Context, s sqlx.Session) error {
		rd, err := m.FindOneByUserIdVideoId(ctx, userID, videoID)
		if err != nil {
			return err
		}

		vd, err := m.videoModel.FindOne(ctx, videoID)
		if err != nil {
			return err
		}
		vd.FavoriteCount--
		err = m.videoModel.Update(ctx, vd)
		if err != nil {
			return err
		}

		return m.Delete(ctx, rd.Id)
	})
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

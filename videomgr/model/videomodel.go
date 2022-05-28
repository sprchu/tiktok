package model

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ VideoModel = (*customVideoModel)(nil)

type (
	// VideoModel is an interface to be customized, add more methods here,
	// and implement the added methods in customVideoModel.
	VideoModel interface {
		videoModel
		MGetLatest(ctx context.Context, n int) ([]Video, error)
		GetByUser(ctx context.Context, uid int64) ([]Video, error)
	}

	customVideoModel struct {
		*defaultVideoModel
	}
)

// NewVideoModel returns a model for the database table.
func NewVideoModel(conn sqlx.SqlConn, c cache.CacheConf) VideoModel {
	return &customVideoModel{
		defaultVideoModel: newVideoModel(conn, c),
	}
}

func (m *customVideoModel) MGetLatest(ctx context.Context, n int) ([]Video, error) {
	var videos []Video
	query := fmt.Sprintf("select %s from %s order by create_time desc limit %d", videoRows, m.table, n)
	err := m.CachedConn.QueryRowsNoCacheCtx(ctx, &videos, query)
	switch err {
	case nil:
		return videos, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *customVideoModel) GetByUser(ctx context.Context, uid int64) ([]Video, error) {
	var videos []Video
	query := fmt.Sprintf("select %s from %s where `user_id` = ?", videoRows, m.table)
	err := m.CachedConn.QueryRowsNoCacheCtx(ctx, &videos, query, uid)
	switch err {
	case nil:
		return videos, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

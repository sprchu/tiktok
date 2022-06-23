package model

import (
	"context"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UserModel = (*customUserModel)(nil)

type (
	// UserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserModel.
	UserModel interface {
		userModel
		MGetByIDs(ctx context.Context, ids []int64) ([]User, error)
	}

	customUserModel struct {
		*defaultUserModel
	}
)

// NewUserModel returns a model for the database table.
func NewUserModel(conn sqlx.SqlConn, c cache.CacheConf) UserModel {
	return &customUserModel{
		defaultUserModel: newUserModel(conn, c),
	}
}

func (m *customUserModel) MGetByIDs(ctx context.Context, ids []int64) ([]User, error) {
	var users []User
	query := fmt.Sprintf("select * from %s where id in (%s)", m.table, strings.Trim(fmt.Sprint(ids), "[]"))
	err := m.CachedConn.QueryRowsNoCacheCtx(ctx, &users, query)
	switch err {
	case nil:
		return users, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

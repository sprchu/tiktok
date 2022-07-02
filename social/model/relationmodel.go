package model

import (
	"context"
	"fmt"

	user "github.com/sprchu/tiktok/user/model"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ RelationModel = (*customRelationModel)(nil)

type (
	// RelationModel is an interface to be customized, add more methods here,
	// and implement the added methods in customRelationModel.
	RelationModel interface {
		relationModel
		IsFollow(ctx context.Context, userID, followerID int64) (bool, error)
		Follow(ctx context.Context, userID, followID int64) error
		Unfollow(ctx context.Context, userID, followID int64) error
		ListFollow(ctx context.Context, userID int64) ([]int64, error)
		ListFollower(ctx context.Context, userID int64) ([]int64, error)
	}

	customRelationModel struct {
		*defaultRelationModel
		userModel user.UserModel
	}
)

// NewRelationModel returns a model for the database table.
func NewRelationModel(conn sqlx.SqlConn, c cache.CacheConf) RelationModel {
	return &customRelationModel{
		defaultRelationModel: newRelationModel(conn, c),
		userModel:            user.NewUserModel(conn, c),
	}
}

func (m *customRelationModel) IsFollow(ctx context.Context, userID, followID int64) (bool, error) {
	var cnt int64
	query := fmt.Sprintf("select count(*) from %s where `user_id` = ? and `follow_id` = ?", m.table)
	err := m.CachedConn.QueryRowNoCacheCtx(ctx, &cnt, query, userID, followID)
	if err != nil {
		return false, err
	}

	return cnt > 0, nil
}

func (m *customRelationModel) Follow(ctx context.Context, userID, followID int64) error {
	return m.CachedConn.TransactCtx(ctx, func(ctx context.Context, tx sqlx.Session) error {
		insertStat := fmt.Sprintf("insert into %s (%s) values (?, ?)", m.table, relationRowsExpectAutoSet)
		_, err := tx.ExecCtx(ctx, insertStat, userID, followID)
		if err != nil {
			return err
		}

		err = updateFollowCount(ctx, tx, userID, 1)
		if err != nil {
			return err
		}
		return updateFollowerCount(ctx, tx, followID, 1)
	})
}

func (m *customRelationModel) Unfollow(ctx context.Context, userID, followID int64) error {
	return m.CachedConn.TransactCtx(ctx, func(ctx context.Context, tx sqlx.Session) error {
		deleteStat := fmt.Sprintf("delete from %s where `user_id` = ? and `follow_id` = ?", m.table)
		res, err := tx.ExecCtx(ctx, deleteStat, userID, followID)
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

		err = updateFollowCount(ctx, tx, userID, -1)
		if err != nil {
			return err
		}
		return updateFollowerCount(ctx, tx, followID, -1)
	})
}

func updateFollowCount(ctx context.Context, tx sqlx.Session, userID int64, n int64) error {
	var user user.User
	query := fmt.Sprintf("select * from %s where `id` = ?", userTable)
	err := tx.QueryRowCtx(ctx, &user, query, userID)
	if err != nil {
		return err
	}

	updateStat := fmt.Sprintf("update %s set `follow_count` = ? where `id` = ?", userTable)
	_, err = tx.ExecCtx(ctx, updateStat, user.FollowCount+n, userID)
	return err
}

func updateFollowerCount(ctx context.Context, tx sqlx.Session, userID int64, n int64) error {
	var user user.User
	query := fmt.Sprintf("select * from %s where `id` = ?", userTable)
	err := tx.QueryRowCtx(ctx, &user, query, userID)
	if err != nil {
		return err
	}

	updateStat := fmt.Sprintf("update %s set `follower_count` = ? where `id` = ?", userTable)
	_, err = tx.ExecCtx(ctx, updateStat, user.FollowerCount+n, userID)
	return err
}

func (m *customRelationModel) ListFollow(ctx context.Context, userID int64) ([]int64, error) {
	query := fmt.Sprintf("select `follow_id` from %s where `user_id` = ?", m.table)
	return m.getUserIds(ctx, query, userID)
}

func (m *customRelationModel) ListFollower(ctx context.Context, userID int64) ([]int64, error) {
	query := fmt.Sprintf("select `user_id` from %s where `follow_id` = ?", m.table)
	return m.getUserIds(ctx, query, userID)
}

func (m *customRelationModel) getUserIds(ctx context.Context, query string, args ...interface{}) ([]int64, error) {
	var ids []int64
	err := m.CachedConn.QueryRowsNoCacheCtx(ctx, &ids, query, args...)
	switch err {
	case nil:
		return ids, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

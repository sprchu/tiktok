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
	return m.CachedConn.TransactCtx(ctx, func(context.Context, sqlx.Session) error {
		_, err := m.Insert(ctx, &Relation{UserId: userID, FollowId: followID})
		if err != nil {
			return err
		}

		ur, err := m.userModel.FindOne(ctx, userID)
		if err != nil {
			return err
		}
		ur.FollowCount++
		err = m.userModel.Update(ctx, ur)
		if err != nil {
			return err
		}

		fur, err := m.userModel.FindOne(ctx, followID)
		if err != nil {
			return err
		}
		fur.FollowerCount++
		err = m.userModel.Update(ctx, fur)
		if err != nil {
			return err
		}

		return nil
	})
}

func (m *customRelationModel) Unfollow(ctx context.Context, userID, followID int64) error {
	return m.CachedConn.TransactCtx(ctx, func(context.Context, sqlx.Session) error {
		r, err := m.FindOneByUserIdFollowId(ctx, userID, followID)
		if err != nil {
			return err
		}
		err = m.Delete(ctx, r.Id)
		if err != nil {
			return err
		}

		ur, err := m.userModel.FindOne(ctx, userID)
		if err != nil {
			return err
		}
		ur.FollowCount--
		err = m.userModel.Update(ctx, ur)
		if err != nil {
			return err
		}

		fur, err := m.userModel.FindOne(ctx, followID)
		if err != nil {
			return err
		}
		fur.FollowerCount--
		err = m.userModel.Update(ctx, fur)
		if err != nil {
			return err
		}

		return nil
	})
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

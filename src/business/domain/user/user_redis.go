package user

import (
	"context"
	"time"

	"github.com/adiatma85/gg-project/src/business/entity"
	"github.com/adiatma85/own-go-sdk/codes"
	"github.com/adiatma85/own-go-sdk/errors"
)

const (
	getUserByIdKey    = `gg-be:user:get:%s`
	deleteUserPattern = `gg-be:user:*`
)

func (u *user) upsertCache(ctx context.Context, key string, value entity.User, expTime time.Duration) error {
	entry, err := u.json.Marshal(value)
	if err != nil {
		return err
	}

	return u.redis.SetEX(ctx, key, string(entry), expTime)
}

func (u *user) getCache(ctx context.Context, key string) (entity.User, error) {
	var result entity.User

	entry, err := u.redis.Get(ctx, key)
	if err != nil {
		return result, err
	}

	data := []byte(entry)

	return result, u.json.Unmarshal(data, &result)
}

func (u *user) deleteCache(ctx context.Context) error {
	if err := u.redis.Del(ctx, deleteUserPattern); err != nil {
		return errors.NewWithCode(codes.CodeCacheDeleteSimpleKey, "delete cache by keys pattern")
	}
	return nil
}

package role

import (
	"context"
	"time"

	"github.com/adiatma85/gg-project/src/business/entity"
	"github.com/adiatma85/own-go-sdk/codes"
	"github.com/adiatma85/own-go-sdk/errors"
)

const (
	getRoleByIdKey    = `gg-be:role:get:%s`
	deleteRolePattern = `gg-be:role:*`
)

func (r *role) upsertCache(ctx context.Context, key string, value entity.Role, expTime time.Duration) error {
	entry, err := r.json.Marshal(value)
	if err != nil {
		return err
	}

	return r.redis.SetEX(ctx, key, string(entry), expTime)
}

func (r *role) getCache(ctx context.Context, key string) (entity.Role, error) {
	var result entity.Role

	entry, err := r.redis.Get(ctx, key)
	if err != nil {
		return result, err
	}

	data := []byte(entry)

	return result, r.json.Unmarshal(data, &result)
}

func (r *role) deleteCache(ctx context.Context) error {
	if err := r.redis.Del(ctx, deleteRolePattern); err != nil {
		return errors.NewWithCode(codes.CodeCacheDeleteSimpleKey, "delete cache by keys pattern")
	}
	return nil
}

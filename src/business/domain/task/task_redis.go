package task

import (
	"context"
	"time"

	"github.com/adiatma85/gg-project/src/business/entity"
	"github.com/adiatma85/own-go-sdk/codes"
	"github.com/adiatma85/own-go-sdk/errors"
)

const (
	getTaskByIdKey    = `gg-be:task:get:%s`
	deleteTaskPattern = `gg-be:task:*`
)

func (t *task) upsertCache(ctx context.Context, key string, value entity.Task, expTime time.Duration) error {
	entry, err := t.json.Marshal(value)
	if err != nil {
		return err
	}

	return t.redis.SetEX(ctx, key, string(entry), expTime)
}

func (t *task) getCache(ctx context.Context, key string) (entity.Task, error) {
	var result entity.Task

	entry, err := t.redis.Get(ctx, key)
	if err != nil {
		return result, err
	}

	data := []byte(entry)

	return result, t.json.Unmarshal(data, &result)
}

func (t *task) deleteCache(ctx context.Context) error {
	if err := t.redis.Del(ctx, deleteTaskPattern); err != nil {
		return errors.NewWithCode(codes.CodeCacheDeleteSimpleKey, "delete cache by keys pattern")
	}
	return nil
}

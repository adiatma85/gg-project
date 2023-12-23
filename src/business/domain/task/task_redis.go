package task

import (
	"context"
	"time"

	"github.com/adiatma85/new-go-template/src/business/entity"
	"github.com/adiatma85/own-go-sdk/codes"
	"github.com/adiatma85/own-go-sdk/errors"
)

const (
	getTaskByIdKey    = `gantenggaming:task:get:%s`
	deleteTaskPattern = `gantenggaming:task:*`
)

func (t *task) upsertCache(ctx context.Context, key string, value entity.Task, expTime time.Duration) error {
	task, err := t.json.Marshal(value)
	if err != nil {
		return err
	}

	return t.redis.SetEX(ctx, key, string(task), expTime)
}

func (t *task) getCache(ctx context.Context, key string) (entity.Task, error) {
	var result entity.Task

	task, err := t.redis.Get(ctx, key)
	if err != nil {
		return result, err
	}

	data := []byte(task)

	return result, t.json.Unmarshal(data, &result)
}

func (t *task) deleteUserCache(ctx context.Context) error {
	if err := t.redis.Del(ctx, deleteTaskPattern); err != nil {
		return errors.NewWithCode(codes.CodeCacheDeleteSimpleKey, "delete cache by keys pattern")
	}
	return nil
}

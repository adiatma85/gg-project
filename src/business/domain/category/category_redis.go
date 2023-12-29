package category

import (
	"context"
	"time"

	"github.com/adiatma85/gg-project/src/business/entity"
	"github.com/adiatma85/own-go-sdk/codes"
	"github.com/adiatma85/own-go-sdk/errors"
)

const (
	getCategoryByIdKey    = `gg-be:category:get:%s`
	deleteCategoryPattern = `gg-be:category:*`
)

func (c *category) upsertCache(ctx context.Context, key string, value entity.Category, expTime time.Duration) error {
	entry, err := c.json.Marshal(value)
	if err != nil {
		return err
	}

	return c.redis.SetEX(ctx, key, string(entry), expTime)
}

func (c *category) getCache(ctx context.Context, key string) (entity.Category, error) {
	var result entity.Category

	entry, err := c.redis.Get(ctx, key)
	if err != nil {
		return result, err
	}

	data := []byte(entry)

	return result, c.json.Unmarshal(data, &result)
}

func (c *category) deleteCache(ctx context.Context) error {
	if err := c.redis.Del(ctx, deleteCategoryPattern); err != nil {
		return errors.NewWithCode(codes.CodeCacheDeleteSimpleKey, "delete cache by keys pattern")
	}
	return nil
}

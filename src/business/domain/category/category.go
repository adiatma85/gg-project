package category

import (
	"context"

	"github.com/adiatma85/gg-project/src/business/entity"
	"github.com/adiatma85/own-go-sdk/codes"
	"github.com/adiatma85/own-go-sdk/errors"
	"github.com/adiatma85/own-go-sdk/log"
	"github.com/adiatma85/own-go-sdk/null"
	"github.com/adiatma85/own-go-sdk/parser"
	"github.com/adiatma85/own-go-sdk/redis"
	"github.com/adiatma85/own-go-sdk/sql"
)

type Interface interface {
	Create(ctx context.Context, userParam entity.CreateCategoryParam) (entity.Category, error)
	Get(ctx context.Context, params entity.CategoryParam) (entity.Category, error)
	GetList(ctx context.Context, params entity.CategoryParam) ([]entity.Category, *entity.Pagination, error)
	Update(ctx context.Context, updateParam entity.UpdateCategoryParam, selectParam entity.CategoryParam) error
}

type InitParam struct {
	Log   log.Interface
	Db    sql.Interface
	Json  parser.JSONInterface
	Redis redis.Interface
}

type category struct {
	log   log.Interface
	db    sql.Interface
	json  parser.JSONInterface
	redis redis.Interface
}

func Init(param InitParam) Interface {
	c := &category{
		log:   param.Log,
		db:    param.Db,
		json:  param.Json,
		redis: param.Redis,
	}

	return c
}

func (c *category) Create(ctx context.Context, userParam entity.CreateCategoryParam) (entity.Category, error) {
	category := entity.Category{}

	tx, err := c.db.Leader().BeginTx(ctx, "txcUser", sql.TxOptions{})
	if err != nil {
		return category, errors.NewWithCode(codes.CodeSQLTxBegin, err.Error())
	}
	defer tx.Rollback()

	tx, category, err = c.createSQLCategory(tx, userParam)
	if err != nil {
		return category, err
	}

	if err = tx.Commit(); err != nil {
		return category, errors.NewWithCode(codes.CodeSQLTxCommit, err.Error())
	}

	if err := c.deleteCache(ctx); err != nil {
		c.log.Error(ctx, err)
	}

	return c.Get(ctx, entity.CategoryParam{
		ID: null.Int64From(category.ID),
	})
}

func (c *category) Get(ctx context.Context, params entity.CategoryParam) (entity.Category, error) {
	return c.getSQLCategory(ctx, params)
}

func (c *category) GetList(ctx context.Context, params entity.CategoryParam) ([]entity.Category, *entity.Pagination, error) {
	return c.getSQLCategoryList(ctx, params)
}

func (c *category) Update(ctx context.Context, updateParam entity.UpdateCategoryParam, selectParam entity.CategoryParam) error {
	return c.updateSQLCategory(ctx, updateParam, selectParam)
}

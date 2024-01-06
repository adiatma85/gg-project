package task

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
	Create(ctx context.Context, insertParam entity.CreateTaskParam) (entity.Task, error)
	Get(ctx context.Context, params entity.TaskParam) (entity.Task, error)
	GetList(ctx context.Context, params entity.TaskParam) ([]entity.Task, *entity.Pagination, error)
	Update(ctx context.Context, updateParam entity.UpdateTaskParam, selectParam entity.TaskParam) error
}

type InitParam struct {
	Log   log.Interface
	Db    sql.Interface
	Json  parser.JSONInterface
	Redis redis.Interface
}

type task struct {
	log   log.Interface
	db    sql.Interface
	json  parser.JSONInterface
	redis redis.Interface
}

func Init(param InitParam) Interface {
	t := &task{
		log:   param.Log,
		db:    param.Db,
		json:  param.Json,
		redis: param.Redis,
	}

	return t
}

func (t *task) Create(ctx context.Context, insertParam entity.CreateTaskParam) (entity.Task, error) {
	result := entity.Task{}

	tx, err := t.db.Leader().BeginTx(ctx, "txcUser", sql.TxOptions{})
	if err != nil {
		return result, errors.NewWithCode(codes.CodeSQLTxBegin, err.Error())
	}
	defer tx.Rollback()

	tx, result, err = t.createSQLTask(tx, insertParam)
	if err != nil {
		return result, err
	}

	if err = tx.Commit(); err != nil {
		return result, errors.NewWithCode(codes.CodeSQLTxCommit, err.Error())
	}

	if err := t.deleteCache(ctx); err != nil {
		t.log.Error(ctx, err)
	}

	return t.Get(ctx, entity.TaskParam{
		ID: null.Int64From(result.ID),
	})
}

func (t *task) Get(ctx context.Context, params entity.TaskParam) (entity.Task, error) {
	return t.getSQLTask(ctx, params)
}

func (t *task) GetList(ctx context.Context, params entity.TaskParam) ([]entity.Task, *entity.Pagination, error) {
	return t.getSQLTaskList(ctx, params)
}

func (t *task) Update(ctx context.Context, updateParam entity.UpdateTaskParam, selectParam entity.TaskParam) error {
	return t.updateSQLTask(ctx, updateParam, selectParam)
}

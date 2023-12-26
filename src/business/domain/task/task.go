package task

import (
	"context"

	"github.com/adiatma85/new-go-template/src/business/entity"
	"github.com/adiatma85/own-go-sdk/codes"
	"github.com/adiatma85/own-go-sdk/errors"
	"github.com/adiatma85/own-go-sdk/log"
	"github.com/adiatma85/own-go-sdk/null"
	"github.com/adiatma85/own-go-sdk/parser"
	"github.com/adiatma85/own-go-sdk/redis"
	"github.com/adiatma85/own-go-sdk/sql"
)

type Interface interface {
	// Create(ctx context.Context)
	// Get()
	// GetList()
	// Update()
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

func (t *task) Create(ctx context.Context, taskParam entity.CreateTaskParam) (entity.Task, error) {
	task := entity.Task{}

	tx, err := t.db.Leader().BeginTx(ctx, "txTask", sql.TxOptions{})
	if err != nil {
		return task, err
	}
	defer tx.Rollback()

	tx, task, err = t.createSQLTask(tx, taskParam)
	if err != nil {
		return task, err
	}

	if err = tx.Commit(); err != nil {
		return task, errors.NewWithCode(codes.CodeSQLTxCommit, err.Error())
	}

	if err := t.deleteTaskCache(ctx); err != nil {
		t.log.Error(ctx, err)
	}

	return t.Get(ctx, entity.TaskParam{
		ID: null.Int64From(task.ID),
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

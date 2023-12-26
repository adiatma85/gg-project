package task

import (
	"context"
	"fmt"
	"time"

	"github.com/adiatma85/new-go-template/src/business/entity"
	"github.com/adiatma85/own-go-sdk/codes"
	"github.com/adiatma85/own-go-sdk/errors"
	"github.com/adiatma85/own-go-sdk/query"
	"github.com/adiatma85/own-go-sdk/redis"
	"github.com/adiatma85/own-go-sdk/sql"
)

func (t *task) createSQLTask(tx sql.CommandTx, v entity.CreateTaskParam) (sql.CommandTx, entity.Task, error) {
	task := entity.Task{}

	res, err := tx.NamedExec("iCreateUser", createTask, v)
	if err != nil {
		return tx, task, errors.NewWithCode(codes.CodeSQLTxExec, err.Error())
	}

	rowCount, err := res.RowsAffected()
	if err != nil || rowCount < 1 {
		return tx, task, errors.NewWithCode(codes.CodeSQLNoRowsAffected, "no rows affected")
	}

	lastID, err := res.LastInsertId()
	if err != nil {
		return tx, task, errors.NewWithCode(codes.CodeSQLNoRowsAffected, err.Error())
	}

	task.ID = lastID

	return tx, task, nil
}

func (t *task) getSQLTask(ctx context.Context, params entity.TaskParam) (entity.Task, error) {
	task := entity.Task{}

	key, err := t.json.Marshal(params)
	if err != nil {
		return task, nil
	}

	cachedTask, err := t.getCache(ctx, fmt.Sprintf(getTaskByIdKey, string(key)))
	switch {
	case errors.Is(err, redis.Nil):
		t.log.Info(ctx, fmt.Sprintf(entity.ErrorRedisNil, err.Error()))
	case err != nil:
		t.log.Error(ctx, fmt.Sprintf(entity.ErrorRedis, err.Error()))
	default:
		return cachedTask, nil
	}

	qb := query.NewSQLQueryBuilder(t.db, "param", "db", &params.QueryOption)
	queryExt, queryArgs, _, _, err := qb.Build(&params)
	if err != nil {
		return task, errors.NewWithCode(codes.CodeSQLBuilder, err.Error())
	}

	row, err := t.db.Follower().QueryRow(ctx, "rTaskByID", getTask+queryExt, queryArgs...)
	if err != nil && errors.Is(err, sql.ErrNotFound) {
		return task, errors.NewWithCode(codes.CodeSQLRead, err.Error())
	} else if errors.Is(err, sql.ErrNotFound) {
		return task, errors.NewWithCode(codes.CodeSQLRecordDoesNotExist, err.Error())
	}

	if err := row.StructScan(&task); err != nil && !errors.Is(err, sql.ErrNotFound) {
		return task, errors.NewWithCode(codes.CodeSQLRowScan, err.Error())
	} else if errors.Is(err, sql.ErrNotFound) {
		return task, errors.NewWithCode(codes.CodeSQLRecordDoesNotExist, err.Error())
	}

	if err = t.upsertCache(ctx, fmt.Sprintf(getTaskByIdKey, string(key)), task, time.Minute); err != nil {
		t.log.Error(ctx, err)
	}

	return task, nil
}

func (t *task) getSQLTaskList(ctx context.Context, params entity.TaskParam) ([]entity.Task, *entity.Pagination, error) {
	tasks := []entity.Task{}

	qb := query.NewSQLQueryBuilder(t.db, "param", "db", &params.QueryOption)
	queryExt, queryArgs, countExt, countArgs, err := qb.Build(&params)
	if err != nil {
		return tasks, nil, errors.NewWithCode(codes.CodeSQLBuilder, err.Error())
	}

	rows, err := t.db.Follower().Query(ctx, "rListTask", getTask+queryExt, queryArgs...)
	if err != nil && !errors.Is(err, sql.ErrNotFound) {
		return tasks, nil, errors.NewWithCode(codes.CodeSQLRead, err.Error())
	}

	defer rows.Close()

	for rows.Next() {
		temp := entity.Task{}
		if err := rows.StructScan(&temp); err != nil {
			t.log.Error(ctx, errors.NewWithCode(codes.CodeSQLRowScan, err.Error()))
			continue
		}
		tasks = append(tasks, temp)
	}

	pg := entity.Pagination{
		CurrentPage:     params.Page,
		CurrentElements: int64(len(tasks)),
	}

	if len(tasks) > 0 && !params.QueryOption.DisableLimit && params.IncludePagination {
		if err := t.db.Follower().Get(ctx, "cTask", readTaskCount+countExt, &pg.TotalElements, countArgs...); err != nil {
			return tasks, nil, errors.NewWithCode(codes.CodeSQLRead, err.Error())
		}
	}

	pg.ProcessPagination(params.Limit)

	return tasks, &pg, nil
}

func (t *task) updateSQLTask(ctx context.Context, updateParam entity.UpdateTaskParam, selectParam entity.TaskParam) error {
	t.log.Debug(ctx, fmt.Sprintf("update task by : %v", selectParam))

	qb := query.NewSQLQueryBuilder(t.db, "param", "db", &selectParam.QueryOption)

	var err error
	queryUpdate, args, err := qb.BuildUpdate(&updateParam, &selectParam)
	if err != nil {
		return errors.NewWithCode(codes.CodeSQLBuilder, err.Error())
	}
	_, err = t.db.Leader().Exec(ctx, "tValue", updateTask+queryUpdate, args...)
	if err != nil {
		return errors.NewWithCode(codes.CodeSQLTxExec, err.Error())
	}

	t.log.Debug(ctx, fmt.Sprintf("successfully update task: %v", updateParam))

	if err := t.deleteTaskCache(ctx); err != nil {
		t.log.Error(ctx, err)
	}

	return nil
}

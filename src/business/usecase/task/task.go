package task

import (
	"context"
	"fmt"
	"time"

	taskDom "github.com/adiatma85/gg-project/src/business/domain/task"
	"github.com/adiatma85/gg-project/src/business/entity"
	"github.com/adiatma85/own-go-sdk/jwtAuth"
	"github.com/adiatma85/own-go-sdk/log"
	"github.com/adiatma85/own-go-sdk/null"
)

type Interface interface {
	Create(ctx context.Context, req entity.CreateTaskParam) (entity.Task, error)
	Get(ctx context.Context, params entity.TaskParam) (entity.Task, error)
	GetList(ctx context.Context, params entity.TaskParam) ([]entity.Task, *entity.Pagination, error)
	Update(ctx context.Context, updateParam entity.UpdateTaskParam, selectParam entity.TaskParam) error
	Delete(ctx context.Context, selectParam entity.TaskParam) error
}

type InitParam struct {
	Log     log.Interface
	Task    taskDom.Interface
	JwtAuth jwtAuth.Interface
}

type task struct {
	log     log.Interface
	task    taskDom.Interface
	jwtAuth jwtAuth.Interface
}

var Now = time.Now

func Init(param InitParam) Interface {
	t := &task{
		log:     param.Log,
		task:    param.Task,
		jwtAuth: param.JwtAuth,
	}

	return t
}

func (t *task) Create(ctx context.Context, req entity.CreateTaskParam) (entity.Task, error) {
	user, err := t.jwtAuth.GetUserAuthInfo(ctx)
	if err != nil {
		return entity.Task{}, err
	}

	req.UserId = user.User.ID
	req.CreatedBy = null.StringFrom(fmt.Sprintf("%v", user.User.ID))

	return t.task.Create(ctx, req)
}

func (t *task) Get(ctx context.Context, params entity.TaskParam) (entity.Task, error) {
	params.QueryOption.IsActive = true
	return t.task.Get(ctx, params)
}

func (t *task) GetList(ctx context.Context, params entity.TaskParam) ([]entity.Task, *entity.Pagination, error) {
	params.IncludePagination = true
	params.QueryOption.IsActive = true

	tasks, pg, err := t.task.GetList(ctx, params)
	if err != nil {
		return nil, nil, err
	}

	return tasks, pg, nil
}

func (t *task) Update(ctx context.Context, updateParam entity.UpdateTaskParam, selectParam entity.TaskParam) error {
	return t.task.Update(ctx, updateParam, selectParam)
}

func (t *task) Delete(ctx context.Context, selectParam entity.TaskParam) error {
	user, err := t.jwtAuth.GetUserAuthInfo(ctx)
	if err != nil {
		return err
	}

	deleteParam := entity.UpdateTaskParam{
		Status:    null.Int64From(-1),
		DeletedAt: null.TimeFrom(Now()),
		DeletedBy: null.StringFrom(fmt.Sprintf("%v", user.User.ID)),
	}

	return t.task.Update(ctx, deleteParam, selectParam)
}

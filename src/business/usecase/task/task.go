package task

import (
	"context"
	"time"

	taskDom "github.com/adiatma85/new-go-template/src/business/domain/task"
	"github.com/adiatma85/new-go-template/src/business/entity"
	"github.com/adiatma85/own-go-sdk/log"
)

type Interface interface {
	Create(ctx context.Context, req entity.CreateTaskParam) (entity.Task, error)
	Get(ctx context.Context, params entity.TaskParam) (entity.Task, error)
	GetList(ctx context.Context, params entity.TaskParam) ([]entity.Task, error)
	Update(ctx context.Context, updateParam entity.UpdateTaskParam, selectParam entity.TaskParam) error
	Delete(ctx context.Context, selectParam entity.TaskParam) error
}

type InitParam struct {
	Log  log.Interface
	Task taskDom.Interface
}

type task struct {
	log  log.Interface
	task taskDom.Interface
}

var Now = time.Now()

func Init(param InitParam) Interface {
	t := &task{
		log:  param.Log,
		task: param.Task,
	}

	return t
}

// Create implements Interface.
func (t *task) Create(ctx context.Context, req entity.CreateTaskParam) (entity.Task, error) {
	// var result entity.Task

	// result, err := t.validateTask(ctx, req)
	// if err != nil {
	// 	return result, err
	// }

	// req.CreatedBy = null.StringFrom(fmt.Sprintf("%v", entity.SystemUser))
	panic("unimplemented")
}

// Delete implements Interface.
func (*task) Delete(ctx context.Context, selectParam entity.TaskParam) error {
	panic("unimplemented")
}

// Get implements Interface.
func (*task) Get(ctx context.Context, params entity.TaskParam) (entity.Task, error) {
	panic("unimplemented")
}

// GetList implements Interface.
func (*task) GetList(ctx context.Context, params entity.TaskParam) ([]entity.Task, error) {
	panic("unimplemented")
}

// Update implements Interface.
func (*task) Update(ctx context.Context, updateParam entity.UpdateTaskParam, selectParam entity.TaskParam) error {
	panic("unimplemented")
}

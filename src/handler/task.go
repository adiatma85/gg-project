package handler

import (
	"github.com/adiatma85/gg-project/src/business/entity"
	"github.com/adiatma85/own-go-sdk/codes"
	"github.com/gin-gonic/gin"
)

// @Summary Create Task
// @Description Create new entry for Task
// @Security BearerAuth
// @Tags Task
// @Param data body entity.CreateTaskParam true "Input New Task Data"
// @Produce json
// @Success 200 {object} entity.HTTPResp{data=entity.Task{}}
// @Failure 400 {object} entity.HTTPResp{}
// @Failure 500 {object} entity.HTTPResp{}
// @Router /v1/task [post]
func (r *rest) CreateTask(ctx *gin.Context) {
	var param entity.CreateTaskParam
	if err := r.Bind(ctx, &param); err != nil {
		r.httpRespError(ctx, err)
		return
	}

	task, err := r.uc.Task.Create(ctx.Request.Context(), param)
	if err != nil {
		r.httpRespError(ctx, err)
		return
	}

	r.httpRespSuccess(ctx, codes.CodeSuccess, task, nil)
}

// @Summary Get Task List
// @Description Get list all Task
// @Security BearerAuth
// @Tags Task
// @Param limit query integer false "limit"
// @Param page query integer false "page"
// @Param disableLimit query boolean false "disable limit" Enums(true, false)
// @Produce json
// @Success 200 {object} entity.HTTPResp{data=[]entity.Task{}}
// @Failure 500 {object} entity.HTTPResp{}
// @Router /v1/task [GET]
func (r *rest) GetListTask(ctx *gin.Context) {
	var param entity.TaskParam
	if err := r.BindParams(ctx, &param); err != nil {
		r.httpRespError(ctx, err)
		return
	}

	tasks, pg, err := r.uc.Task.GetList(ctx.Request.Context(), param)
	if err != nil {
		r.httpRespError(ctx, err)
		return
	}

	r.httpRespSuccess(ctx, codes.CodeSuccess, tasks, pg)
}

// @Summary Get Task List by User Id
// @Description Get list all Task by User Id
// @Security BearerAuth
// @Tags Task
// @Param user_id path integer true "user id"
// @Param limit query integer false "limit"
// @Param page query integer false "page"
// @Param disableLimit query boolean false "disable limit" Enums(true, false)
// @Produce json
// @Success 200 {object} entity.HTTPResp{data=[]entity.Task{}}
// @Failure 500 {object} entity.HTTPResp{}
// @Router /v1/user/{user_id}/task [GET]
func (r *rest) GetListTaskWithUserId(ctx *gin.Context) {
	var param entity.TaskParam
	if err := r.BindParams(ctx, &param); err != nil {
		r.httpRespError(ctx, err)
		return
	}

	tasks, pg, err := r.uc.Task.GetList(ctx.Request.Context(), param)
	if err != nil {
		r.httpRespError(ctx, err)
		return
	}

	r.httpRespSuccess(ctx, codes.CodeSuccess, tasks, pg)
}

// @Summary Get Task By ID
// @Description Get Task details by Task ID
// @Security BearerAuth
// @Tags Task
// @Param task_id path integer true "Task id"
// @Produce json
// @Success 200 {object} entity.HTTPResp{data=entity.Task{}}
// @Failure 500 {object} entity.HTTPResp{}
// @Router /v1/task/{task_id} [GET]
func (r *rest) GetTaskById(ctx *gin.Context) {
	var param entity.TaskParam
	if err := r.BindParams(ctx, &param); err != nil {
		r.httpRespError(ctx, err)
		return
	}

	task, err := r.uc.Task.Get(ctx.Request.Context(), param)
	if err != nil {
		r.httpRespError(ctx, err)
		return
	}

	r.httpRespSuccess(ctx, codes.CodeSuccess, task, nil)
}

// @Summary Update One Task
// @Description Update one Task detail
// @Security BearerAuth
// @Tags Task
// @Param task_id path integer true "Task id"
// @Param task body entity.UpdateTaskParam true "Task data"
// @Produce json
// @Success 200 {object} entity.HTTPResp{}
// @Failure 500 {object} entity.HTTPResp{}
// @Router /v1/task/{task_id} [PUT]
func (r *rest) UpdateTask(ctx *gin.Context) {
	var updateParam entity.UpdateTaskParam
	if err := r.Bind(ctx, &updateParam); err != nil {
		r.httpRespError(ctx, err)
		return
	}

	var selectParam entity.TaskParam
	if err := r.BindParams(ctx, &selectParam); err != nil {
		r.httpRespError(ctx, err)
		return
	}

	if err := r.uc.Task.Update(ctx.Request.Context(), updateParam, selectParam); err != nil {
		r.httpRespError(ctx, err)
		return
	}

	r.httpRespSuccess(ctx, codes.CodeSuccess, nil, nil)
}

// @Summary Delete Task
// @Description Soft delete Task data
// @Security BearerAuth
// @Tags Task
// @Param task_id path integer true "task id"
// @Produce json
// @Success 200 {object} entity.HTTPResp{}
// @Failure 500 {object} entity.HTTPResp{}
// @Router /v1/task/{task_id} [DELETE]
func (r *rest) DeleteTask(ctx *gin.Context) {
	var selectParam entity.TaskParam
	if err := r.BindParams(ctx, &selectParam); err != nil {
		r.httpRespError(ctx, err)
		return
	}

	if err := r.uc.Task.Delete(ctx.Request.Context(), selectParam); err != nil {
		r.httpRespError(ctx, err)
		return
	}

	r.httpRespSuccess(ctx, codes.CodeSuccess, nil, nil)
}

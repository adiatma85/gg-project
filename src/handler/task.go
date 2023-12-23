package handler

import (
	"github.com/adiatma85/new-go-template/src/business/entity"
	"github.com/gin-gonic/gin"
)

func (r *rest) getTaskById(ctx *gin.Context) {
	var param entity.TaskParam

	if err := r.BindParams(ctx, &param); err != nil {
		r.httpRespError(ctx, err)
	}
	// task := r
}

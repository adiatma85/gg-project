package handler

import (
	"github.com/adiatma85/gg-project/src/business/entity"
	"github.com/adiatma85/own-go-sdk/codes"
	"github.com/gin-gonic/gin"
)

// @Summary Get User List
// @Description Get list all user
// @Security BearerAuth
// @Tags Admin
// @Param limit query integer false "limit"
// @Param page query integer false "page"
// @Param disableLimit query boolean false "disable limit" Enums(true, false)
// @Produce json
// @Success 200 {object} entity.HTTPResp{data=[]entity.User{}}
// @Failure 500 {object} entity.HTTPResp{}
// @Router /v1/admin/user [GET]
func (r *rest) GetListUser(ctx *gin.Context) {
	var param entity.UserParam
	if err := r.BindParams(ctx, &param); err != nil {
		r.httpRespError(ctx, err)
		return
	}

	user, pg, err := r.uc.User.GetListAsAdmin(ctx.Request.Context(), param)
	if err != nil {
		r.httpRespError(ctx, err)
		return
	}

	r.httpRespSuccess(ctx, codes.CodeSuccess, user, pg)
}

// @Summary Get User By ID
// @Description Get user details by user ID
// @Security BearerAuth
// @Tags User
// @Param user_id path integer true "user id"
// @Produce json
// @Success 200 {object} entity.HTTPResp{data=entity.User{}}
// @Failure 500 {object} entity.HTTPResp{}
// @Router /v1/user/{user_id} [GET]
func (r *rest) GetUserByID(ctx *gin.Context) {
	var param entity.UserParam
	if err := r.BindParams(ctx, &param); err != nil {
		r.httpRespError(ctx, err)
		return
	}

	user, err := r.uc.User.Get(ctx.Request.Context(), param)
	if err != nil {
		r.httpRespError(ctx, err)
		return
	}

	r.httpRespSuccess(ctx, codes.CodeSuccess, user, nil)
}

// @Summary Update One User
// @Description Update one user detail
// @Security BearerAuth
// @Tags User
// @Param user_id path integer true "user id"
// @Param user body entity.UpdateUserParam true "user data"
// @Produce json
// @Success 200 {object} entity.HTTPResp{data=entity.User{}}
// @Failure 500 {object} entity.HTTPResp{}
// @Router /v1/user/{user_id} [PUT]
func (r *rest) UpdateUser(ctx *gin.Context) {
	var updateParam entity.UpdateUserParam

	if err := r.Bind(ctx, &updateParam); err != nil {
		r.httpRespError(ctx, err)
		return
	}

	var selectParam entity.UserParam
	if err := r.BindParams(ctx, &selectParam); err != nil {
		r.httpRespError(ctx, err)
		return
	}

	if err := r.uc.User.Update(ctx.Request.Context(), updateParam, selectParam); err != nil {
		r.httpRespError(ctx, err)
		return
	}

	r.httpRespSuccess(ctx, codes.CodeSuccess, nil, nil)
}

// @Summary Delete User
// @Description Soft delete user data
// @Security BearerAuth
// @Tags User
// @Param user_id path integer true "user id"
// @Produce json
// @Success 200 {object} entity.HTTPResp{data=entity.User{}}
// @Failure 500 {object} entity.HTTPResp{}
// @Router /v1/user/{user_id} [DELETE]
func (r *rest) DeleteUser(ctx *gin.Context) {
	var selectParam entity.UserParam

	if err := r.BindParams(ctx, &selectParam); err != nil {
		r.httpRespError(ctx, err)
		return
	}

	if err := r.uc.User.Delete(ctx.Request.Context(), selectParam); err != nil {
		r.httpRespError(ctx, err)
		return
	}

	r.httpRespSuccess(ctx, codes.CodeSuccess, nil, nil)
}

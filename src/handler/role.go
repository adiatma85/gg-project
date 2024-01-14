package handler

import (
	"github.com/adiatma85/gg-project/src/business/entity"
	"github.com/adiatma85/own-go-sdk/codes"
	"github.com/gin-gonic/gin"
)

// @Summary Create Role
// @Description Create new entry for Role
// @Security BearerAuth
// @Tags Role
// @Param data body entity.CreateRoleParam true "Input New Role Data"
// @Produce json
// @Success 200 {object} entity.HTTPResp{data=entity.Role{}}
// @Failure 400 {object} entity.HTTPResp{}
// @Failure 500 {object} entity.HTTPResp{}
// @Router /v1/role [post]
func (r *rest) CreateRole(ctx *gin.Context) {
	var param entity.CreateRoleParam
	if err := r.Bind(ctx, &param); err != nil {
		r.httpRespError(ctx, err)
		return
	}

	role, err := r.uc.Role.Create(ctx.Request.Context(), param)
	if err != nil {
		r.httpRespError(ctx, err)
		return
	}

	r.httpRespSuccess(ctx, codes.CodeSuccess, role, nil)
}

// @Summary Get Role List
// @Description Get list all Role
// @Security BearerAuth
// @Tags Role
// @Param limit query integer false "limit"
// @Param page query integer false "page"
// @Param disableLimit query boolean false "disable limit" Enums(true, false)
// @Produce json
// @Success 200 {object} entity.HTTPResp{data=[]entity.Role{}}
// @Failure 500 {object} entity.HTTPResp{}
// @Router /v1/role [GET]
func (r *rest) GetListRole(ctx *gin.Context) {
	var param entity.RoleParam
	if err := r.BindParams(ctx, &param); err != nil {
		r.httpRespError(ctx, err)
		return
	}

	roles, pg, err := r.uc.Role.GetList(ctx.Request.Context(), param)
	if err != nil {
		r.httpRespError(ctx, err)
		return
	}

	r.httpRespSuccess(ctx, codes.CodeSuccess, roles, pg)
}

// @Summary Get Role By ID
// @Description Get Role details by Role ID
// @Security BearerAuth
// @Tags Role
// @Param role_id path integer true "Role id"
// @Produce json
// @Success 200 {object} entity.HTTPResp{data=entity.Role{}}
// @Failure 500 {object} entity.HTTPResp{}
// @Router /v1/role/{role_id} [GET]
func (r *rest) GetRoleById(ctx *gin.Context) {
	var param entity.RoleParam
	if err := r.BindParams(ctx, &param); err != nil {
		r.httpRespError(ctx, err)
		return
	}

	role, err := r.uc.Role.Get(ctx.Request.Context(), param)
	if err != nil {
		r.httpRespError(ctx, err)
		return
	}

	r.httpRespSuccess(ctx, codes.CodeSuccess, role, nil)
}

// @Summary Update One Role
// @Description Update one Role detail
// @Security BearerAuth
// @Tags Role
// @Param role_id path integer true "Role id"
// @Param role body entity.UpdateRoleParam true "Role data"
// @Produce json
// @Success 200 {object} entity.HTTPResp{}
// @Failure 500 {object} entity.HTTPResp{}
// @Router /v1/role/{role_id} [PUT]
func (r *rest) UpdateRole(ctx *gin.Context) {
	var updateParam entity.UpdateRoleParam
	if err := r.Bind(ctx, &updateParam); err != nil {
		r.httpRespError(ctx, err)
		return
	}

	var selectParam entity.RoleParam
	if err := r.BindParams(ctx, &selectParam); err != nil {
		r.httpRespError(ctx, err)
		return
	}

	if err := r.uc.Role.Update(ctx.Request.Context(), updateParam, selectParam); err != nil {
		r.httpRespError(ctx, err)
		return
	}

	r.httpRespSuccess(ctx, codes.CodeSuccess, nil, nil)
}

// @Summary Delete Role
// @Description Soft delete Role data
// @Security BearerAuth
// @Tags Role
// @Param role_id path integer true "role id"
// @Produce json
// @Success 200 {object} entity.HTTPResp{}
// @Failure 500 {object} entity.HTTPResp{}
// @Router /v1/role/{role_id} [DELETE]
func (r *rest) DeleteRole(ctx *gin.Context) {
	var selectParam entity.RoleParam
	if err := r.BindParams(ctx, &selectParam); err != nil {
		r.httpRespError(ctx, err)
		return
	}

	if err := r.uc.Role.Delete(ctx.Request.Context(), selectParam); err != nil {
		r.httpRespError(ctx, err)
		return
	}

	r.httpRespSuccess(ctx, codes.CodeSuccess, nil, nil)
}

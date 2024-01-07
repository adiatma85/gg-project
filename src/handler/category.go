package handler

import (
	"github.com/adiatma85/gg-project/src/business/entity"
	"github.com/adiatma85/own-go-sdk/codes"
	"github.com/gin-gonic/gin"
)

// @Summary Create Category
// @Description Create new entry for Category
// @Security BearerAuth
// @Tags Category
// @Param data body entity.CreateCategoryParam true "Input New Category Data"
// @Produce json
// @Success 200 {object} entity.HTTPResp{data=entity.Category{}}
// @Failure 400 {object} entity.HTTPResp{}
// @Failure 500 {object} entity.HTTPResp{}
// @Router /v1/category [post]
func (r *rest) CreateCategory(ctx *gin.Context) {
	var param entity.CreateCategoryParam
	if err := r.Bind(ctx, &param); err != nil {
		r.httpRespError(ctx, err)
		return
	}

	category, err := r.uc.Category.Create(ctx.Request.Context(), param)
	if err != nil {
		r.httpRespError(ctx, err)
		return
	}

	r.httpRespSuccess(ctx, codes.CodeSuccess, category, nil)
}

// @Summary Get Category List
// @Description Get list all Category
// @Security BearerAuth
// @Tags Category
// @Param limit query integer false "limit"
// @Param page query integer false "page"
// @Param disableLimit query boolean false "disable limit" Enums(true, false)
// @Produce json
// @Success 200 {object} entity.HTTPResp{data=[]entity.Category{}}
// @Failure 500 {object} entity.HTTPResp{}
// @Router /v1/category [GET]
func (r *rest) GetListCategory(ctx *gin.Context) {
	var param entity.CategoryParam
	if err := r.BindParams(ctx, &param); err != nil {
		r.httpRespError(ctx, err)
		return
	}

	categories, pg, err := r.uc.Category.GetList(ctx.Request.Context(), param)
	if err != nil {
		r.httpRespError(ctx, err)
		return
	}

	r.httpRespSuccess(ctx, codes.CodeSuccess, categories, pg)
}

// @Summary Get Category By ID
// @Description Get Category details by Category ID
// @Security BearerAuth
// @Tags Category
// @Param category_id path integer true "Category id"
// @Produce json
// @Success 200 {object} entity.HTTPResp{data=entity.Category{}}
// @Failure 500 {object} entity.HTTPResp{}
// @Router /v1/category/{category_id} [GET]
func (r *rest) GetCategoryByID(ctx *gin.Context) {
	var param entity.CategoryParam
	if err := r.BindParams(ctx, &param); err != nil {
		r.httpRespError(ctx, err)
		return
	}

	category, err := r.uc.Category.Get(ctx.Request.Context(), param)
	if err != nil {
		r.httpRespError(ctx, err)
		return
	}

	r.httpRespSuccess(ctx, codes.CodeSuccess, category, nil)
}

// @Summary Update One Category
// @Description Update one category detail
// @Security BearerAuth
// @Tags Category
// @Param category_id path integer true "Category id"
// @Param category body entity.UpdateCategoryParam true "category data"
// @Produce json
// @Success 200 {object} entity.HTTPResp{data=entity.Category{}}
// @Failure 500 {object} entity.HTTPResp{}
// @Router /v1/category/{category_id} [PUT]
func (r *rest) UpdateCategory(ctx *gin.Context) {
	var updateParam entity.UpdateCategoryParam
	if err := r.Bind(ctx, &updateParam); err != nil {
		r.httpRespError(ctx, err)
		return
	}

	var selectParam entity.CategoryParam
	if err := r.BindParams(ctx, &selectParam); err != nil {
		r.httpRespError(ctx, err)
		return
	}

	if err := r.uc.Category.Update(ctx.Request.Context(), updateParam, selectParam); err != nil {
		r.httpRespError(ctx, err)
		return
	}

	r.httpRespSuccess(ctx, codes.CodeSuccess, nil, nil)
}

// @Summary Delete Category
// @Description Soft delete category data
// @Security BearerAuth
// @Tags Category
// @Param category_id path integer true "category id"
// @Produce json
// @Success 200 {object} entity.HTTPResp{data=entity.Category{}}
// @Failure 500 {object} entity.HTTPResp{}
// @Router /v1/category/{category_id} [DELETE]
func (r *rest) DeleteCategory(ctx *gin.Context) {
	var selectParam entity.CategoryParam
	if err := r.BindParams(ctx, &selectParam); err != nil {
		r.httpRespError(ctx, err)
		return
	}

	if err := r.uc.Category.Delete(ctx.Request.Context(), selectParam); err != nil {
		r.httpRespError(ctx, err)
		return
	}

	r.httpRespSuccess(ctx, codes.CodeSuccess, nil, nil)
}

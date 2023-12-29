package category

import (
	"context"
	"fmt"
	"time"

	categoryDom "github.com/adiatma85/gg-project/src/business/domain/category"
	"github.com/adiatma85/gg-project/src/business/entity"
	"github.com/adiatma85/own-go-sdk/jwtAuth"
	"github.com/adiatma85/own-go-sdk/log"
	"github.com/adiatma85/own-go-sdk/null"
)

type Interface interface {
	Create(ctx context.Context, req entity.CreateCategoryParam) (entity.Category, error)
	Get(ctx context.Context, params entity.CategoryParam) (entity.Category, error)
	GetList(ctx context.Context, params entity.CategoryParam) ([]entity.Category, *entity.Pagination, error)
	Update(ctx context.Context, updateParam entity.UpdateCategoryParam, selectParam entity.CategoryParam) error
	Delete(ctx context.Context, selectParam entity.CategoryParam) error
}

type InitParam struct {
	Log      log.Interface
	Category categoryDom.Interface
	JwtAuth  jwtAuth.Interface
}

type category struct {
	log      log.Interface
	category categoryDom.Interface
	jwtAuth  jwtAuth.Interface
}

var Now = time.Now

func Init(param InitParam) Interface {
	c := &category{
		log:      param.Log,
		category: param.Category,
		jwtAuth:  param.JwtAuth,
	}

	return c
}

func (c *category) Create(ctx context.Context, req entity.CreateCategoryParam) (entity.Category, error) {
	user, err := c.jwtAuth.GetUserAuthInfo(ctx)
	if err != nil {
		return entity.Category{}, err
	}

	req.CreatedBy = null.StringFrom(fmt.Sprintf("%v", user.User.ID))

	return c.category.Create(ctx, req)
}

func (c *category) Get(ctx context.Context, params entity.CategoryParam) (entity.Category, error) {
	return c.category.Get(ctx, params)
}

func (c *category) GetList(ctx context.Context, params entity.CategoryParam) ([]entity.Category, *entity.Pagination, error) {
	params.IncludePagination = true
	params.QueryOption.IsActive = true

	categories, pg, err := c.category.GetList(ctx, params)
	if err != nil {
		return nil, nil, err
	}

	return categories, pg, nil
}

func (c *category) Update(ctx context.Context, updateParam entity.UpdateCategoryParam, selectParam entity.CategoryParam) error {
	return c.category.Update(ctx, updateParam, selectParam)
}

func (c *category) Delete(ctx context.Context, selectParam entity.CategoryParam) error {
	user, err := c.jwtAuth.GetUserAuthInfo(ctx)
	if err != nil {
		return err
	}

	deleteParam := entity.UpdateCategoryParam{
		Status:    null.Int64From(-1),
		DeletedAt: null.TimeFrom(Now()),
		DeletedBy: null.StringFrom(fmt.Sprintf("%v", user.User.ID)),
	}

	return c.category.Update(ctx, deleteParam, selectParam)
}

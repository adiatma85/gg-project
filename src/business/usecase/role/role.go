package role

import (
	"context"
	"fmt"
	"time"

	roleDom "github.com/adiatma85/gg-project/src/business/domain/role"
	"github.com/adiatma85/gg-project/src/business/entity"
	"github.com/adiatma85/own-go-sdk/jwtAuth"
	"github.com/adiatma85/own-go-sdk/log"
	"github.com/adiatma85/own-go-sdk/null"
)

type Interface interface {
	Create(ctx context.Context, req entity.CreateRoleParam) (entity.Role, error)
	Get(ctx context.Context, params entity.RoleParam) (entity.Role, error)
	GetList(ctx context.Context, params entity.RoleParam) ([]entity.Role, *entity.Pagination, error)
	Update(ctx context.Context, updateParam entity.UpdateRoleParam, selectParam entity.RoleParam) error
	Delete(ctx context.Context, selectParam entity.RoleParam) error
}

type InitParam struct {
	Log     log.Interface
	Role    roleDom.Interface
	JwtAuth jwtAuth.Interface
}

type role struct {
	log     log.Interface
	role    roleDom.Interface
	jwtAuth jwtAuth.Interface
}

var Now = time.Now

func Init(param InitParam) Interface {
	r := &role{}

	return r
}

func (r *role) Create(ctx context.Context, req entity.CreateRoleParam) (entity.Role, error) {
	user, err := r.jwtAuth.GetUserAuthInfo(ctx)
	if err != nil {
		return entity.Role{}, err
	}

	req.CreatedBy = null.StringFrom(fmt.Sprintf("%v", user.User.ID))
	req.UpdatedBy = null.StringFrom(fmt.Sprintf("%v", user.User.ID))

	return r.role.Create(ctx, req)
}

func (r *role) Get(ctx context.Context, params entity.RoleParam) (entity.Role, error) {
	params.QueryOption.IsActive = true
	return r.role.Get(ctx, params)
}

func (r *role) GetList(ctx context.Context, params entity.RoleParam) ([]entity.Role, *entity.Pagination, error) {
	params.IncludePagination = true
	params.QueryOption.IsActive = true

	roles, pg, err := r.role.GetList(ctx, params)
	if err != nil {
		return nil, nil, err
	}

	return roles, pg, nil
}

func (r *role) Update(ctx context.Context, updateParam entity.UpdateRoleParam, selectParam entity.RoleParam) error {
	return r.role.Update(ctx, updateParam, selectParam)
}

func (r *role) Delete(ctx context.Context, selectParam entity.RoleParam) error {
	user, err := r.jwtAuth.GetUserAuthInfo(ctx)
	if err != nil {
		return err
	}

	deleteParam := entity.UpdateRoleParam{
		Status:    null.Int64From(-1),
		DeletedAt: null.TimeFrom(Now()),
		DeletedBy: null.StringFrom(fmt.Sprintf("%v", user.User.ID)),
	}

	return r.role.Update(ctx, deleteParam, selectParam)
}

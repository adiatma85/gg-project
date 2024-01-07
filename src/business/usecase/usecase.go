package usecase

import (
	"github.com/adiatma85/gg-project/src/business/domain"
	"github.com/adiatma85/gg-project/src/business/usecase/category"
	"github.com/adiatma85/gg-project/src/business/usecase/task"
	"github.com/adiatma85/gg-project/src/business/usecase/user"
	"github.com/adiatma85/own-go-sdk/jwtAuth"
	"github.com/adiatma85/own-go-sdk/log"
)

type Usecase struct {
	User     user.Interface
	Category category.Interface
	Task     task.Interface
}

type InitParam struct {
	Log     log.Interface
	Dom     *domain.Domain
	JwtAuth jwtAuth.Interface
}

func Init(param InitParam) *Usecase {
	usecase := &Usecase{
		User:     user.Init(user.InitParam{Log: param.Log, User: param.Dom.User, JwtAuth: param.JwtAuth}),
		Category: category.Init(category.InitParam{Log: param.Log, Category: param.Dom.Category, JwtAuth: param.JwtAuth}),
		Task:     task.Init(task.InitParam{Log: param.Log, Task: param.Dom.Task, JwtAuth: param.JwtAuth}),
	}

	return usecase
}

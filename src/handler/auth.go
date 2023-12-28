package handler

import (
	"github.com/adiatma85/gg-project/src/business/entity"
	"github.com/adiatma85/own-go-sdk/codes"
	"github.com/gin-gonic/gin"
)

// @Summary Register New User
// @Description Register new user
// @Tags Auth
// @Param data body entity.CreateUserParam true "Input New User Data"
// @Produce json
// @Success 200 {object} entity.HTTPResp{data=entity.User{}}
// @Failure 400 {object} entity.HTTPResp{}
// @Failure 500 {object} entity.HTTPResp{}
// @Router /public/v1/register [post]
func (r *rest) RegisterNewUserWithoutToken(ctx *gin.Context) {
	var param entity.CreateUserParam
	if err := r.Bind(ctx, &param); err != nil {
		r.httpRespError(ctx, err)
		return
	}

	authInfo, err := r.uc.User.CreateWithoutAuthInfo(ctx.Request.Context(), param)
	if err != nil {
		r.httpRespError(ctx, err)
		return
	}

	r.httpRespSuccess(ctx, codes.CodeSuccess, authInfo, nil)
}

// @Summary Sign In With Password
// @Description This endpoint will sign in user with email and password
// @Tags Auth
// @Param data body entity.UserLoginRequest true "Input Email And Password"
// @Produce json
// @Success 200 {object} entity.HTTPResp{}
// @Failure 400 {object} entity.HTTPResp{}
// @Failure 404 {object} entity.HTTPResp{}
// @Failure 500 {object} entity.HTTPResp{}
// @Router /auth/v1/login [post]
func (r *rest) SignInWithPassword(ctx *gin.Context) {
	var inputUserLoginData entity.UserLoginRequest
	err := r.Bind(ctx, &inputUserLoginData)
	if err != nil {
		r.httpRespError(ctx, err)
		return
	}

	authInfo, err := r.uc.User.SignInWithPassword(ctx.Request.Context(), inputUserLoginData)
	if err != nil {
		r.httpRespError(ctx, err)
		return
	}

	r.httpRespSuccess(ctx, codes.CodeSuccess, authInfo, nil)
}

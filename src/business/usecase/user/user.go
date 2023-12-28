package user

import (
	"context"
	"fmt"
	"time"

	userDom "github.com/adiatma85/gg-project/src/business/domain/user"
	"github.com/adiatma85/gg-project/src/business/entity"
	"github.com/adiatma85/own-go-sdk/codes"
	"github.com/adiatma85/own-go-sdk/errors"
	"github.com/adiatma85/own-go-sdk/jwtAuth"
	"github.com/adiatma85/own-go-sdk/log"
	"github.com/adiatma85/own-go-sdk/null"
	"github.com/adiatma85/own-go-sdk/query"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type Interface interface {
	CreateWithoutAuthInfo(ctx context.Context, params entity.CreateUserParam) (entity.User, error)
	Create(ctx context.Context, req entity.CreateUserParam) (entity.User, error)
	Get(ctx context.Context, params entity.UserParam) (entity.User, error)
	GetListAsAdmin(ctx context.Context, params entity.UserParam) ([]entity.User, *entity.Pagination, error)
	Update(ctx context.Context, updateParam entity.UpdateUserParam, selectParam entity.UserParam) error
	Delete(ctx context.Context, selectParam entity.UserParam) error
	SignInWithPassword(ctx context.Context, req entity.UserLoginRequest) (entity.UserLoginResponse, error)

	// Improvement kedepannya
	// CheckPassword(ctx context.Context, params entity.UserCheckPasswordParam, userParam entity.UserParam) (entity.HTTPMessage, error)
	// ChangePassword(ctx context.Context, passwordChangeParam entity.UserChangePasswordParam, userParam entity.UserParam) (entity.HTTPMessage, error)
	// Activate(ctx context.Context, selectParam entity.UserParam) error
	// SignInWithPassword(ctx context.Context, req entity.UserLoginRequest) (entity.UserLoginResponse, error)
	// RefreshToken(ctx context.Context, param entity.UserRefreshTokenParam) (entity.RefreshTokenResponse, error)
}

type InitParam struct {
	Log     log.Interface
	User    userDom.Interface
	JwtAuth jwtAuth.Interface
}

type user struct {
	log     log.Interface
	user    userDom.Interface
	jwtAuth jwtAuth.Interface
}

var Now = time.Now

func Init(param InitParam) Interface {
	u := &user{
		log:     param.Log,
		user:    param.User,
		jwtAuth: param.JwtAuth,
	}

	return u
}

func (u *user) Create(ctx context.Context, req entity.CreateUserParam) (entity.User, error) {
	var result entity.User

	result, err := u.validateUser(ctx, req)
	if err != nil {
		return result, err
	}

	req.CreatedBy = null.StringFrom(fmt.Sprintf("%v", entity.SystemUser))

	return u.user.Create(ctx, req)
}

func (u *user) CreateWithoutAuthInfo(ctx context.Context, req entity.CreateUserParam) (entity.User, error) {
	var result entity.User
	req.ConfirmPassword = req.Password

	result, err := u.validateUser(ctx, req)
	if err != nil {
		return result, err
	}

	// Hash the password in here
	req.Password, err = u.getHashPassowrd(req.Password)

	return u.user.Create(ctx, req)
}

func (u *user) validateUser(ctx context.Context, req entity.CreateUserParam) (entity.User, error) {
	var result entity.User

	if req.Password != req.ConfirmPassword {
		return result, errors.NewWithCode(codes.CodePasswordDoesNotMatch, "password does not match")
	}

	user, err := u.user.Get(ctx, entity.UserParam{
		Email: null.StringFrom(req.Email),
	})
	if err != nil && errors.GetCode(err) != codes.CodeSQLRecordDoesNotExist {
		return result, err
	}

	if user != result {
		return result, errors.NewWithCode(codes.CodeConflict, "email is exists")
	}

	return result, nil
}

func (u *user) Get(ctx context.Context, params entity.UserParam) (entity.User, error) {
	return u.user.Get(ctx, params)
}

func (u *user) GetListAsAdmin(ctx context.Context, params entity.UserParam) ([]entity.User, *entity.Pagination, error) {
	params.IncludePagination = true
	users, pg, err := u.user.GetList(ctx, params)
	if err != nil {
		return nil, nil, err
	}

	return users, pg, nil
}

func (u *user) Update(ctx context.Context, updateParam entity.UpdateUserParam, selectParam entity.UserParam) error {
	return u.user.Update(ctx, updateParam, selectParam)
}

func (u *user) Delete(ctx context.Context, selectParam entity.UserParam) error {
	deleteParam := entity.UpdateUserParam{
		Status:    null.Int64From(-1),
		DeletedAt: null.TimeFrom(Now()),
		DeletedBy: null.StringFrom(fmt.Sprintf("%v", entity.SystemUser)),
	}

	return u.user.Update(ctx, deleteParam, selectParam)
}

func (u *user) getHashPassowrd(password string) (string, error) {
	bytePassword := []byte(password)
	hash, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func (u *user) checkHashPassword(hashPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
	return err == nil
}

func (u *user) SignInWithPassword(ctx context.Context, req entity.UserLoginRequest) (entity.UserLoginResponse, error) {
	// validate body request
	if err := u.validateUserLoginRequest(req); err != nil {
		return entity.UserLoginResponse{}, err
	}

	// validate user is exist on db and status is active
	user, err := u.user.Get(ctx, entity.UserParam{
		Email: null.StringFrom(req.Email),
		QueryOption: query.Option{
			IsActive: true,
		},
	})
	if err != nil {
		if errors.GetCode(err) == codes.CodeSQLRecordDoesNotExist {
			return entity.UserLoginResponse{}, errors.NewWithCode(codes.CodeNotFound, "email not found")
		}
		return entity.UserLoginResponse{}, err
	}

	// Create the JWT token in here
	accessToken, err := u.jwtAuth.CreateAccessToken(user.ConvertToAuthUser())
	if err != nil {
		return entity.UserLoginResponse{}, err
	}

	result := entity.UserLoginResponse{
		Email:       user.Email,
		DisplayName: user.DisplayName,
		AccessToken: accessToken,
	}

	return result, nil
}

func (u *user) validateUserLoginRequest(req entity.UserLoginRequest) error {
	if req.Email == "" {
		return errors.NewWithCode(codes.CodeBadRequest, "email is required")
	}

	if req.Password == "" {
		return errors.NewWithCode(codes.CodeBadRequest, "password is required")
	}

	return nil
}

func (u *user) tempCreateToken(user jwtAuth.User) (string, error) {
	expirationTime := time.Now().Add(5 * time.Hour)

	claims := &jwtAuth.Claim{
		UserID: user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("JUANCOK"))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

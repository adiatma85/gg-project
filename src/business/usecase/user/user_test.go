package user

import (
	"context"
	"reflect"
	"testing"

	mock_user_dom "github.com/adiatma85/gg-project/src/business/domain/mock/user"
	"github.com/adiatma85/gg-project/src/business/entity"
	"github.com/adiatma85/own-go-sdk/null"
	mock_jwt_auth "github.com/adiatma85/own-go-sdk/tests/mock/jwtAuth"
	mock_log "github.com/adiatma85/own-go-sdk/tests/mock/log"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

type mockInterface struct {
	logger  *mock_log.MockInterface
	userDom *mock_user_dom.MockInterface
	jwtAuth *mock_jwt_auth.MockInterface
}

func initMockTest(t *testing.T) (Interface, user, mockInterface) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	logger := mock_log.NewMockInterface(ctrl)
	logger.EXPECT().Info(gomock.Any(), gomock.Any()).AnyTimes()
	logger.EXPECT().Error(gomock.Any(), gomock.Any()).AnyTimes()
	logger.EXPECT().Debug(gomock.Any(), gomock.Any()).AnyTimes()

	mockUserDom := mock_user_dom.NewMockInterface(ctrl)
	mockJwtAuth := mock_jwt_auth.NewMockInterface(ctrl)

	ucInterface := Init(InitParam{
		Log:     logger,
		User:    mockUserDom,
		JwtAuth: mockJwtAuth,
	})

	ucStruct := user{
		log:     logger,
		user:    mockUserDom,
		jwtAuth: mockJwtAuth,
	}

	mockInterface := mockInterface{
		logger:  logger,
		userDom: mockUserDom,
		jwtAuth: mockJwtAuth,
	}

	return ucInterface, ucStruct, mockInterface
}

func Test_user_Create(t *testing.T) {
	usecase, _, mocks := initMockTest(t)

	// Type in here
	type args struct {
		ctx         context.Context
		insertParam entity.CreateUserParam
	}

	// Mock in here
	mockInsertParam := entity.CreateUserParam{
		Email: "random@email.com",
	}

	mockUserFetchParam := entity.UserParam{
		Email: null.StringFrom("random@email.com"),
	}

	// Test cases in here
	tests := []struct {
		name     string
		arg      args
		mockFunc func(mock mockInterface, arg args)
		want     entity.User
		wantErr  bool
	}{
		{
			name: "password does not match",
			arg: args{
				ctx: context.Background(),
				insertParam: entity.CreateUserParam{
					Password:        "password 1",
					ConfirmPassword: "password 2",
				},
			},
			mockFunc: func(mock mockInterface, arg args) {},
			want:     entity.User{},
			wantErr:  true,
		},
		{
			name: "failed to fetch from user domain",
			arg: args{
				ctx:         context.Background(),
				insertParam: mockInsertParam,
			},
			mockFunc: func(mock mockInterface, arg args) {
				mock.userDom.EXPECT().Get(arg.ctx, mockUserFetchParam).Return(entity.User{}, assert.AnError)
			},
			want:    entity.User{},
			wantErr: true,
		},
		{
			name: "failed to fetch from user domain",
			arg: args{
				ctx:         context.Background(),
				insertParam: mockInsertParam,
			},
			mockFunc: func(mock mockInterface, arg args) {
				mock.userDom.EXPECT().Get(arg.ctx, mockUserFetchParam).Return(entity.User{}, assert.AnError)
			},
			want:    entity.User{},
			wantErr: true,
		},
	}

	// Iterate the test in here
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFunc(mocks, tt.arg)

			got, err := usecase.Create(tt.arg.ctx, tt.arg.insertParam)
			if (err != nil) != tt.wantErr {
				t.Errorf("usecase.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				assert.Equal(t, tt.want, got)
				t.Errorf("usecase.Create() got = %v, want %v", got, tt.want)
				return
			}
		})
	}
}

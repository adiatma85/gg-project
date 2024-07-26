package category

import (
	"context"
	"reflect"
	"testing"
	"time"

	mock_category_dom "github.com/adiatma85/gg-project/src/business/domain/mock/category"
	"github.com/adiatma85/gg-project/src/business/entity"
	"github.com/adiatma85/own-go-sdk/jwtAuth"
	"github.com/adiatma85/own-go-sdk/null"
	mock_jwt_auth "github.com/adiatma85/own-go-sdk/tests/mock/jwtAuth"
	mock_log "github.com/adiatma85/own-go-sdk/tests/mock/log"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

type mockInterface struct {
	logger      *mock_log.MockInterface
	categoryDom *mock_category_dom.MockInterface
	jwtAuth     *mock_jwt_auth.MockInterface
}

func initMockTest(t *testing.T) (Interface, category, mockInterface) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	logger := mock_log.NewMockInterface(ctrl)
	logger.EXPECT().Info(gomock.Any(), gomock.Any()).AnyTimes()
	logger.EXPECT().Error(gomock.Any(), gomock.Any()).AnyTimes()
	logger.EXPECT().Debug(gomock.Any(), gomock.Any()).AnyTimes()

	mockCategoryDom := mock_category_dom.NewMockInterface(ctrl)
	mockJwtAuth := mock_jwt_auth.NewMockInterface(ctrl)

	ucInterface := Init(InitParam{
		Log:      logger,
		Category: mockCategoryDom,
		JwtAuth:  mockJwtAuth,
	})

	ucStruct := category{
		log:      logger,
		category: mockCategoryDom,
		jwtAuth:  mockJwtAuth,
	}

	mockInterface := mockInterface{
		logger:      logger,
		categoryDom: mockCategoryDom,
		jwtAuth:     mockJwtAuth,
	}

	return ucInterface, ucStruct, mockInterface
}

func Test_category_Create(t *testing.T) {
	usecase, _, mocks := initMockTest(t)

	// Type in here
	type args struct {
		ctx context.Context
		req entity.CreateCategoryParam
	}

	// Mock in here
	mockTime := time.Now()
	Now = func() time.Time {
		return mockTime
	}
	mockCreateCategoryArgs := entity.CreateCategoryParam{
		Name: "Ini adalah nama yang panjang dan lebar",
	}

	mockUserInfo := jwtAuth.UserAuthInfo{
		User: jwtAuth.User{
			ID: 10,
		},
	}

	mockFinalInserParam := entity.CreateCategoryParam{
		Name:      "Ini adalah nama yang panjang dan lebar",
		CreatedBy: null.StringFrom("10"),
		UpdatedBy: null.StringFrom("10"),
	}

	mockCategoryResult := entity.Category{
		ID:        1,
		Name:      "Ini adalah nama yang panjang dan lebar",
		CreatedBy: null.StringFrom("10"),
		UpdatedBy: null.StringFrom("10"),
		CreatedAt: null.TimeFrom(mockTime),
		UpdatedAt: null.TimeFrom(mockTime),
		Status:    null.Int64From(1),
	}

	// Test cases in here
	tests := []struct {
		name     string
		arg      args
		mockFunc func(mock mockInterface, arg args)
		want     entity.Category
		wantErr  bool
	}{
		{
			name: "failed to get user info",
			arg: args{
				ctx: context.Background(),
				req: mockCreateCategoryArgs,
			},
			mockFunc: func(mock mockInterface, arg args) {
				mock.jwtAuth.EXPECT().GetUserAuthInfo(arg.ctx).Return(jwtAuth.UserAuthInfo{}, assert.AnError)
			},
			want:    entity.Category{},
			wantErr: true,
		},
		{
			name: "failed to insert new category to category domain",
			arg: args{
				ctx: context.Background(),
				req: mockCreateCategoryArgs,
			},
			mockFunc: func(mock mockInterface, arg args) {
				mock.jwtAuth.EXPECT().GetUserAuthInfo(arg.ctx).Return(mockUserInfo, nil)
				mock.categoryDom.EXPECT().Create(arg.ctx, mockFinalInserParam).Return(entity.Category{}, assert.AnError)
			},
			want:    entity.Category{},
			wantErr: true,
		},
		{
			name: "success",
			arg: args{
				ctx: context.Background(),
				req: mockCreateCategoryArgs,
			},
			mockFunc: func(mock mockInterface, arg args) {
				mock.jwtAuth.EXPECT().GetUserAuthInfo(arg.ctx).Return(mockUserInfo, nil)
				mock.categoryDom.EXPECT().Create(arg.ctx, mockFinalInserParam).Return(mockCategoryResult, nil)
			},
			want:    mockCategoryResult,
			wantErr: false,
		},
	}

	// Iterate the tests in here
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFunc(mocks, tt.arg)

			got, err := usecase.Create(tt.arg.ctx, tt.arg.req)
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

// func Test_category_Get(t *testing.T) {
// 	usecase, _, mocks := initMockTest(t)

// 	// Type in here
// 	type args struct {
// 		ctx    context.Context
// 		params entity.CategoryParam
// 	}

// 	// Mock in here

// 	// Test cases in here
// 	tests := []struct {
// 		name     string
// 		arg      args
// 		mockFunc func(mock mockInterface, arg args)
// 		want     entity.Category
// 		wantErr  bool
// 	}{}

// 	// Iterate the tests in here
// }

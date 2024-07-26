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
	"github.com/adiatma85/own-go-sdk/query"
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

func Test_category_Get(t *testing.T) {
	usecase, _, mocks := initMockTest(t)

	// Type in here
	type args struct {
		ctx    context.Context
		params entity.CategoryParam
	}

	// Mock in here
	mockCategoryParam := entity.CategoryParam{
		ID: null.Int64From(1),
		QueryOption: query.Option{
			IsActive: true,
		},
	}

	mockCategoryResult := entity.Category{
		ID:     1,
		Name:   "Nama Kategori yang panjang dan lebar",
		Status: null.Int64From(1),
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
			name: "failed to get from category domain",
			arg: args{
				ctx:    context.Background(),
				params: mockCategoryParam,
			},
			mockFunc: func(mock mockInterface, arg args) {
				mock.categoryDom.EXPECT().Get(arg.ctx, mockCategoryParam).Return(entity.Category{}, assert.AnError)
			},
			want:    entity.Category{},
			wantErr: true,
		},
		{
			name: "success",
			arg: args{
				ctx:    context.Background(),
				params: mockCategoryParam,
			},
			mockFunc: func(mock mockInterface, arg args) {
				mock.categoryDom.EXPECT().Get(arg.ctx, mockCategoryParam).Return(mockCategoryResult, nil)
			},
			want:    mockCategoryResult,
			wantErr: false,
		},
	}

	// Iterate the tests in here
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFunc(mocks, tt.arg)

			got, err := usecase.Get(tt.arg.ctx, tt.arg.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("usecase.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				assert.Equal(t, tt.want, got)
				t.Errorf("usecase.Get() got = %v, want %v", got, tt.want)
				return
			}
		})
	}
}

func Test_category_GetAsAdmin(t *testing.T) {
	usecase, _, mocks := initMockTest(t)

	// Type in here
	type args struct {
		ctx    context.Context
		params entity.CategoryParam
	}

	// Mock in here
	mockCategoryParam := entity.CategoryParam{
		ID: null.Int64From(1),
	}

	mockCategoryResult := entity.Category{
		ID:     1,
		Name:   "Nama Kategori yang panjang dan lebar",
		Status: null.Int64From(1),
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
			name: "failed to get from category domain",
			arg: args{
				ctx:    context.Background(),
				params: mockCategoryParam,
			},
			mockFunc: func(mock mockInterface, arg args) {
				mock.categoryDom.EXPECT().Get(arg.ctx, mockCategoryParam).Return(entity.Category{}, assert.AnError)
			},
			want:    entity.Category{},
			wantErr: true,
		},
		{
			name: "success",
			arg: args{
				ctx:    context.Background(),
				params: mockCategoryParam,
			},
			mockFunc: func(mock mockInterface, arg args) {
				mock.categoryDom.EXPECT().Get(arg.ctx, mockCategoryParam).Return(mockCategoryResult, nil)
			},
			want:    mockCategoryResult,
			wantErr: false,
		},
	}

	// Iterate the tests in here
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFunc(mocks, tt.arg)

			got, err := usecase.GetAsAdmin(tt.arg.ctx, tt.arg.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("usecase.GetAsAdmin() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				assert.Equal(t, tt.want, got)
				t.Errorf("usecase.GetAsAdmin() got = %v, want %v", got, tt.want)
				return
			}
		})
	}
}

func Test_category_GetList(t *testing.T) {
	usecase, _, mocks := initMockTest(t)

	// Type in here
	type args struct {
		ctx    context.Context
		params entity.CategoryParam
	}

	// Mock in here
	mockParams := entity.CategoryParam{
		ID: null.Int64From(1),
		PaginationParam: entity.PaginationParam{
			IncludePagination: true,
		},
		QueryOption: query.Option{
			IsActive: true,
		},
	}

	mockPagination := entity.Pagination{
		CurrentPage:     1,
		CurrentElements: 1,
		TotalPages:      1,
		TotalElements:   1,
	}

	mockResult := []entity.Category{
		{
			ID:     1,
			Status: null.Int64From(1),
		},
	}

	// Test cases in here
	tests := []struct {
		name       string
		arg        args
		mockFunc   func(mock mockInterface, arg args)
		want       []entity.Category
		pagination *entity.Pagination
		wantErr    bool
	}{
		{
			name: "error fetching from domain",
			arg: args{
				ctx:    context.Background(),
				params: mockParams,
			},
			mockFunc: func(mock mockInterface, arg args) {
				mock.categoryDom.EXPECT().GetList(arg.ctx, mockParams).Return([]entity.Category{}, nil, assert.AnError)
			},
			want:       nil,
			pagination: nil,
			wantErr:    true,
		},
		{
			name: "error fetching from domain",
			arg: args{
				ctx:    context.Background(),
				params: mockParams,
			},
			mockFunc: func(mock mockInterface, arg args) {
				mock.categoryDom.EXPECT().GetList(arg.ctx, mockParams).Return(mockResult, &mockPagination, nil)
			},
			want:       mockResult,
			pagination: &mockPagination,
			wantErr:    false,
		},
	}

	// Iterate the tests cases in here
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFunc(mocks, tt.arg)

			got, pagination, err := usecase.GetList(tt.arg.ctx, tt.arg.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("usecase.GetList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				assert.Equal(t, tt.want, got)
				t.Errorf("usecase.GetList() got = %v, want %v", got, tt.want)
				return
			}

			if !reflect.DeepEqual(pagination, tt.pagination) {
				assert.Equal(t, tt.pagination, pagination)
				t.Errorf("usecase.GetList() got = %v, want %v", pagination, tt.pagination)
				return
			}
		})
	}
}

func Test_category_GetListAsAdmin(t *testing.T) {
	usecase, _, mocks := initMockTest(t)

	// Type in here
	type args struct {
		ctx    context.Context
		params entity.CategoryParam
	}

	// Mock in here
	mockParams := entity.CategoryParam{
		ID: null.Int64From(1),
		PaginationParam: entity.PaginationParam{
			IncludePagination: true,
		},
	}

	mockPagination := entity.Pagination{
		CurrentPage:     1,
		CurrentElements: 1,
		TotalPages:      1,
		TotalElements:   1,
	}

	mockResult := []entity.Category{
		{
			ID:     1,
			Status: null.Int64From(1),
		},
	}

	// Test cases in here
	tests := []struct {
		name       string
		arg        args
		mockFunc   func(mock mockInterface, arg args)
		want       []entity.Category
		pagination *entity.Pagination
		wantErr    bool
	}{
		{
			name: "error fetching from domain",
			arg: args{
				ctx:    context.Background(),
				params: mockParams,
			},
			mockFunc: func(mock mockInterface, arg args) {
				mock.categoryDom.EXPECT().GetList(arg.ctx, mockParams).Return([]entity.Category{}, nil, assert.AnError)
			},
			want:       nil,
			pagination: nil,
			wantErr:    true,
		},
		{
			name: "error fetching from domain",
			arg: args{
				ctx:    context.Background(),
				params: mockParams,
			},
			mockFunc: func(mock mockInterface, arg args) {
				mock.categoryDom.EXPECT().GetList(arg.ctx, mockParams).Return(mockResult, &mockPagination, nil)
			},
			want:       mockResult,
			pagination: &mockPagination,
			wantErr:    false,
		},
	}

	// Iterate the tests cases in here
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFunc(mocks, tt.arg)

			got, pagination, err := usecase.GetListAsAdmin(tt.arg.ctx, tt.arg.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("usecase.GetListAsAdmin() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				assert.Equal(t, tt.want, got)
				t.Errorf("usecase.GetListAsAdmin() got = %v, want %v", got, tt.want)
				return
			}

			if !reflect.DeepEqual(pagination, tt.pagination) {
				assert.Equal(t, tt.pagination, pagination)
				t.Errorf("usecase.GetListAsAdmin() got = %v, want %v", pagination, tt.pagination)
				return
			}
		})
	}
}

func Test_category_Update(t *testing.T) {
	usecase, _, mocks := initMockTest(t)

	// Type in here
	type args struct {
		ctx         context.Context
		updateParam entity.UpdateCategoryParam
		selectParam entity.CategoryParam
	}

	// Mock in here
	mockTime := time.Now()
	Now = func() time.Time {
		return mockTime
	}
	mockUpdateArgs := entity.UpdateCategoryParam{
		Name: "Nama yang diganti",
	}

	mockSelectArgs := entity.CategoryParam{
		ID: null.Int64From(1),
		QueryOption: query.Option{
			IsActive: true,
		},
	}

	mockUserInfo := jwtAuth.UserAuthInfo{
		User: jwtAuth.User{
			ID: 10,
		},
	}

	mockUpdateParam := entity.UpdateCategoryParam{
		Name:      "Nama yang diganti",
		UpdatedAt: null.TimeFrom(mockTime),
		UpdatedBy: null.StringFrom("10"),
	}

	// Test cases in here
	tests := []struct {
		name     string
		arg      args
		mockFunc func(mock mockInterface, arg args)
		wantErr  bool
	}{
		{
			name: "failed to get user info",
			arg: args{
				ctx:         context.Background(),
				updateParam: mockUpdateArgs,
				selectParam: mockSelectArgs,
			},
			mockFunc: func(mock mockInterface, arg args) {
				mock.jwtAuth.EXPECT().GetUserAuthInfo(arg.ctx).Return(jwtAuth.UserAuthInfo{}, assert.AnError)
			},
			wantErr: true,
		},
		{
			name: "failed to update to category domain",
			arg: args{
				ctx:         context.Background(),
				updateParam: mockUpdateArgs,
				selectParam: mockSelectArgs,
			},
			mockFunc: func(mock mockInterface, arg args) {
				mock.jwtAuth.EXPECT().GetUserAuthInfo(arg.ctx).Return(mockUserInfo, nil)
				mock.categoryDom.EXPECT().Update(arg.ctx, mockUpdateParam, mockSelectArgs).Return(assert.AnError)
			},
			wantErr: true,
		},
		{
			name: "success",
			arg: args{
				ctx:         context.Background(),
				updateParam: mockUpdateArgs,
				selectParam: mockSelectArgs,
			},
			mockFunc: func(mock mockInterface, arg args) {
				mock.jwtAuth.EXPECT().GetUserAuthInfo(arg.ctx).Return(mockUserInfo, nil)
				mock.categoryDom.EXPECT().Update(arg.ctx, mockUpdateParam, mockSelectArgs).Return(nil)
			},
			wantErr: false,
		},
	}

	// Iterate the test in here
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFunc(mocks, tt.arg)

			err := usecase.Update(tt.arg.ctx, tt.arg.updateParam, tt.arg.selectParam)
			if (err != nil) != tt.wantErr {
				t.Errorf("usecase.Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func Test_category_Delete(t *testing.T) {
	usecase, _, mocks := initMockTest(t)

	// Type in here
	type args struct {
		ctx         context.Context
		selectParam entity.CategoryParam
	}

	// Mock in here
	mockTime := time.Now()
	Now = func() time.Time {
		return mockTime
	}

	mockSelectArgs := entity.CategoryParam{
		ID: null.Int64From(1),
		QueryOption: query.Option{
			IsActive: true,
		},
	}

	mockUserInfo := jwtAuth.UserAuthInfo{
		User: jwtAuth.User{
			ID: 10,
		},
	}

	mockDeleteParam := entity.UpdateCategoryParam{
		Status:    null.Int64From(-1),
		DeletedAt: null.TimeFrom(mockTime),
		DeletedBy: null.StringFrom("10"),
	}

	// Test cases in here
	tests := []struct {
		name     string
		arg      args
		mockFunc func(mock mockInterface, arg args)
		wantErr  bool
	}{
		{
			name: "failed to get user info",
			arg: args{
				ctx:         context.Background(),
				selectParam: mockSelectArgs,
			},
			mockFunc: func(mock mockInterface, arg args) {
				mock.jwtAuth.EXPECT().GetUserAuthInfo(arg.ctx).Return(jwtAuth.UserAuthInfo{}, assert.AnError)
			},
			wantErr: true,
		},
		{
			name: "failed to update to category domain",
			arg: args{
				ctx:         context.Background(),
				selectParam: mockSelectArgs,
			},
			mockFunc: func(mock mockInterface, arg args) {
				mock.jwtAuth.EXPECT().GetUserAuthInfo(arg.ctx).Return(mockUserInfo, nil)
				mock.categoryDom.EXPECT().Update(arg.ctx, mockDeleteParam, mockSelectArgs).Return(assert.AnError)
			},
			wantErr: true,
		},
		{
			name: "success",
			arg: args{
				ctx:         context.Background(),
				selectParam: mockSelectArgs,
			},
			mockFunc: func(mock mockInterface, arg args) {
				mock.jwtAuth.EXPECT().GetUserAuthInfo(arg.ctx).Return(mockUserInfo, nil)
				mock.categoryDom.EXPECT().Update(arg.ctx, mockDeleteParam, mockSelectArgs).Return(nil)
			},
			wantErr: false,
		},
	}

	// Iterate the test in here
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFunc(mocks, tt.arg)

			err := usecase.Delete(tt.arg.ctx, tt.arg.selectParam)
			if (err != nil) != tt.wantErr {
				t.Errorf("usecase.Delete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func Test_category_Activate(t *testing.T) {
	usecase, _, mocks := initMockTest(t)

	// Type in here
	type args struct {
		ctx         context.Context
		selectParam entity.CategoryParam
	}

	// Mock in here
	mockTime := time.Now()
	Now = func() time.Time {
		return mockTime
	}

	mockSelectArgs := entity.CategoryParam{
		ID: null.Int64From(1),
		QueryOption: query.Option{
			IsActive: true,
		},
	}

	mockUserInfo := jwtAuth.UserAuthInfo{
		User: jwtAuth.User{
			ID: 10,
		},
	}

	mockActivateParam := entity.UpdateCategoryParam{
		Status:    null.Int64From(1),
		UpdatedAt: null.TimeFrom(mockTime),
		UpdatedBy: null.StringFrom("10"),
	}

	// Test cases in here
	tests := []struct {
		name     string
		arg      args
		mockFunc func(mock mockInterface, arg args)
		wantErr  bool
	}{
		{
			name: "failed to get user info",
			arg: args{
				ctx:         context.Background(),
				selectParam: mockSelectArgs,
			},
			mockFunc: func(mock mockInterface, arg args) {
				mock.jwtAuth.EXPECT().GetUserAuthInfo(arg.ctx).Return(jwtAuth.UserAuthInfo{}, assert.AnError)
			},
			wantErr: true,
		},
		{
			name: "failed to update to category domain",
			arg: args{
				ctx:         context.Background(),
				selectParam: mockSelectArgs,
			},
			mockFunc: func(mock mockInterface, arg args) {
				mock.jwtAuth.EXPECT().GetUserAuthInfo(arg.ctx).Return(mockUserInfo, nil)
				mock.categoryDom.EXPECT().Update(arg.ctx, mockActivateParam, mockSelectArgs).Return(assert.AnError)
			},
			wantErr: true,
		},
		{
			name: "success",
			arg: args{
				ctx:         context.Background(),
				selectParam: mockSelectArgs,
			},
			mockFunc: func(mock mockInterface, arg args) {
				mock.jwtAuth.EXPECT().GetUserAuthInfo(arg.ctx).Return(mockUserInfo, nil)
				mock.categoryDom.EXPECT().Update(arg.ctx, mockActivateParam, mockSelectArgs).Return(nil)
			},
			wantErr: false,
		},
	}

	// Iterate the test in here
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFunc(mocks, tt.arg)

			err := usecase.Activate(tt.arg.ctx, tt.arg.selectParam)
			if (err != nil) != tt.wantErr {
				t.Errorf("usecase.Activate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

package task

import (
	"context"
	"reflect"
	"testing"
	"time"

	mock_task_dom "github.com/adiatma85/gg-project/src/business/domain/mock/task"
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
	logger  *mock_log.MockInterface
	taskDom *mock_task_dom.MockInterface
	jwtAuth *mock_jwt_auth.MockInterface
}

func initMockTest(t *testing.T) (Interface, task, mockInterface) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	logger := mock_log.NewMockInterface(ctrl)
	logger.EXPECT().Info(gomock.Any(), gomock.Any()).AnyTimes()
	logger.EXPECT().Error(gomock.Any(), gomock.Any()).AnyTimes()
	logger.EXPECT().Debug(gomock.Any(), gomock.Any()).AnyTimes()

	mockTaskDom := mock_task_dom.NewMockInterface(ctrl)
	mockJwtAuth := mock_jwt_auth.NewMockInterface(ctrl)

	ucInterface := Init(InitParam{
		Log:     logger,
		Task:    mockTaskDom,
		JwtAuth: mockJwtAuth,
	})

	ucStruct := task{
		log:     logger,
		task:    mockTaskDom,
		jwtAuth: mockJwtAuth,
	}

	mockInterface := mockInterface{
		logger:  logger,
		taskDom: mockTaskDom,
		jwtAuth: mockJwtAuth,
	}

	return ucInterface, ucStruct, mockInterface
}

func Test_task_Create(t *testing.T) {
	usecase, _, mocks := initMockTest(t)

	// Type in here
	type args struct {
		ctx context.Context
		req entity.CreateTaskParam
	}

	// Mock in here
	mockTime := time.Now()
	Now = func() time.Time {
		return mockTime
	}

	mockCreateTaskArgs := entity.CreateTaskParam{
		Title: "Title yang panjang dan lebar",
	}

	mockUserInfo := jwtAuth.UserAuthInfo{
		User: jwtAuth.User{
			ID: 10,
		},
	}

	mockFinalInserParam := entity.CreateTaskParam{
		Title:     "Title yang panjang dan lebar",
		UserId:    10,
		CreatedBy: null.StringFrom("10"),
		UpdatedBy: null.StringFrom("10"),
	}

	mockTaskResult := entity.Task{
		ID:        1,
		Title:     "Title yang panjang dan lebar",
		CreatedBy: null.StringFrom("10"),
		UpdatedBy: null.StringFrom("10"),
		CreatedAt: null.TimeFrom(mockTime),
		UpdatedAt: null.TimeFrom(mockTime),
		Status:    1,
	}

	// Test cases in here
	tests := []struct {
		name     string
		arg      args
		mockFunc func(mock mockInterface, arg args)
		want     entity.Task
		wantErr  bool
	}{
		{
			name: "failed to get user info",
			arg: args{
				ctx: context.Background(),
				req: mockCreateTaskArgs,
			},
			mockFunc: func(mock mockInterface, arg args) {
				mock.jwtAuth.EXPECT().GetUserAuthInfo(arg.ctx).Return(jwtAuth.UserAuthInfo{}, assert.AnError)
			},
			want:    entity.Task{},
			wantErr: true,
		},
		{
			name: "failed to insert new task to task domain",
			arg: args{
				ctx: context.Background(),
				req: mockCreateTaskArgs,
			},
			mockFunc: func(mock mockInterface, arg args) {
				mock.jwtAuth.EXPECT().GetUserAuthInfo(arg.ctx).Return(mockUserInfo, nil)
				mock.taskDom.EXPECT().Create(arg.ctx, mockFinalInserParam).Return(entity.Task{}, assert.AnError)
			},
			want:    entity.Task{},
			wantErr: true,
		},
		{
			name: "success",
			arg: args{
				ctx: context.Background(),
				req: mockCreateTaskArgs,
			},
			mockFunc: func(mock mockInterface, arg args) {
				mock.jwtAuth.EXPECT().GetUserAuthInfo(arg.ctx).Return(mockUserInfo, nil)
				mock.taskDom.EXPECT().Create(arg.ctx, mockFinalInserParam).Return(mockTaskResult, nil)
			},
			want:    mockTaskResult,
			wantErr: false,
		},
	}

	// Iterate the test cases in here
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

func Test_task_Get(t *testing.T) {
	usecase, _, mocks := initMockTest(t)

	// Type in here
	type args struct {
		ctx    context.Context
		params entity.TaskParam
	}

	// Mock in here
	mockUserInfo := jwtAuth.UserAuthInfo{
		User: jwtAuth.User{
			ID: 10,
		},
	}

	mockTaskParam := entity.TaskParam{
		ID:     null.Int64From(1),
		UserId: null.Int64From(10),
		QueryOption: query.Option{
			IsActive: true,
		},
	}

	mockTaskResult := entity.Task{
		ID:     1,
		Title:  "Nama Kategori yang panjang dan lebar",
		Status: 1,
	}

	// Test cases in here
	tests := []struct {
		name     string
		arg      args
		mockFunc func(mock mockInterface, arg args)
		want     entity.Task
		wantErr  bool
	}{
		{
			name: "failed to get user info",
			arg: args{
				ctx:    context.Background(),
				params: mockTaskParam,
			},
			mockFunc: func(mock mockInterface, arg args) {
				mock.jwtAuth.EXPECT().GetUserAuthInfo(arg.ctx).Return(jwtAuth.UserAuthInfo{}, assert.AnError)
			},
			want:    entity.Task{},
			wantErr: true,
		},
		{
			name: "failed to get from task domain",
			arg: args{
				ctx:    context.Background(),
				params: mockTaskParam,
			},
			mockFunc: func(mock mockInterface, arg args) {
				mock.jwtAuth.EXPECT().GetUserAuthInfo(arg.ctx).Return(mockUserInfo, nil)
				mock.taskDom.EXPECT().Get(arg.ctx, mockTaskParam).Return(entity.Task{}, assert.AnError)
			},
			want:    entity.Task{},
			wantErr: true,
		},
		{
			name: "success",
			arg: args{
				ctx:    context.Background(),
				params: mockTaskParam,
			},
			mockFunc: func(mock mockInterface, arg args) {
				mock.jwtAuth.EXPECT().GetUserAuthInfo(arg.ctx).Return(mockUserInfo, nil)
				mock.taskDom.EXPECT().Get(arg.ctx, mockTaskParam).Return(mockTaskResult, nil)
			},
			want:    mockTaskResult,
			wantErr: false,
		},
	}

	// Iterate the test cases in here
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

func Test_task_GetList(t *testing.T) {
	usecase, _, mocks := initMockTest(t)

	// Type in here
	type args struct {
		ctx    context.Context
		params entity.TaskParam
	}

	// Mock In here
	mockUserInfo := jwtAuth.UserAuthInfo{
		User: jwtAuth.User{
			ID: 10,
		},
	}

	mockParams := entity.TaskParam{
		ID:     null.Int64From(1),
		UserId: null.Int64From(10),
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

	mockResult := []entity.Task{
		{
			ID:     1,
			Status: 1,
		},
	}

	// Test cases in here
	tests := []struct {
		name       string
		arg        args
		mockFunc   func(mock mockInterface, arg args)
		want       []entity.Task
		pagination *entity.Pagination
		wantErr    bool
	}{
		{
			name: "failed to get user info",
			arg: args{
				ctx:    context.Background(),
				params: mockParams,
			},
			mockFunc: func(mock mockInterface, arg args) {
				mock.jwtAuth.EXPECT().GetUserAuthInfo(arg.ctx).Return(jwtAuth.UserAuthInfo{}, assert.AnError)
			},
			want:    []entity.Task{},
			wantErr: true,
		},
		{
			name: "error fetching from domain",
			arg: args{
				ctx:    context.Background(),
				params: mockParams,
			},
			mockFunc: func(mock mockInterface, arg args) {
				mock.jwtAuth.EXPECT().GetUserAuthInfo(arg.ctx).Return(mockUserInfo, nil)
				mock.taskDom.EXPECT().GetList(arg.ctx, mockParams).Return([]entity.Task{}, nil, assert.AnError)
			},
			want:       []entity.Task{},
			pagination: nil,
			wantErr:    true,
		},
		{
			name: "success",
			arg: args{
				ctx:    context.Background(),
				params: mockParams,
			},
			mockFunc: func(mock mockInterface, arg args) {
				mock.jwtAuth.EXPECT().GetUserAuthInfo(arg.ctx).Return(mockUserInfo, nil)
				mock.taskDom.EXPECT().GetList(arg.ctx, mockParams).Return(mockResult, &mockPagination, nil)
			},
			want:       mockResult,
			pagination: &mockPagination,
			wantErr:    false,
		},
	}

	// Iterate the test in here
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

func Test_task_Update(t *testing.T) {
	usecase, _, mocks := initMockTest(t)

	// Type in here
	type args struct {
		ctx         context.Context
		updateParam entity.UpdateTaskParam
		selectParam entity.TaskParam
	}

	// Mock in here
	mockTime := time.Now()
	Now = func() time.Time {
		return mockTime
	}

	mockUpdateArgs := entity.UpdateTaskParam{
		Title: "Title yang panjang dan lebar",
	}

	mockSelectArgs := entity.TaskParam{
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

	mockUpdateParam := entity.UpdateTaskParam{
		Title:     "Title yang panjang dan lebar",
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
			name: "failed to update to task domain",
			arg: args{
				ctx:         context.Background(),
				updateParam: mockUpdateArgs,
				selectParam: mockSelectArgs,
			},
			mockFunc: func(mock mockInterface, arg args) {
				mock.jwtAuth.EXPECT().GetUserAuthInfo(arg.ctx).Return(mockUserInfo, nil)
				mock.taskDom.EXPECT().Update(arg.ctx, mockUpdateParam, mockSelectArgs).Return(assert.AnError)
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
				mock.taskDom.EXPECT().Update(arg.ctx, mockUpdateParam, mockSelectArgs).Return(nil)
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

func Test_task_Delete(t *testing.T) {
	usecase, _, mocks := initMockTest(t)

	// Type in here
	type args struct {
		ctx         context.Context
		selectParam entity.TaskParam
	}

	// Mock in here
	mockTime := time.Now()
	Now = func() time.Time {
		return mockTime
	}

	mockSelectArgs := entity.TaskParam{
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

	mockDeleteParam := entity.UpdateTaskParam{
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
			name: "failed to update to task domain",
			arg: args{
				ctx:         context.Background(),
				selectParam: mockSelectArgs,
			},
			mockFunc: func(mock mockInterface, arg args) {
				mock.jwtAuth.EXPECT().GetUserAuthInfo(arg.ctx).Return(mockUserInfo, nil)
				mock.taskDom.EXPECT().Update(arg.ctx, mockDeleteParam, mockSelectArgs).Return(assert.AnError)
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
				mock.taskDom.EXPECT().Update(arg.ctx, mockDeleteParam, mockSelectArgs).Return(nil)
			},
			wantErr: false,
		},
	}

	// Iterate the test cases in here
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

func Test_task_Activate(t *testing.T) {
	usecase, _, mocks := initMockTest(t)

	// Type in here
	type args struct {
		ctx         context.Context
		selectParam entity.TaskParam
	}

	// Mock in here
	mockTime := time.Now()
	Now = func() time.Time {
		return mockTime
	}

	mockSelectArgs := entity.TaskParam{
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

	mockActivateParam := entity.UpdateTaskParam{
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
			name: "failed to update to role domain",
			arg: args{
				ctx:         context.Background(),
				selectParam: mockSelectArgs,
			},
			mockFunc: func(mock mockInterface, arg args) {
				mock.jwtAuth.EXPECT().GetUserAuthInfo(arg.ctx).Return(mockUserInfo, nil)
				mock.taskDom.EXPECT().Update(arg.ctx, mockActivateParam, mockSelectArgs).Return(assert.AnError)
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
				mock.taskDom.EXPECT().Update(arg.ctx, mockActivateParam, mockSelectArgs).Return(nil)
			},
			wantErr: false,
		},
	}

	// Iterate the test cases in here
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

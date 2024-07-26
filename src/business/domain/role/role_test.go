package role

import (
	"context"
	"database/sql"
	"reflect"
	"testing"

	"github.com/adiatma85/gg-project/src/business/entity"

	libsql "github.com/adiatma85/own-go-sdk/sql"
	mock_log "github.com/adiatma85/own-go-sdk/tests/mock/log"
	mock_json "github.com/adiatma85/own-go-sdk/tests/mock/parser"

	"go.uber.org/mock/gomock"
)

func Test_role_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	logger := mock_log.NewMockInterface(ctrl)
	logger.EXPECT().Error(gomock.Any(), gomock.Any()).AnyTimes()
	logger.EXPECT().Info(gomock.Any(), gomock.Any()).AnyTimes()
	logger.EXPECT().Debug(gomock.Any(), gomock.Any()).AnyTimes()

	mockJsonParser := mock_json.NewMockJSONInterface(ctrl)

	// Type in here
	type args struct {
		ctx         context.Context
		createParam entity.CreateRoleParam
	}

	// Mock in here
	// query := regexp.QuoteMeta(`INSERT INTO category (name, created_by) VALUES (?, ?)`)
	// queryGet := regexp.QuoteMeta(getRole)

	// Test cases in here
	tests := []struct {
		name        string
		args        args
		prepSqlMock func() (*sql.DB, error)
		want        entity.Role
		wantErr     bool
	}{}

	// Iterate the test in here
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sqlServer, err := tt.prepSqlMock()
			if err != nil {
				t.Error(err)
			}
			defer sqlServer.Close()

			sqlClient := libsql.Init(libsql.Config{
				Driver: "sqlmock",
				Leader: libsql.ConnConfig{
					MockDB: sqlServer,
				},
				Follower: libsql.ConnConfig{
					MockDB: sqlServer,
				},
			}, logger, nil)

			// Initialize the Domain
			domain := Init(InitParam{
				Log:  logger,
				Db:   sqlClient,
				Json: mockJsonParser,
			})

			got, err := domain.Create(tt.args.ctx, tt.args.createParam)
			if (err != nil) != tt.wantErr {
				t.Errorf("category.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("category.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

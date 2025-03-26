package shop_app_test

import (
	"context"
	"errors"
	proto "pim-sys/gen/go/logs"
	app "pim-sys/internal/logs/app"
	"testing"

	mock "github.com/stretchr/testify/mock"
	"google.golang.org/grpc/metadata"
)

//go:generate mockery --testonly --name Storage --dir . --outpkg shop_app_test --output .

func TestLogs_GetLogs(t *testing.T) {
	type args struct {
		ctx     context.Context
		content *proto.GetLogsRequest
	}
	tests := []struct {
		name    string
		logs    func() *app.Logs
		args    args
		want    *proto.GetLogsResponse
		wantErr bool
	}{
		{
			name: "deafult",
			logs: func() *app.Logs {
				storage := &Storage{}
				storage.Mock.On(
					"GetLogs",
					mock.Anything,
					mock.Anything).Return(nil, nil)

				return &app.Logs{
					LogsStorage: storage,
				}
			},
			args: args{
				ctx:     metadata.NewIncomingContext(context.TODO(), metadata.New(map[string]string{"user_id": "1"})),
				content: &proto.GetLogsRequest{},
			},
			wantErr: false,
		},
		{
			name: "db error",
			logs: func() *app.Logs {
				storage := &Storage{}
				storage.Mock.On(
					"GetLogs",
					mock.Anything,
					mock.Anything).Return(nil, errors.New("err"))

				return &app.Logs{
					LogsStorage: storage,
				}
			},
			args: args{
				ctx:     metadata.NewIncomingContext(context.TODO(), metadata.New(map[string]string{"user_id": "1"})),
				content: &proto.GetLogsRequest{},
			},
			wantErr: true,
		},
		{
			name: "user id parse error",
			logs: func() *app.Logs {
				storage := &Storage{}
				storage.Mock.On(
					"GetLogs",
					mock.Anything,
					mock.Anything).Return(nil, nil)

				return &app.Logs{
					LogsStorage: storage,
				}
			},
			args: args{
				ctx:     metadata.NewIncomingContext(context.TODO(), metadata.New(map[string]string{"user_id": "asdasd"})),
				content: &proto.GetLogsRequest{},
			},
			wantErr: true,
		},
		{
			name: "user not found",
			logs: func() *app.Logs {
				storage := &Storage{}
				storage.Mock.On(
					"GetLogs",
					mock.Anything,
					mock.Anything).Return(nil, nil)

				return &app.Logs{
					LogsStorage: storage,
				}
			},
			args: args{
				ctx:     metadata.NewIncomingContext(context.TODO(), metadata.New(map[string]string{})),
				content: &proto.GetLogsRequest{},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.logs().GetLogs(tt.args.ctx, tt.args.content)
			if (err != nil) != tt.wantErr {
				t.Errorf("Logs.GetLogs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestLogs_GetGraph(t *testing.T) {
	type args struct {
		ctx     context.Context
		content *proto.GetGraphRequest
	}
	tests := []struct {
		name    string
		logs    func() *app.Logs
		args    args
		want    *proto.GetGraphResponse
		wantErr bool
	}{
		{
			name: "deafult",
			logs: func() *app.Logs {
				storage := &Storage{}
				storage.Mock.On(
					"GetMinDate",
					mock.Anything,
					mock.Anything).Return(nil, nil)
				storage.Mock.On(
					"GetSales",
					mock.Anything,
					mock.Anything,
					mock.Anything,
					mock.Anything).Return(nil, nil)

				return &app.Logs{
					LogsStorage: storage,
				}
			},
			args: args{
				ctx:     metadata.NewIncomingContext(context.TODO(), metadata.New(map[string]string{"user_id": "1"})),
				content: &proto.GetGraphRequest{DateFrom: 1, DateTo: 2},
			},
			wantErr: false,
		},
		{
			name: "db error (GetMinDate)",
			logs: func() *app.Logs {
				storage := &Storage{}
				storage.Mock.On(
					"GetMinDate",
					mock.Anything,
					mock.Anything).Return(errors.New("err"))
				storage.Mock.On(
					"GetSales",
					mock.Anything,
					mock.Anything,
					mock.Anything,
					mock.Anything).Return(nil, nil)

				return &app.Logs{
					LogsStorage: storage,
				}
			},
			args: args{
				ctx:     metadata.NewIncomingContext(context.TODO(), metadata.New(map[string]string{"user_id": "1"})),
				content: &proto.GetGraphRequest{DateFrom: 1, DateTo: 2},
			},
			wantErr: true,
		},
		{
			name: "db error(GetSales)",
			logs: func() *app.Logs {
				storage := &Storage{}
				storage.Mock.On(
					"GetMinDate",
					mock.Anything,
					mock.Anything).Return(nil, nil)
				storage.Mock.On(
					"GetSales",
					mock.Anything,
					mock.Anything,
					mock.Anything,
					mock.Anything).Return(nil, errors.New("err"))

				return &app.Logs{
					LogsStorage: storage,
				}
			},
			args: args{
				ctx:     metadata.NewIncomingContext(context.TODO(), metadata.New(map[string]string{"user_id": "1"})),
				content: &proto.GetGraphRequest{DateFrom: 1, DateTo: 2},
			},
			wantErr: true,
		},
		{
			name: "user__id parse err",
			logs: func() *app.Logs {
				storage := &Storage{}
				storage.Mock.On(
					"GetMinDate",
					mock.Anything,
					mock.Anything).Return(nil, nil)
				storage.Mock.On(
					"GetSales",
					mock.Anything,
					mock.Anything,
					mock.Anything,
					mock.Anything).Return(nil, nil)

				return &app.Logs{
					LogsStorage: storage,
				}
			},
			args: args{
				ctx:     metadata.NewIncomingContext(context.TODO(), metadata.New(map[string]string{"user_id": "asdasds"})),
				content: &proto.GetGraphRequest{DateFrom: 1, DateTo: 2},
			},
			wantErr: true,
		},
		{
			name: "user not found",
			logs: func() *app.Logs {
				storage := &Storage{}
				storage.Mock.On(
					"GetMinDate",
					mock.Anything,
					mock.Anything).Return(nil, nil)
				storage.Mock.On(
					"GetSales",
					mock.Anything,
					mock.Anything,
					mock.Anything,
					mock.Anything).Return(nil, nil)

				return &app.Logs{
					LogsStorage: storage,
				}
			},
			args: args{
				ctx:     metadata.NewIncomingContext(context.TODO(), metadata.New(map[string]string{})),
				content: &proto.GetGraphRequest{DateFrom: 1, DateTo: 2},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.logs().GetGraph(tt.args.ctx, tt.args.content)
			if (err != nil) != tt.wantErr {
				t.Errorf("Logs.GetGraph() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

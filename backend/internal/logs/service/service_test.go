package logs_service_test

import (
	"context"
	proto "pim-sys/gen/go/logs"
	service "pim-sys/internal/logs/service"
	"reflect"
	"testing"

	"github.com/stretchr/testify/mock"
)

//go:generate mockery --testonly --name Logs --dir . --outpkg logs_service_test --output .

func TestServerAPI_GetGraph(t *testing.T) {
	type args struct {
		ctx context.Context
		in  *proto.GetGraphRequest
	}
	tests := []struct {
		name    string
		s       func() *service.ServerAPI
		args    args
		want    *proto.GetGraphResponse
		wantErr bool
	}{
		{
			name: "default",
			s: func() *service.ServerAPI {
				logs := &Logs{}
				logs.Mock.On(
					"GetGraph",
					mock.Anything,
					mock.Anything,
				).Return(nil, nil)
				return &service.ServerAPI{Logs: logs}
			},
			args: args{
				ctx: context.TODO(),
				in: &proto.GetGraphRequest{
					DateFrom: 1,
					DateTo:   2,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s().GetGraph(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("ServerAPI.GetGraph() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ServerAPI.GetGraph() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServerAPI_GetLogs(t *testing.T) {
	type args struct {
		ctx context.Context
		in  *proto.GetLogsRequest
	}
	tests := []struct {
		name    string
		s       func() *service.ServerAPI
		args    args
		want    *proto.GetLogsResponse
		wantErr bool
	}{
		{
			name: "default",
			s: func() *service.ServerAPI {
				logs := &Logs{}
				logs.Mock.On(
					"GetLogs",
					mock.Anything,
					mock.Anything,
				).Return(nil, nil)
				return &service.ServerAPI{Logs: logs}
			},
			args: args{
				ctx: context.TODO(),
				in: &proto.GetLogsRequest{
					ProductId: 1,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s().GetLogs(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("ServerAPI.GetLogs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ServerAPI.GetLogs() = %v, want %v", got, tt.want)
			}
		})
	}
}
